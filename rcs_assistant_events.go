package plivo

type RcsAssistantEventService struct {
	client *Client
	RcsAssistantEvent
}

type RcsAssistantEventCreateParams struct {
}

type RcsAssistantEvent struct {
	ApiID       string   `json:"api_id,omitempty"`
	PhoneNumber string   `json:"phone_number,omitempty"`
	IsCapable   bool     `json:"is_capable,omitempty"`
	Features    []string `json:"features,omitempty"`
	Message     string   `json:"message,omitempty"`
	Error       string   `json:"error,omitempty"`
}

func (service *RcsAssistantEventService) Create(params RcsAssistantEventCreateParams) (response *RcsAssistantEvent, err error) {
	req, err := service.client.NewRequest("POST", params, "RCS/AssistantEvents")
	if err != nil {
		return
	}
	response = &RcsAssistantEvent{}
	err = service.client.ExecuteRequest(req, response)
	return
}