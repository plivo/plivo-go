package whatsappclient

import "github.com/plivo/plivo-go"

type WhatsAppTemplateService struct {
	client *plivo.Client
}

type WhatsAppTemplateComponent struct {
	Type       string                          `json:"type,omitempty"`
	Format     string                          `json:"format,omitempty"`
	Text       string                          `json:"text,omitempty"`
	Buttons    []WhatsAppTemplateButton        `json:"buttons,omitempty"`
	Cards      []WhatsAppTemplateCard          `json:"cards,omitempty"`
	Parameters []WhatsAppTemplateParameter     `json:"parameters,omitempty"`
	Example    *WhatsAppTemplateComponentExample `json:"example,omitempty"`
}

type WhatsAppTemplateButton struct {
	Type        string `json:"type,omitempty"`
	Text        string `json:"text,omitempty"`
	URL         string `json:"url,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	OTPType     string `json:"otp_type,omitempty"`
	AutofillText string `json:"autofill_text,omitempty"`
	PackageName  string `json:"package_name,omitempty"`
	SignatureHash string `json:"signature_hash,omitempty"`
}

type WhatsAppTemplateCard struct {
	Components []WhatsAppTemplateComponent `json:"components,omitempty"`
}

type WhatsAppTemplateParameter struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type WhatsAppTemplateComponentExample struct {
	HeaderText   []string   `json:"header_text,omitempty"`
	BodyText     [][]string `json:"body_text,omitempty"`
	HeaderHandle []string   `json:"header_handle,omitempty"`
}

type WhatsAppTemplateCreateParams struct {
	ElementName    string                      `json:"elementName,omitempty"`
	LanguageCode   string                      `json:"languageCode,omitempty"`
	LanguagePolicy string                      `json:"languagePolicy,omitempty"`
	Category       string                      `json:"category,omitempty"`
	Components     []WhatsAppTemplateComponent `json:"components,omitempty"`
	// Optional parameters.
	ApplicationID string `json:"application_id,omitempty"`
}

type WhatsAppTemplateUpdateParams struct {
	Category   string                      `json:"category,omitempty"`
	Components []WhatsAppTemplateComponent `json:"components,omitempty"`
	// Optional parameters.
	ApplicationID string `json:"application_id,omitempty"`
}

type WhatsAppTemplateResponse struct {
	ApiID   string `json:"api_id,omitempty"`
	Message string `json:"message,omitempty"`
}

type WhatsAppTemplate struct {
	ID           string                      `json:"id,omitempty"`
	ElementName  string                      `json:"elementName,omitempty"`
	LanguageCode string                      `json:"languageCode,omitempty"`
	Category     string                      `json:"category,omitempty"`
	Status       string                      `json:"status,omitempty"`
	Components   []WhatsAppTemplateComponent `json:"components,omitempty"`
}

type WhatsAppTemplateListResponse struct {
	ApiID    string             `json:"api_id,omitempty"`
	Response []WhatsAppTemplate `json:"response,omitempty"`
}

func (service *WhatsAppTemplateService) Create(wabaID string, params WhatsAppTemplateCreateParams) (response *WhatsAppTemplateResponse, err error) {
	req, err := service.client.NewRequest("POST", params, "WhatsApp/Template/%s/", wabaID)
	if err != nil {
		return
	}
	response = &WhatsAppTemplateResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *WhatsAppTemplateService) Update(wabaID string, templateID string, params WhatsAppTemplateUpdateParams) (response *WhatsAppTemplateResponse, err error) {
	req, err := service.client.NewRequest("POST", params, "WhatsApp/Template/%s/%s/", wabaID, templateID)
	if err != nil {
		return
	}
	response = &WhatsAppTemplateResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *WhatsAppTemplateService) Get(wabaID string, templateName string) (response *WhatsAppTemplateListResponse, err error) {
	req, err := service.client.NewRequest("GET", nil, "WhatsApp/Template/%s/%s/", wabaID, templateName)
	if err != nil {
		return
	}
	response = &WhatsAppTemplateListResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *WhatsAppTemplateService) List(wabaID string) (response *WhatsAppTemplateListResponse, err error) {
	req, err := service.client.NewRequest("GET", nil, "WhatsApp/Template/%s/", wabaID)
	if err != nil {
		return
	}
	response = &WhatsAppTemplateListResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *WhatsAppTemplateService) Delete(wabaID string, templateID string) (err error) {
	req, err := service.client.NewRequest("DELETE", nil, "WhatsApp/Template/%s/%s/", wabaID, templateID)
	if err != nil {
		return
	}
	err = service.client.ExecuteRequest(req, nil)
	return
}