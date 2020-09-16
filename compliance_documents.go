package plivo

import "time"

type ComplianceDocumentService struct {
	client *Client
}

type GetComplianceDocumentResponse struct {
	APIID           string `json:"api_id"`
	DocumentID      string `json:"document_id"`
	EndUserID       string `json:"end_user_id"`
	DocumentTypeID  string `json:"document_type_id"`
	Alias           string `json:"alias"`
	FileName        string `json:"file_name"`
	MetaInformation struct {
		LastName    string `json:"last_name"`
		FirstName   string `json:"first_name"`
		DateOfBirth string `json:"date_of_birth"`
	} `json:"meta_information"`
	CreatedAt string `json:"created_at"`
}


type ListComplianceDocumentResponse struct {
	APIID string `json:"api_id"`
	Meta  struct {
		Limit      int         `json:"limit"`
		Next       string      `json:"next"`
		Offset     int         `json:"offset"`
		Previous   interface{} `json:"previous"`
		TotalCount int         `json:"total_count"`
	} `json:"meta"`
	Objects []struct {
		CreatedAt            time.Time `json:"created_at"`
		ComplianceDocumentID string    `json:"compliance_document_id"`
		Alias                string    `json:"alias"`
		MetaInformation      struct {
			LastName    string `json:"last_name"`
			FirstName   string `json:"first_name"`
			DateOfBirth string `json:"date_of_birth"`
		} `json:"meta_information"`
		File           string `json:"file"`
		EndUserID      string `json:"end_user_id"`
		DocumentTypeID string `json:"document_type_id"`
	} `json:"objects"`
}

type CreateComplianceDocumentParams struct {

}

type UpdateComplianceDocumentParams struct {
	CreateEndUserParams
	EndUserID string `json:"end_user_id"`
}

func (service *ComplianceDocumentService) Get(complianceDocumentId string) (response *GetComplianceDocumentResponse, err error) {
	req, err := service.client.NewRequest("GET", nil, "ComplianceDocument/%s/", complianceDocumentId)
	response = &GetComplianceDocumentResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *ComplianceDocumentService) List(params EndUserListParams) (response *ListComplianceDocumentResponse, err error) {
	request, err := service.client.NewRequest("GET", params, "ComplianceDocument/")
	if err != nil {
		return
	}
	response = &ListComplianceDocumentResponse{}
	err = service.client.ExecuteRequest(request, response)
	return
}

func (service *ComplianceDocumentService) Create(params CreateEndUserParams) (response *CreateEndUserResponse, err error) {
	// TODO
	request, err := service.client.NewRequest("POST", params, "ComplianceDocument/")
	if err != nil {
		return
	}
	response = &CreateEndUserResponse{}
	err = service.client.ExecuteRequest(request, response)
	return
}

func (service *ComplianceDocumentService) Update(params UpdateEndUserParams) (response *CreateEndUserResponse, err error) {
	// TODO: file upload
	request, err := service.client.NewRequest("POST", params, "ComplianceDocument/%s/", )
	if err != nil {
		return
	}
	response = &CreateEndUserResponse{}
	err = service.client.ExecuteRequest(request, response)
	return
}

func (service *ComplianceDocumentService) Delete(complianceDocumentId string) (err error) {
	req, err := service.client.NewRequest("DELETE", nil, "ComplianceDocument/%s/", complianceDocumentId)
	if err != nil {
		return
	}
	err = service.client.ExecuteRequest(req, nil, isVoiceRequest())
	return
}
