package plivo

type BulkMessageService struct {
	client *Client
	BulkMessage
}

type BulkMessageCreateParams struct {
	Src  string   `json:"src" url:"src"`
	Dst  []string `json:"dst" url:"dst"`
	Text string   `json:"text" url:"text"`
	// Optional parameters.
	Type          string      `json:"type,omitempty" url:"type,omitempty"`
	URL           string      `json:"url,omitempty" url:"url,omitempty"`
	Method        string      `json:"method,omitempty" url:"method,omitempty"`
	Log           interface{} `json:"log,omitempty" url:"log,omitempty"`
	PowerpackUUID string      `json:"powerpack_uuid,omitempty" url:"powerpack_uuid,omitempty"`
}

type BulkMessage struct {
	ApiID         string   `json:"api_id,omitempty" url:"api_id,omitempty"`
	MessageUUID   []string `json:"message_uuid,omitempty" url:"message_uuid,omitempty"`
	Message       string   `json:"message,omitempty" url:"message,omitempty"`
	InvalidNumber []string `json:"invalid_number,omitempty" url:"invalid_number,omitempty"`
}

// Stores response for sending a bulk message.
type BulkMessageCreateResponseBody struct {
	ApiID         string   `json:"api_id" url:"api_id"`
	MessageUUID   []string `json:"message_uuid" url:"message_uuid"`
	Message       string   `json:"message" url:"message"`
	InvalidNumber []string `json:"invalid_number,omitempty" url:"invalid_number,omitempty"`
}

func (service *BulkMessageService) Create(params BulkMessageCreateParams) (response *BulkMessageCreateResponseBody, err error) {
	req, err := service.client.NewRequest("POST", params, "Message/Bulk")
	if err != nil {
		return
	}
	response = &BulkMessageCreateResponseBody{}
	err = service.client.ExecuteRequest(req, response)
	return
}