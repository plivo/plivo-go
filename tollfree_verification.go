package plivo

import "time"

// TollFreeRequestVerificationService - TF verification service struct
type TollFreeRequestVerificationService struct {
	client *Client
}

// TollFreeResponse - Default response
type TollFreeResponse struct {
	APIID   string `json:"api_id" url:"api_id"`
	Message string `json:"message,omitempty" url:"message,omitempty"`
}

// TollFreeCreateResponse - Default response
type TollFreeCreateResponse struct {
	APIID   string `json:"api_id" url:"api_id"`
	Message string `json:"message,omitempty" url:"message,omitempty"`
	UUID    string `json:"uuid" url:"uuid"`
}

// TollFreeCreateParams - List of params to create a TF verification request
type TollFreeCreateParams struct {
	ProfileUUID           string `json:"profile_uuid,omitempty" url:"profile_uuid,omitempty"`
	Usecase               string `json:"usecase,omitempty" url:"usecase,omitempty"`
	UsecaseSummary        string `json:"usecase_summary,omitempty" url:"usecase_summary,omitempty"`
	MessageSample         string `json:"message_sample,omitempty" url:"message_sample,omitempty"`
	OptInImageURL         string `json:"optin_image_url,omitempty" url:"optin_image_url,omitempty"`
	OptInType             string `json:"optin_type,omitempty" url:"optin_type,omitempty"`
	Volume                string `json:"volume,omitempty" url:"volume,omitempty"`
	AdditionalInformation string `json:"additional_information,omitempty" url:"additional_information,omitempty"`
	ExtraData             string `json:"extra_data,omitempty" url:"extra_data,omitempty"`
	Number                string `json:"number,omitempty" url:"number,omitempty"`
	CallbackURL           string `json:"callback_url,omitempty" url:"callback_url,omitempty"`
	CallbackMethod        string `json:"callback_method,omitempty" url:"callback_method,omitempty"`
}

// TollFreeUpdateParams - List of update params to update in TF verification request
type TollFreeUpdateParams struct {
	ProfileUUID           string `json:"profile_uuid,omitempty" url:"profile_uuid,omitempty"`
	Usecase               string `json:"usecase,omitempty" url:"usecase,omitempty"`
	UsecaseSummary        string `json:"usecase_summary,omitempty" url:"usecase_summary,omitempty"`
	MessageSample         string `json:"message_sample,omitempty" url:"message_sample,omitempty"`
	OptInImageURL         string `json:"optin_image_url,omitempty" url:"optin_image_url,omitempty"`
	OptInType             string `json:"optin_type,omitempty" url:"optin_type,omitempty"`
	Volume                string `json:"volume,omitempty" url:"volume,omitempty"`
	AdditionalInformation string `json:"additional_information,omitempty" url:"additional_information,omitempty"`
	ExtraData             string `json:"extra_data,omitempty" url:"extra_data,omitempty"`
	CallbackURL           string `json:"callback_url,omitempty" url:"callback_url,omitempty"`
	CallbackMethod        string `json:"callback_method,omitempty" url:"callback_method,omitempty"`
}

// TollFreeListParams - List of params to search in list API
type TollFreeListParams struct {
	Number      string `json:"number,omitempty"  url:"number,omitempty"`
	Status      string `json:"status,omitempty"  url:"status,omitempty"`
	ProfileUUID string `json:"profile_uuid,omitempty" url:"profile_uuid,omitempty"`
	CreatedGT   string `json:"created__gt,omitempty" url:"created__gt,omitempty"`
	CreatedGTE  string `json:"created__gte,omitempty" url:"created__gte,omitempty"`
	CreatedLT   string `json:"created__lt,omitempty" url:"created__lt,omitempty"`
	CreatedLTE  string `json:"created__lte,omitempty" url:"created__lte,omitempty"`
	Usecase     string `json:"usecase,omitempty" url:"usecase,omitempty"`
	Limit       int64  `json:"limit,omitempty" url:"limit,omitempty"`
	Offset      int64  `json:"offset,omitempty" url:"offset,omitempty"`
}

// TollfreeVerificationRequest struct
type TollfreeVerificationRequest struct {
	UUID                  string    `json:"uuid"  url:"uuid"`
	ProfileUUID           string    `json:"profile_uuid" url:"profile_uuid"`
	Number                string    `json:"number" url:"number"`
	Usecase               string    `json:"usecase" url:"usecase"`
	UsecaseSummary        string    `json:"usecase_summary" url:"usecase_summary"`
	MessageSample         string    `json:"message_sample" url:"message_sample"`
	OptinImageURL         *string   `json:"optin_image_url" url:"optin_image_url"`
	OptinType             string    `json:"optin_type" url:"optin_type"`
	Volume                string    `json:"volume" url:"volume"`
	AdditionalInformation string    `json:"additional_information" url:"additional_information"`
	ExtraData             string    `json:"extra_data" url:"extra_data"`
	CallbackURL           string    `json:"callback_url" url:"callback_url"`
	CallbackMethod        string    `json:"callback_method" url:"callback_method"`
	Status                string    `json:"status" url:"status"`
	ErrorMessage          string    `json:"error_message" url:"error_message"`
	Created               time.Time `json:"created" url:"created"`
	LastModified          time.Time `json:"last_modified" url:"last_modified"`
}

// TollfreeVerificationRequestListResponse - list API response struct
type TollfreeVerificationRequestListResponse struct {
	APIID   string                        `json:"api_id" url:"api_id"`
	Meta    *Meta                         `json:"meta,omitempty" url:"meta,omitempty"`
	Objects []TollfreeVerificationRequest `json:"objects,omitempty" url:"objects,omitempty"`
}

// Create - create API for Tollfree Verification Request
func (service *TollFreeRequestVerificationService) Create(params TollFreeCreateParams) (response *TollFreeCreateResponse, err error) {
	req, err := service.client.NewRequest("POST", params, "TollfreeVerification")
	if err != nil {
		return
	}
	response = &TollFreeCreateResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

// Update - Update API for Tollfree Verification Request
func (service *TollFreeRequestVerificationService) Update(UUID string, params TollFreeUpdateParams) (response *TollFreeResponse, err error) {
	req, err := service.client.NewRequest("POST", params, "TollfreeVerification/%s", UUID)
	if err != nil {
		return
	}
	response = &TollFreeResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

// Get - Get API for Tollfree Verification Request
func (service *TollFreeRequestVerificationService) Get(UUID string) (response *TollfreeVerificationRequest, err error) {
	req, err := service.client.NewRequest("GET", nil, "TollfreeVerification/%s", UUID)
	if err != nil {
		return
	}
	response = &TollfreeVerificationRequest{}
	err = service.client.ExecuteRequest(req, response)
	return
}

// List - List API for Tollfree Verification Request
func (service *TollFreeRequestVerificationService) List(params TollFreeListParams) (response *TollfreeVerificationRequestListResponse, err error) {
	req, err := service.client.NewRequest("GET", params, "TollfreeVerification")
	if err != nil {
		return
	}
	response = &TollfreeVerificationRequestListResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

// Delete - Delete API for Tollfree Verification Request
func (service *TollFreeRequestVerificationService) Delete(UUID string) (err error) {
	req, err := service.client.NewRequest("DELETE", nil, "TollfreeVerification/%s", UUID)
	if err != nil {
		return
	}
	err = service.client.ExecuteRequest(req, nil)
	return
}
