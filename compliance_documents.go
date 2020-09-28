package plivo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"time"
)

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
		LastName                   string `json:"last_name"`
		FirstName                  string `json:"first_name"`
		DateOfBirth                string `json:"date_of_birth,omitempty"`
		AddressLine1               string `json:"address_line_1,omitempty"`
		AddressLine2               string `json:"address_line_2,omitempty"`
		City                       string `json:"city,omitempty"`
		Country                    string `json:"country,omitempty"`
		PostalCode                 string `json:"postal_code,omitempty"`
		UniqueIdentificationNumber string `json:"unique_identification_number,omitempty"`
		Nationality                string `json:"nationality,omitempty"`
		PlaceOfBirth               string `json:"place_of_birth,omitempty"`
		DateOfIssue                string `json:"date_of_issue,omitempty"`
		DateOfExpiration           string `json:"date_of_expiration,omitempty"`
		TypeOfUtility              string `json:"type_of_utility,omitempty"`
		BillingId                  string `json:"billing_id,omitempty"`
		BillingDate                string `json:"billing_date,omitempty"`
		BusinessName               string `json:"business_name,omitempty"`
		TypeOfId                   string `json:"type_of_id,omitempty"`
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
			LastName                   string `json:"last_name"`
			FirstName                  string `json:"first_name"`
			DateOfBirth                string `json:"date_of_birth,omitempty"`
			AddressLine1               string `json:"address_line_1,omitempty"`
			AddressLine2               string `json:"address_line_2,omitempty"`
			City                       string `json:"city,omitempty"`
			Country                    string `json:"country,omitempty"`
			PostalCode                 string `json:"postal_code,omitempty"`
			UniqueIdentificationNumber string `json:"unique_identification_number,omitempty"`
			Nationality                string `json:"nationality,omitempty"`
			PlaceOfBirth               string `json:"place_of_birth,omitempty"`
			DateOfIssue                string `json:"date_of_issue,omitempty"`
			DateOfExpiration           string `json:"date_of_expiration,omitempty"`
			TypeOfUtility              string `json:"type_of_utility,omitempty"`
			BillingId                  string `json:"billing_id,omitempty"`
			BillingDate                string `json:"billing_date,omitempty"`
			BusinessName               string `json:"business_name,omitempty"`
			TypeOfId                   string `json:"type_of_id,omitempty"`
		} `json:"meta_information"`
		FileName       string `json:"file_name,omitempty"`
		EndUserID      string `json:"end_user_id"`
		DocumentTypeID string `json:"document_type_id"`
	} `json:"objects"`
}

type CreateComplianceDocumentParams struct {
	FileName                   string
	EndUserID                  string `json:"end_user_id,omitempty"`
	DocumentTypeID             string `json:"document_type_id,omitempty"`
	Alias                      string `json:"alias,omitempty"`
	LastName                   string `json:"last_name,omitempty"`
	FirstName                  string `json:"first_name,omitempty"`
	DateOfBirth                string `json:"date_of_birth,omitempty"`
	AddressLine1               string `json:"address_line_1,omitempty"`
	AddressLine2               string `json:"address_line_2,omitempty"`
	City                       string `json:"city,omitempty"`
	Country                    string `json:"country,omitempty"`
	PostalCode                 string `json:"postal_code,omitempty"`
	UniqueIdentificationNumber string `json:"unique_identification_number,omitempty"`
	Nationality                string `json:"nationality,omitempty"`
	PlaceOfBirth               string `json:"place_of_birth,omitempty"`
	DateOfIssue                string `json:"date_of_issue,omitempty"`
	DateOfExpiration           string `json:"date_of_expiration,omitempty"`
	TypeOfUtility              string `json:"type_of_utility,omitempty"`
	BillingId                  string `json:"billing_id,omitempty"`
	BillingDate                string `json:"billing_date,omitempty"`
	BusinessName               string `json:"business_name,omitempty"`
	TypeOfId                   string `json:"type_of_id,omitempty"`
}

type UpdateComplianceDocumentParams struct {
	ComplianceDocumentID string `json:compliance_document_id`
	CreateComplianceDocumentParams
}

//type UpdateComplianceDocumentParams struct {
//	CreateEndUserParams
//	EndUserID string `json:"end_user_id"`
//}

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

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	if path != "" {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		fileContents, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		fi, err := file.Stat()
		if err != nil {
			return nil, err
		}
		file.Close()

		part, err := writer.CreateFormFile(paramName, fi.Name())
		if err != nil {
			return nil, err
		}
		part.Write(fileContents)
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	return http.NewRequest("POST", uri, body)
}

func (service *ComplianceDocumentService) Create(params CreateComplianceDocumentParams) (response *GetComplianceDocumentResponse, err error) {
	requestUrl := service.client.BaseUrl
	requestUrl.Path = fmt.Sprintf(baseRequestString, fmt.Sprintf(service.client.AuthId+"/ComplianceDocument/"))

	var requestParams map[string]string
	fields := reflect.TypeOf(params)
	values := reflect.ValueOf(params)
	num := fields.NumField()
	for i:=0; i<num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		requestParams[field.Name] = value.String()
	}

	request, err := newfileUploadRequest(requestUrl.String(), requestParams, "file", params.FileName)
	//request, err := service.client.NewRequest("POST", params, "ComplianceDocument/")
	if err != nil {
		return
	}
	response = &GetComplianceDocumentResponse{}
	err = service.client.ExecuteRequest(request, response)
	return
}

func (service *ComplianceDocumentService) Update(params UpdateComplianceDocumentParams) (response *GetComplianceDocumentResponse, err error) {
	requestUrl := service.client.BaseUrl
	requestUrl.Path = fmt.Sprintf(baseRequestString, fmt.Sprintf(service.client.AuthId+"/ComplianceDocument/"+params.ComplianceDocumentID+"/"))

	var requestParams map[string]string
	fields := reflect.TypeOf(params)
	values := reflect.ValueOf(params)
	num := fields.NumField()
	for i:=0; i<num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		requestParams[field.Name] = value.String()
	}

	request, err := newfileUploadRequest(requestUrl.String(), requestParams, "file", params.FileName)
	//request, err := service.client.NewRequest("POST", params, "ComplianceDocument/%s/", params.ComplianceDocumentID)
	if err != nil {
		return
	}
	response = &GetComplianceDocumentResponse{}
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
