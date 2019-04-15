package plivo

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

func Numbers(numbers ...string) string {
	return strings.Join(numbers, "<")
}


func getLanguageVoices() map[string][]string {
	languageVoices := make(map[string][]string)
	languageVoices["Australian English"] = append(languageVoices["Australian English"], "Nicole", "Russell")
	languageVoices["Brazilian Portuguese"] = append(languageVoices["Brazilian Portuguese"], "Vitória", "Ricardo")
	languageVoices["Canadian French"] = append(languageVoices["Canadian French"], "Chantal", "Chantal")
	languageVoices["Danish"] = append(languageVoices["Danish"], "Naja", "Mads")
	languageVoices["Dutch"] = append(languageVoices["Dutch"], "Lotte", "Ruben")
	languageVoices["French"] = append(languageVoices["French"], "Léa", "Céline", "Mathieu")
	languageVoices["German"] = append(languageVoices["German"], "Vicki", "Hans")
	languageVoices["Hindi"] = append(languageVoices["Hindi"], "Aditi")
	languageVoices["Icelandic"] = append(languageVoices["Icelandic"], "Dóra","Karl")
	languageVoices["Indian English"] = append(languageVoices["Indian English"], "Raveena", "Aditi")
	languageVoices["Italian"] = append(languageVoices["Italian"], "Carla", "Giorgio")
	languageVoices["Japanese"] = append(languageVoices["Japanese"], "Mizuki", "Takumi")
	languageVoices["Korean"] = append(languageVoices["Korean"], "Seoyeon")
	languageVoices["Mandarin Chinese"] = append(languageVoices["Mandarin Chinese"], "Zhiyu")
	languageVoices["Norwegian"] = append(languageVoices["Norwegian"], "Liv")
	languageVoices["Polish"] = append(languageVoices["Polish"], "Ewa", "Maja","Jacek","Jan")
	languageVoices["Portuguese-Iberic"] = append(languageVoices["Portuguese-Iberic"], "Inês", "Cristiano")
	languageVoices["Romanian"] = append(languageVoices["Romanian"], "Carmen")
	languageVoices["Russian"] = append(languageVoices["Russian"], "Tatyana","Maxim")
	languageVoices["Spanish-Castilian"] = append(languageVoices["Spanish-Castilian"], "Conchita","Enrique")
	languageVoices["Swedish"] = append(languageVoices["Swedish"], "Astrid")
	languageVoices["Turkish"] = append(languageVoices["Turkish"], "Filiz")
	languageVoices["UK English"] = append(languageVoices["UK English"], "Amy","Emma","Brian")
	languageVoices["US English"] = append(languageVoices["US English"], "Joanna", "Salli", "Kendra", "Kimberly", "Ivy", "Matthew", "Justin", "Joey")
	languageVoices["US Spanish"] = append(languageVoices["US Spanish"], "Penélope","Miguel")
	languageVoices["Welsh"] = append(languageVoices["Welsh"], "Gwyneth")
	languageVoices["Welsh English"] = append(languageVoices["Welsh English"], "Geraint")
	return languageVoices
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func ValidateLanguageVoice(language string, voice string) error{
	voiceparts := strings.Split(voice, ".")
	if len(voiceparts) != 2 || voiceparts[0] != "Polly" || len(voiceparts[1]) == 0 {
		return  errors.New("XML Validation Error: Invalid language. Voice " + voice + " is not valid." +
			" Refer <link> for the list of supported voices.")
	}

	languageVoicesList := getLanguageVoices()

	if (languageVoicesList[language] == nil) {
		return  errors.New("XML Validation Error: Invalid language. Language " + language + " is not supported.")
	}

	availableLanguageVoicesList := languageVoicesList[language]

	for i := range availableLanguageVoicesList {
		availableLanguageVoicesList[i] = TransformString(availableLanguageVoicesList[i])
	}
	transformedVoiceName := TransformString(voiceparts[1])

	if strings.Compare(voiceparts[1], "*") == 0 ||  Contains(availableLanguageVoicesList,transformedVoiceName) == false {
		return errors.New("XML Validation Error: <Speak> voice '" + voice + "' is not valid. Refer <link> for list of supported voices.")
	}
	return nil
}

func TransformString(s string) string{
	tc := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ = transform.String(tc, s)
	s = strings.Title(s)
	s = strings.Replace(s, " ", "_", -1)
	return s
}

func headersWithSep(headers map[string]string, keyValSep, itemSep string, escape bool) string {
	v := url.Values{}
	for key, value := range headers {
		v.Set(key, value)
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		var prefix string
		if escape {
			prefix = url.QueryEscape(k) + keyValSep
		} else {
			prefix = k + keyValSep
		}

		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteString(itemSep)
			}
			buf.WriteString(prefix)
			if escape {
				buf.WriteString(url.QueryEscape(v))
			} else {
				buf.WriteString(v)
			}
		}
	}
	return buf.String()
}

