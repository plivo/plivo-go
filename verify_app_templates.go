package plivo

type VerifyAppTemplateService struct {
	client *Client
}

type VerifyTemplate struct {
	TemplateUUID string `json:"template_uuid,omitempty"`
	Text         string `json:"text,omitempty"`
	FriendlyName string `json:"friendly_name,omitempty"`
	Locale       string `json:"locale,omitempty"`
}

type VerifyAppTemplateListResponse struct {
	ApiID     string           `json:"api_id"`
	Templates []VerifyTemplate `json:"templates"`
}

func (service *VerifyAppTemplateService) List() (response *VerifyAppTemplateListResponse, err error) {
	req, err := service.client.NewRequest("GET", nil, "Verify/App/templates/")
	if err != nil {
		return
	}
	response = &VerifyAppTemplateListResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}