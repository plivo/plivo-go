package plivo

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/url"
	"reflect"
	"sort"
	"strings"
)

func Numbers(numbers ...string) string {
	return strings.Join(numbers, "<")
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
			queryParamMap := getMapFromQueryString(parsedUrl.RawQuery)
			for k, v := range params {
				queryParamMap[k] = v
			}
			uri += GetSortedQueryParamString(queryParamMap, true)
		} else {
			uri += GetSortedQueryParamString(getMapFromQueryString(parsedUrl.RawQuery), true) + "." + GetSortedQueryParamString(params, false)
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

func getMapFromQueryString(query string) map[string]string {
	mp := make(map[string]string, 0)
	if len(query) == 0 {
		return mp
	}
	keyValuePairs := strings.Split(query, "&")
	sort.Strings(keyValuePairs)
	for _, element := range keyValuePairs {
		params := strings.SplitN(element, "=", 2)
		if len(params) == 2 {
			mp[params[0]] = params[1]
		}
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
	return messageMAC
}

func ValidateSignatureV3(uri, nonce, method, signature, authToken string, params map[string]string) bool {
	multipleSignatures := strings.Split(signature, ",")
	return Find(ComputeSignatureV3(authToken, uri, method, nonce, params), multipleSignatures)
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