// Some code from encode.go from the Go Standard Library
func Headers(headers map[string]string) string {
	return headersWithSep(headers, "=", ",", true)
}

// The old signature validation is deprecated. Will be marked deprecated in the next release.

func ComputeSignature(authToken, uri string, params map[string]string) string {
	originalString := fmt.Sprintf("%s%s", uri, headersWithSep(params, "", "", false))
	logrus.Infof("originalString: %s\n", originalString)
	mac := hmac.New(sha1.New, []byte(authToken))
	mac.Write([]byte(originalString))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func ValidateSignature(authToken, uri string, params map[string]string, signature string) bool {
	return ComputeSignature(authToken, uri, params) == signature
}

// Adding V2 signature validation

func ComputeSignatureV2(authToken, uri string, nonce string) string {
	parsedUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}

	var originalString string = parsedUrl.Scheme + "://" + parsedUrl.Host + parsedUrl.Path + nonce
	mac := hmac.New(sha256.New, []byte(authToken))
	mac.Write([]byte(originalString))
	var messageMAC string = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return messageMAC
}

func ValidateSignatureV2(uri string, nonce string, signature string, authToken string) bool {
	return ComputeSignatureV2(authToken, uri, nonce) == signature
}

func GenerateUrl(uri string, params map[string]string, method string) string {
	parsedUrl, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	uri = parsedUrl.Scheme + "://" + parsedUrl.Host + parsedUrl.Path
	if len(params) > 0 || len(parsedUrl.RawQuery) > 0 {
		uri += "?"
	}
	if len(parsedUrl.RawQuery) > 0 {
		if method == "GET" {
			queryParamMap := getMapFromQueryString(parsedUrl.Query())
			for k, v := range params {
				queryParamMap[k] = v
			}
			uri += GetSortedQueryParamString(queryParamMap, true)
		} else {
			uri += GetSortedQueryParamString(getMapFromQueryString(parsedUrl.Query()), true) + "." + GetSortedQueryParamString(params, false)
			uri = strings.TrimRight(uri, ".")
		}
	} else {
		if method == "GET" {
			uri += GetSortedQueryParamString(params, true)
		} else {
			uri += GetSortedQueryParamString(params, false)
		}
	}
	return uri
}

func getMapFromQueryString(query url.Values) map[string]string {
	/*
		Example: input  "a=b&c=d&z=x"
		output (string): {"z":x", "c": "d", "a": "b"}
	*/
	mp := make(map[string]string, 0)
	if len(query) == 0 {
		return mp
	}
	for key, val := range query {
		mp[key] = val[0]
	}
	return mp
}

func GetSortedQueryParamString(params map[string]string, queryParams bool) string {
	url := ""
	keys := make([]string, 0, len(params))
	for key, _ := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	if queryParams {
		for _, key := range keys {
			url += key + "=" + params[key] + "&"
		}
		url = strings.TrimRight(url, "&")
	} else {
		for _, key := range keys {
			url += key + params[key]
		}
	}
	return url
}

func ComputeSignatureV3(authToken, uri, method string, nonce string, params map[string]string) string {
	var newUrl = GenerateUrl(uri, params, method) + "." + nonce
	mac := hmac.New(sha256.New, []byte(authToken))
	mac.Write([]byte(newUrl))
	var messageMAC = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	logrus.Info(messageMAC)
	return messageMAC
}

func ValidateSignatureV3(uri, nonce, method, signature, authToken string, params ...map[string]string) bool {
	parameters := map[string]string{}
	if len(params) != 0 {
		parameters = params[0]
	}
	multipleSignatures := strings.Split(signature, ",")
	return Find(ComputeSignatureV3(authToken, uri, method, nonce, parameters), multipleSignatures)
}

func Find(val string, slice []string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func checkAndFetchCallInsightsRequestDetails(param interface{}) (isCallInsightsRequest bool, requestPath string) {
	isCallInsightsRequest = false
	if reflect.TypeOf(param).Kind() == reflect.Map {
		if reflect.TypeOf(param).Key().Kind() == reflect.String {
			if _, ok := param.(map[string]interface{})[CallInsightsParams]; ok {
				isCallInsightsRequest = true
				requestPath = param.(map[string]interface{})[CallInsightsParams].(map[string]interface{})[CallInsightsRequestPath].(string)
			}
		}
	}
	return
}

func isVoiceRequest() (extraData map[string]interface{}) {
	extraData = make(map[string]interface{})
	extraData["is_voice_request"] = true
	extraData["retry"] = 0
	return extraData
}
