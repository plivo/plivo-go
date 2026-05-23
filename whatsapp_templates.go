package plivo

type WhatsappTemplateService struct {
	client *Client
}

type WhatsappTemplateComponent struct {
	Type       string      `json:"type,omitempty"`
	SubType    string      `json:"sub_type,omitempty"`
	Index      string      `json:"index,omitempty"`
	Parameters []Parameter `json:"parameters,omitempty"`
	Cards      []Card      `json:"cards,omitempty"`
}

type WhatsappTemplateCreateParams struct {
	Name                string                      `json:"name,omitempty"`
	Category            string                      `json:"category,omitempty"`
	Language            string                      `json:"language,omitempty"`
	Components          []WhatsappTemplateComponent `json:"components,omitempty"`
	AllowCategoryChange bool                        `json:"allow_category_change,omitempty"`
}

type WhatsappTemplateUpdateParams struct {
	Name                string                      `json:"name,omitempty"`
	Category            string                      `json:"category,omitempty"`
	Language            string                      `json:"language,omitempty"`
	Components          []WhatsappTemplateComponent `json:"components,omitempty"`
	AllowCategoryChange bool                        `json:"allow_category_change,omitempty"`
}

type WhatsappTemplateListParams struct {
	TemplateName string `url:"template_name,omitempty"`
	Limit        int    `url:"limit,omitempty"`
	Offset       int    `url:"offset,omitempty"`
}

type WhatsappTemplateDeleteParams struct {
	Name string `url:"name"`
}

type WhatsappTemplateResponse struct {
	ApiID            string      `json:"api_id,omitempty"`
	TemplateID       string      `json:"template_id,omitempty"`
	TemplateName     string      `json:"template_name,omitempty"`
	TemplateStatus   string      `json:"template_status,omitempty"`
	TemplateCategory string      `json:"template_category,omitempty"`
	TemplateLanguage string      `json:"template_language,omitempty"`
	Status           string      `json:"status,omitempty"`
	Message          string      `json:"message,omitempty"`
	Error            interface{} `json:"error,omitempty"`
}

type WhatsappTemplateGetResponse struct {
	ApiID          string                      `json:"api_id,omitempty"`
	TemplateID     string                      `json:"template_id,omitempty"`
	Name           string                      `json:"name,omitempty"`
	Category       string                      `json:"category,omitempty"`
	Language       string                      `json:"language,omitempty"`
	Status         string                      `json:"status,omitempty"`
	Components     []WhatsappTemplateComponent `json:"components,omitempty"`
	QualityScore   interface{}                 `json:"quality_score,omitempty"`
	RejectedReason string                      `json:"rejected_reason,omitempty"`
	Message        string                      `json:"message,omitempty"`
	Error          interface{}                 `json:"error,omitempty"`
}

type WhatsappTemplateListMeta struct {
	Previous *string `json:"previous,omitempty"`
	Next     *string `json:"next,omitempty"`
	Offset   int64   `json:"offset,omitempty"`
	Limit    int64   `json:"limit,omitempty"`
}

type WhatsappTemplateListResponse struct {
	ApiID   string                      `json:"api_id,omitempty"`
	Objects []WhatsappTemplateGetResponse `json:"objects,omitempty"`
	Meta    WhatsappTemplateListMeta    `json:"meta,omitempty"`
	Status  string                      `json:"status,omitempty"`
	Message string                      `json:"message,omitempty"`
	Error   interface{}                 `json:"error,omitempty"`
}

func (service *WhatsappTemplateService) Create(wabaID string, params WhatsappTemplateCreateParams) (response *WhatsappTemplateResponse, err error) {
	req, err := service.client.NewRequest("POST", params, "WhatsApp/Template/%s/", wabaID)
	if err != nil {
		return
	}
	response = &WhatsappTemplateResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *WhatsappTemplateService) Update(wabaID string, templateID string, params WhatsappTemplateUpdateParams) (response *WhatsappTemplateResponse, err error) {
	req, err := service.client.NewRequest("POST", params, "WhatsApp/Template/%s/%s/", wabaID, templateID)
	if err != nil {
		return
	}
	response = &WhatsappTemplateResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *WhatsappTemplateService) Get(wabaID string, templateID string) (response *WhatsappTemplateGetResponse, err error) {
	req, err := service.client.NewRequest("GET", nil, "WhatsApp/Template/%s/%s/", wabaID, templateID)
	if err != nil {
		return
	}
	response = &WhatsappTemplateGetResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *WhatsappTemplateService) List(wabaID string, params WhatsappTemplateListParams) (response *WhatsappTemplateListResponse, err error) {
	req, err := service.client.NewRequest("GET", params, "WhatsApp/Template/%s/", wabaID)
	if err != nil {
		return
	}
	response = &WhatsappTemplateListResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *WhatsappTemplateService) Delete(wabaID string, templateID string, params WhatsappTemplateDeleteParams) (err error) {
	req, err := service.client.NewRequest("DELETE", params, "WhatsApp/Template/%s/%s/", wabaID, templateID)
	if err != nil {
		return
	}
	err = service.client.ExecuteRequest(req, nil)
	return
}