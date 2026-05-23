package plivo

type VerifySessionService struct {
	client *Client
	VerifySession
}

type VerifySessionCreateParams struct {
	AppHash             string `json:"app_hash,omitempty" url:"app_hash,omitempty"`
	BrandName           string `json:"brand_name,omitempty" url:"brand_name,omitempty"`
	CodeLength          int    `json:"code_length,omitempty" url:"code_length,omitempty"`
	DLTEntityID         string `json:"dlt_entity_id,omitempty" url:"dlt_entity_id,omitempty"`
	DLTSenderID         string `json:"dlt_sender_id,omitempty" url:"dlt_sender_id,omitempty"`
	DLTTemplateCategory string `json:"dlt_template_category,omitempty" url:"dlt_template_category,omitempty"`
	DLTTemplateID       string `json:"dlt_template_id,omitempty" url:"dlt_template_id,omitempty"`
	DLTText             string `json:"dlt_text,omitempty" url:"dlt_text,omitempty"`
	DTMF                int    `json:"dtmf,omitempty" url:"dtmf,omitempty"`
	FraudCheck          string `json:"fraud_check,omitempty" url:"fraud_check,omitempty"`
	Text                string `json:"text,omitempty" url:"text,omitempty"`
}

type VerifySession struct {
	ApiID   string `json:"api_id,omitempty" url:"api_id,omitempty"`
	Message string `json:"message,omitempty" url:"message,omitempty"`
	Error   string `json:"error,omitempty" url:"error,omitempty"`
}

type VerifySessionCreateResponseBody struct {
	ApiID   string `json:"api_id" url:"api_id"`
	Message string `json:"message" url:"message"`
	Error   string `json:"error" url:"error"`
}

func (service *VerifySessionService) Create(params VerifySessionCreateParams) (response *VerifySessionCreateResponseBody, err error) {
	req, err := service.client.NewRequest("POST", params, "Verify/Session")
	if err != nil {
		return
	}
	response = &VerifySessionCreateResponseBody{}
	err = service.client.ExecuteRequest(req, response)
	return
}