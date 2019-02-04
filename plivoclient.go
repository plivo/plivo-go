package plivo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"runtime"

	"github.com/google/go-querystring/query"
)

const baseUrlString = "https://api.plivo.com/"

const sdkVersion = "4.0.5"

type Client struct {
	httpClient *http.Client

	AuthId    string
	AuthToken string

	BaseUrl   *url.URL
	userAgent string

	Messages     *MessageService
	Accounts     *AccountService
	Subaccounts  *SubaccountService
	Applications *ApplicationService
	Endpoints    *EndpointService
	Numbers      *NumberService
	PhoneNumbers *PhoneNumberService
	Pricing      *PricingService // TODO Rename?
	Recordings   *RecordingService
	Calls        *CallService
	LiveCalls    *LiveCallService
	QueuedCalls  *QueuedCallService
	Conferences  *ConferenceService
	Addresses    *AddressService
	Identities   *IdentityService

	RequestInterceptor  func(request *http.Request)
	ResponseInterceptor func(response *http.Response)
}

type ClientOptions struct {
	HttpClient *http.Client
	//RequestInterceptor  func(request *http.Request)
	//ResponseInterceptor func(response *http.Response)
}

/*
To set a proxy for all requests, configure the Transport for the HttpClient passed in:

	&http.Client{
 		Transport: &http.Transport{
 			Proxy: http.ProxyURL("http//your.proxy.here"),
 		},
 	}

Similarly, to configure the timeout, set it on the HttpClient passed in:

	&http.Client{
 		Timeout: time.Minute,
 	}
*/
func NewClient(authId, authToken string, options *ClientOptions) (client *Client, err error) {
	if len(authId) == 0 {
		authId = os.Getenv("PLIVO_AUTH_ID")
	}
	if len(authToken) == 0 {
		authToken = os.Getenv("PLIVO_AUTH_TOKEN")
	}

	baseUrl, err := url.Parse(baseUrlString)

	httpClient := &http.Client{
		Timeout: time.Minute,
	}

	client = &Client{
		userAgent:  fmt.Sprintf("%s/%s (Go: %s)", "plivo-go", sdkVersion, runtime.Version()),
		BaseUrl:    baseUrl,
		httpClient: httpClient,

		AuthId:    authId,
		AuthToken: authToken,
	}

	if options.HttpClient != nil {
		client.httpClient = options.HttpClient
	}

	client.Messages = &MessageService{client: client}
	client.Accounts = &AccountService{client: client}
	client.Subaccounts = &SubaccountService{client: client}
	client.Applications = &ApplicationService{client: client}
	client.Endpoints = &EndpointService{client: client}
	client.Numbers = &NumberService{client: client}
	client.PhoneNumbers = &PhoneNumberService{client: client}
	client.Pricing = &PricingService{client: client}
	client.Recordings = &RecordingService{client: client}
	client.Calls = &CallService{client: client}
	client.LiveCalls = &LiveCallService{client: client}
	client.QueuedCalls = &QueuedCallService{client: client}
	client.Conferences = &ConferenceService{client: client}
	client.Addresses = &AddressService{client:client}
	client.Identities = &IdentityService{client:client}

	return
}

func (client *Client) NewRequest(method string, params interface{}, formatString string, formatParams ...interface{}) (request *http.Request, err error) {
	if client == nil || client.httpClient == nil {
		err = errors.New("client and httpClient cannot be nil")
		return
	}

	for i, param := range formatParams {
		if param == nil || param == "" {
			err = errors.New(fmt.Sprintf("Request path parameter #%d is nil/empty but should not be so.", i))
			return
		}
	}

	requestUrl := *client.BaseUrl
	requestUrl.Path = fmt.Sprintf("/v1/Account/%s/%s/", client.AuthId, fmt.Sprintf(formatString, formatParams...))

	var buffer = new(bytes.Buffer)
	if method == "GET" {
		var values url.Values
		if values, err = query.Values(params); err != nil {
			return
		}

		requestUrl.RawQuery = values.Encode()
	} else {
		if err = json.NewEncoder(buffer).Encode(params); err != nil {
			return
		}
	}

	request, err = http.NewRequest(method, requestUrl.String(), buffer)

	request.Header.Add("User-Agent", client.userAgent)
	request.Header.Add("Content-Type", "application/json")

	request.SetBasicAuth(client.AuthId, client.AuthToken)

	return
}

func (client *Client) ExecuteRequest(request *http.Request, body interface{}) (err error) {
	if client == nil {
		return errors.New("client cannot be nil")
	}

	if client.httpClient == nil {
		return errors.New("httpClient cannot be nil")
	}

	//if client.RequestInterceptor != nil {
	//	client.RequestInterceptor(request)
	//}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return
	}

	//if client.ResponseInterceptor != nil {
	//	responseCopy := *response
	//	client.ResponseInterceptor(&responseCopy)
	//}

	data, err := ioutil.ReadAll(response.Body)
	if err == nil && data != nil && len(data) > 0 {
		if response.StatusCode >= 200 && response.StatusCode < 300 {
			if body != nil {
				err = json.Unmarshal(data, body)
			}
		} else {
			err = errors.New(string(data))
		}
	}

	return
}
