package plivo

type ComplianceRequirementService struct {
	client *Client
}

type GetComplianceRequirementResponse struct {
	APIID                   string `json:"api_id"`
	ComplianceRequirementID string `json:"compliance_requirement_id"`
	CountryIso2             string `json:"country_iso2"`
	NumberType              string `json:"number_type"`
	EndUserType             string `json:"end_user_type"`
	AcceptableDocumentTypes []struct {
		Name                string `json:"name"`
		Scope               string `json:"scope"`
		AcceptableDocuments []struct {
			DocumentTypeID   string `json:"document_type_id"`
			DocumentTypeName string `json:"document_type_name"`
		} `json:"acceptable_documents"`
	} `json:"acceptable_document_types"`
}

type ListComplianceRequirementParams struct {
	CountryIso2 string `json:"country_iso2,omitempty" url:"country_iso2,omitempty"`
	NumberType  string `json:"number_type,omitempty" url:"number_type,omitempty"`
	EndUserType string `json:"end_user_type,omitempty" url:"end_user_type,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty" url:"phone_number,omitempty"`
}
//type ListComplianceDocumentTypeResponse struct {
//	APIID string `json:"api_id"`
//	Meta  struct {
//		Limit      int         `json:"limit"`
//		Next       interface{} `json:"next"`
//		Offset     int         `json:"offset"`
//		Previous   interface{} `json:"previous"`
//		TotalCount int         `json:"total_count"`
//	} `json:"meta"`
//	Objects []struct {
//		CreatedAt      time.Time `json:"created_at"`
//		Description    string    `json:"description"`
//		DocumentName   string    `json:"document_name"`
//		DocumentTypeID string    `json:"document_type_id"`
//		Information    []struct {
//			FieldName    string `json:"field_name"`
//			FieldType    string `json:"field_type"`
//			FriendlyName string `json:"friendly_name"`
//			MaxLength    int    `json:"max_length,omitempty"`
//			MinLength    int    `json:"min_length,omitempty"`
//		} `json:"information"`
//		ProofRequired interface{} `json:"proof_required"`
//	} `json:"objects"`
//}

func (service *ComplianceRequirementService) Get(complianceRequirementId string) (response *GetComplianceRequirementResponse, err error) {
	req, err := service.client.NewRequest("GET", nil, "ComplianceRequirement/%s/", complianceRequirementId)
	response = &GetComplianceRequirementResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *ComplianceRequirementService) List(params ListComplianceRequirementParams) (response *GetComplianceRequirementResponse, err error) {
	request, err := service.client.NewRequest("GET", params, "ComplianceRequirement/")
	if err != nil {
		return
	}
	response = &GetComplianceRequirementResponse{}
	err = service.client.ExecuteRequest(request, response)
	return
}
