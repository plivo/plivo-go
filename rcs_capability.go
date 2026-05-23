package plivo

type RcsCapabilityService struct {
	client *Client
}

type RcsCapabilityParams struct {
	PhoneNumber string `json:"phone_number,omitempty" url:"phone_number,omitempty"`
	AgentUUID   string `json:"agent_uuid,omitempty" url:"agent_uuid,omitempty"`
}

type RcsCapabilityResponse struct {
	ApiID       string   `json:"api_id,omitempty"`
	PhoneNumber string   `json:"phone_number,omitempty"`
	IsCapable   bool     `json:"is_capable,omitempty"`
	Features    []string `json:"features,omitempty"`
	Message     string   `json:"message,omitempty"`
	Error       string   `json:"error,omitempty"`
}

func (service *RcsCapabilityService) Check(params RcsCapabilityParams) (response *RcsCapabilityResponse, err error) {
	req, err := service.client.NewRequest("GET", params, "RCS/Capability/")
	if err != nil {
		return
	}
	response = &RcsCapabilityResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}