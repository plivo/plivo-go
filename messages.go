package plivo

const (
	SMS = "sms"
	MMS = "mms"
)

type MessageService struct {
	client *Client
	Message
}

type MessageCreateParams struct {
	Src  string `json:"src,omitempty" url:"src,omitempty"`
	Dst  string `json:"dst,omitempty" url:"dst,omitempty"`
	Text string `json:"text,omitempty" url:"text,omitempty"`
	// Optional parameters.
	Type      string      `json:"type,omitempty" url:"type,omitempty"`
	URL       string      `json:"url,omitempty" url:"url,omitempty"`
	Method    string      `json:"method,omitempty" url:"method,omitempty"`
	Trackable bool        `json:"trackable,omitempty" url:"trackable,omitempty"`
	Log       interface{} `json:"log,omitempty" url:"log,omitempty"`
	MediaUrls []string    `json:"media_urls,omitempty" url:"media_urls,omitempty"`
	MediaIds  []string    `json:"media_ids,omitempty" url:"media_ids,omitempty"`
	// Either one of src and powerpackuuid should be given
	PowerpackUUID string `json:"powerpack_uuid,omitempty" url:"powerpack_uuid,omitempty"`
	MessageExpiry int    `json:"message_expiry,omitempty" url:"message_expiry,omitempty"`
}

type Message struct {
	ApiID                    string `json:"api_id,omitempty" url:"api_id,omitempty"`
	ToNumber                 string `json:"to_number,omitempty" url:"to_number,omitempty"`
	FromNumber               string `json:"from_number,omitempty" url:"from_number,omitempty"`
	CloudRate                string `json:"cloud_rate,omitempty" url:"cloud_rate,omitempty"`
	MessageType              string `json:"message_type,omitempty" url:"message_type,omitempty"`
	ResourceURI              string `json:"resource_uri,omitempty" url:"resource_uri,omitempty"`
	CarrierRate              string `json:"carrier_rate,omitempty" url:"carrier_rate,omitempty"`
	MessageDirection         string `json:"message_direction,omitempty" url:"message_direction,omitempty"`
	MessageState             string `json:"message_state,omitempty" url:"message_state,omitempty"`
	TotalAmount              string `json:"total_amount,omitempty" url:"total_amount,omitempty"`
	MessageUUID              string `json:"message_uuid,omitempty" url:"message_uuid,omitempty"`
	MessageTime              string `json:"message_time,omitempty" url:"message_time,omitempty"`
	ErrorCode                string `json:"error_code,omitempty" url:"error_code,omitempty"`
	PowerpackID              string `json:"powerpack_id,omitempty" url:"powerpack_id,omitempty"`
	RequesterIP              string `json:"requester_ip,omitempty" url:"requester_ip,omitempty"`
	IsDomestic               *bool  `json:"is_domestic,omitempty" url:"is_domestic,omitempty"`
	ReplacedSender           string `json:"replaced_sender,omitempty" url:"replaced_sender,omitempty"`
	TendlcCampaignID         string `json:"tendlc_campaign_id" url:"tendlc_campaign_id,omitempty"`
	TendlcRegistrationStatus string `json:"tendlc_registration_status" url:"tendlc_registration_status,omitempty"`
	DestinationCountryISO2   string `json:"destination_country_iso2" url:"destination_country_iso2,omitempty"`
}

// Stores response for ending a message.
type MessageCreateResponseBody struct {
	Message     string   `json:"message" url:"message"`
	ApiID       string   `json:"api_id" url:"api_id"`
	MessageUUID []string `json:"message_uuid" url:"message_uuid"`
	Error       string   `json:"error" url:"error"`
}

type MediaDeleteResponse struct {
	Error string `json:"error,omitempty"`
}
type MMSMedia struct {
	ApiID       string `json:"api_id,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	MediaID     string `json:"media_id,omitempty"`
	MediaURL    string `json:"media_url,omitempty"`
	MessageUUID string `json:"message_uuid,omitempty"`
	Size        int64  `json:"size,omitempty"`
}

type MessageList struct {
	ApiID string `json:"api_id" url:"api_id"`
	Meta  struct {
		Previous *string
		Next     *string
		Offset   int64
		Limit    int64
	} `json:"meta"`
	Objects []Message `json:"objects" url:"objects"`
}

type MediaListResponseBody struct {
	Objects []MMSMedia `json:"objects" url:"objects"`
}

type MessageListParams struct {
	Limit                     int    `url:"limit,omitempty"`
	Offset                    int    `url:"offset,omitempty"`
	PowerpackID               string `url:"powerpack_id,omitempty"`
	Subaccount                string `url:"subaccount,omitempty"`
	MessageDirection          string `url:"message_direction,omitempty"`
	MessageState              string `url:"message_state,omitempty"`
	ErrorCode                 int    `url:"error_code,omitempty"`
	MessageTime               string `url:"message_time,omitempty"`
	MessageTimeGreaterThan    string `url:"message_time__gt,omitempty"`
	MessageTimeGreaterOrEqual string `url:"message_time__gte,omitempty"`
	MessageTimeLessThan       string `url:"message_time__lt,omitempty"`
	MessageTimeLessOrEqual    string `url:"message_time__lte,omitempty"`
	TendlcCampaignID          string `url:"tendlc_campaign_id,omitempty"`
	TendlcRegistrationStatus  string `url:"tendlc_registration_status,omitempty"`
	DestinationCountryISO2    string `url:"destination_country_iso2,omitempty"`
}

func (service *MessageService) List(params MessageListParams) (response *MessageList, err error) {
	req, err := service.client.NewRequest("GET", params, "Message")
	if err != nil {
		return
	}
	response = &MessageList{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *MessageService) Get(messageUuid string) (response *Message, err error) {
	req, err := service.client.NewRequest("GET", nil, "Message/%s", messageUuid)
	if err != nil {
		return
	}
	response = &Message{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *MessageService) Create(params MessageCreateParams) (response *MessageCreateResponseBody, err error) {
	req, err := service.client.NewRequest("POST", params, "Message")
	if err != nil {
		return
	}
	response = &MessageCreateResponseBody{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *MessageService) ListMedia(messageUuid string) (response *MediaListResponseBody, err error) {
	req, err := service.client.NewRequest("GET", nil, "Message/%s/Media/", messageUuid)
	if err != nil {
		return
	}
	response = &MediaListResponseBody{}
	err = service.client.ExecuteRequest(req, response)
	return
}
