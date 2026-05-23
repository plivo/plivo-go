package plivo

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

const (
	voiceBaseUrlString          = "voice.plivo.com"
	voiceBaseUrlStringFallback1 = "voice1.plivo.com"
	voiceBaseUrlStringFallback2 = "voice2.plivo.com"
	CallInsightsBaseURL         = "insights.plivo.com"
	HttpsScheme                 = "https"
)

type Client struct {
	httpClient *http.Client

	AuthId    string
	AuthToken string

	BaseUrl   *url.URL
	userAgent string

	RequestInterceptor  func(request *http.Request)
	ResponseInterceptor func(response *http.Response)

	Account      AccountService
	Subaccounts  SubaccountService
	Applications ApplicationService
	Calls        CallService
	LiveCalls     LiveCallService
	QueuedCalls   QueuedCallService
	ConferenceCalls ConferenceService
	Endpoints    EndpointService
	Messages     MessageService
	Lookup       LookupService
	Powerpacks   PowerpackService
	Media        MediaService
	Numbers      NumberService
	PhoneNumbers PhoneNumberService
	Pricings     PricingService
	Recordings   RecordingService
	Identities   IdentityService
	Addresses    AddressService
	Nodes        *NodeService
	MultiPartyCall MultiPartyCallService
	RcsCapability RcsCapabilityService
}

func NewClient(authId string, authToken string, options *Options) (client *Client, err error) {
	_, authIdIsSet := os.LookupEnv("PLIVO_AUTH_ID")
	_, authTokenIsSet := os.LookupEnv("PLIVO_AUTH_TOKEN")

	if authId == "" {
		if authIdIsSet {
			authId = os.Getenv("PLIVO_AUTH_ID")
		} else {
			return nil, fmt.Errorf("authID not provided and PLIVO_AUTH_ID env var is not set")
		}
	}

	if authToken == "" {
		if authTokenIsSet {
			authToken = os.Getenv("PLIVO_AUTH_TOKEN")
		} else {
			return nil, fmt.Errorf("authToken not provided and PLIVO_AUTH_TOKEN env var is not set")
		}
	}

	if options == nil {
		options = &Options{}
	}

	httpClient := http.DefaultClient
	if options.HttpClient != nil {
		httpClient = options.HttpClient
	}

	baseUrl, err := url.Parse(fmt.Sprintf("https://api.plivo.com/v1/Account/%s/", authId))
	if err != nil {
		return
	}

	userAgent := fmt.Sprintf("plivo-go/%s (Go %s)", sdkVersion, goVersion())

	client = &Client{
		httpClient: httpClient,
		AuthId:     authId,
		AuthToken:  authToken,
		BaseUrl:    baseUrl,
		userAgent:  userAgent,
	}

	client.Account = AccountService{client: client}
	client.Subaccounts = SubaccountService{client: client}
	client.Applications = ApplicationService{client: client}
	client.Calls = CallService{client: client}
	client.LiveCalls = LiveCallService{client: client}
	client.QueuedCalls = QueuedCallService{client: client}
	client.ConferenceCalls = ConferenceService{client: client}
	client.Endpoints = EndpointService{client: client}
	client.Messages = MessageService{client: client}
	client.Lookup = LookupService{client: client}
	client.Powerpacks = PowerpackService{client: client}
	client.Media = MediaService{client: client}
	client.Numbers = NumberService{client: client}
	client.PhoneNumbers = PhoneNumberService{client: client}
	client.Pricings = PricingService{client: client}
	client.Recordings = RecordingService{client: client}
	client.Identities = IdentityService{client: client}
	client.Addresses = AddressService{client: client}
	client.MultiPartyCall = MultiPartyCallService{client: client}
	client.RcsCapability = RcsCapabilityService{client: client}

	return
}

func checkAndFetchCallInsightsRequestDetails(param interface{}) (bool, string) {
	if param == nil {
		return false, ""
	}
	paramStr, ok := param.(string)
	if !ok {
		return false, ""
	}
	re := regexp.MustCompile(`^[a-f0-9-]{36}$`)
	if re.MatchString(paramStr) {
		return false, ""
	}
	insightsRe := regexp.MustCompile(`insights`)
	if insightsRe.MatchString(paramStr) {
		return true, paramStr
	}
	return false, ""
}