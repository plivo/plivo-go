package plivo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

var expectedStatusCode = 200
var expectedResponse = ""
var requestUrl *url.URL
var requestMethod string
var requestHeader http.Header
var testAuthId = "AuthId"
var testAuthToken = "AuthId"

var server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	requestUrl = r.URL
	requestMethod = r.Method
	requestHeader = r.Header
	w.WriteHeader(expectedStatusCode)
	w.Write([]byte(expectedResponse))

	log.Println(expectedResponse)
}))

func expectResponse(fixturePath string, statusCode int) {
	if fixturePath != "" {
		// Some of the tests use uppercased names, so fix up them up.
		fixturePathRunes := []rune(fixturePath)
		fixturePathRunes[0] = unicode.ToLower(fixturePathRunes[0])
		fullFixturePath := fmt.Sprintf("fixtures/%s", string(fixturePathRunes))
		contents, err := ioutil.ReadFile(fullFixturePath)
		if err != nil {
			panic(err)
		}
		log.Printf("Loaded fixture from %s (%d)\n", fullFixturePath, len(contents))
		expectedResponse = string(contents)
	} else {
		expectedResponse = ""
	}

	expectedStatusCode = statusCode
}

func assertRequest(t *testing.T, method, path string, params ...interface{}) {

	path = fmt.Sprintf("/v1/Account/%s/%s/", client.AuthId, fmt.Sprintf(path, params...))
	expectedUrl, _ := url.Parse(path)

	if expectedUrl.Path != requestUrl.Path || method != requestMethod {
		log.Printf("expectedUrl: %s, requestUrl %s\nexpectedMethod: %s, requestMethod: %s\n", expectedUrl, requestUrl, method, requestMethod)
		t.FailNow()
	}

	assert.Contains(t, requestHeader.Get("User-Agent"), "plivo-go")
}

func assertBaseRequest(t *testing.T, method, path string, params ...interface{}) {

	expectedUrl, _ := url.Parse(fmt.Sprintf(path, params...))

	if expectedUrl.Path != requestUrl.Path || method != requestMethod {
		log.Printf("expectedUrl: %s, requestUrl %s\nexpectedMethod: %s, requestMethod: %s\n", expectedUrl, requestUrl, method, requestMethod)
		t.FailNow()
	}

	assert.Contains(t, requestHeader.Get("User-Agent"), "plivo-go")
}

func assertPhloRequest(t *testing.T, method, path string, params ...interface{}) {

	path = fmt.Sprintf("/v1/%s", fmt.Sprintf(path, params...))
	expectedUrl, _ := url.Parse(path)

	if expectedUrl.Path != requestUrl.Path || method != requestMethod {
		log.Printf("expectedUrl: %s, requestUrl %s\nexpectedMethod: %s, requestMethod: %s\n", expectedUrl, requestUrl, method, requestMethod)
		t.FailNow()
	}

	assert.Contains(t, requestHeader.Get("User-Agent"), "plivo-go")
}

func init() {
	client.BaseUrl, _ = url.Parse(server.URL)
	phloClient.BaseUrl, _ = url.Parse(server.URL)

	HttpsScheme = client.BaseUrl.Scheme
	voiceBaseUrlString = client.BaseUrl.Host
	voiceBaseUrlStringFallback1 = client.BaseUrl.Host
	voiceBaseUrlStringFallback2 = client.BaseUrl.Host
}

var client, _ = NewClient(testAuthId, testAuthToken, &ClientOptions{})
var phloClient, _ = NewPhloClient(testAuthId, testAuthToken, &ClientOptions{})
