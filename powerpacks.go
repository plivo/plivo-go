package plivo

import (
	"fmt"
	"strings"
)

type PowerpackService struct {
	client *Client
	Powerpack
}

type PowerackCreateParams struct {
	Name string `json:"name,omitempty"`
	// Optional parameters.
	StickySender    string `json:"sticky_sender,omitempty"`
	LocalConnect    string `json:"local_connect,omitempty"`
	ApplicationType string `json:"application_type,omitempty"`
	ApplicationID   string `json:"application_id,omitempty"`
}

type PowerackUpdateParams struct {
	// Optional parameters.
	Name            string `json:"name,omitempty"`
	StickySender    string `json:"sticky_sender,omitempty"`
	LocalConnect    string `json:"local_connect,omitempty"`
	ApplicationType string `json:"application_type,omitempty"`
	ApplicationID   string `json:"application_id,omitempty"`
}

type NumberPoolResponse struct {
	Number_pool_uuid              string `json:"number_pool_uuid,omitempty"`
	Number                        string `json:"number,omitempty"`
	Type                          string `json:"Type,omitempty"`
	Country_iso2                  string `json:"country_iso2,omitempty"`
	Service                       string `json:"service,omitempty"`
	Added_on                      string `json:"added_on,omitempty"`
	Account_phone_number_resource string `json:"account_phone_number_resource,omitempty"`
}

type NumberResponse struct {
	ApiID string `json:"api_id,omitempty"`
	NumberPoolResponse
	Error string `json:"error,omitempty"`
}

type ShortCode struct {
	Number_pool_uuid string `json:"number_pool_uuid,omitempty"`
	Shortcode        string `json:"shortcode,omitempty"`
	Country_iso2     string `json:"country_iso2,omitempty"`
	Added_on         string `json:"added_on,omitempty"`
	Service          string `json:"service,omitempty"`
}

type Tollfree struct {
	NumberPoolUUID string `json:"number_pool_uuid,omitempty"`
	Tollfree       string `json:"number,omitempty"`
	Country_iso2   string `json:"country_iso2,omitempty"`
	Added_on       string `json:"added_on,omitempty"`
	Service        string `json:"service,omitempty"`
}

type ShortCodeResponse struct {
	BaseListPPKResponse
	Objects []ShortCode `json:"objects" url:"objects"`
}
type TollfreeResponse struct {
	BaseListPPKResponse
	Objects []Tollfree `json:"objects" url:"objects"`
}
type FindShortCodeResponse struct {
	ApiID string `json:"api_id,omitempty"`
	ShortCode
	Error string `json:"error,omitempty"`
}
type FindTollfreeResponse struct {
	ApiID string `json:"api_id,omitempty"`
	Tollfree
	Error string `json:"error,omitempty"`
}

type Powerpack struct {
	UUID            string `json:"uuid,omitempty"`
	Name            string `json:"name,omitempty"`
	StickySender    bool   `json:"sticky_sender,omitempty"`
	LocalConnect    bool   `json:"local_connect,omitempty"`
	ApplicationType string `json:"application_type,omitempty"`
	ApplicationID   string `json:"application_id,omitempty"`
	NumberPoolUUID  string `json:"number_pool,omitempty"`
	CreatedOn       string `json:"created_on,omitempty"`
}

type PowerpackResponse struct {
	CreatedOn            string `json:"created_on,omitempty"`
	LocalConnect         bool   `json:"local_connect,omitempty"`
	Name                 string `json:"name,omitempty"`
	NumberPoolUUID       string `json:"number_pool,omitempty"`
	PowerpackResourceURI string `json:"powerpack_resource_uri,omitempty"`
	StickySender         bool   `json:"sticky_sender,omitempty"`
	UUID                 string `json:"uuid,omitempty"`
}

type PowerpackDeleteResponse struct {
	ApiID    string `json:"api_id,omitempty"`
	Response string `json:"response,omitempty"`
	Error    string `json:"error,omitempty"`
}

type NumberDeleteResponse struct {
	PowerpackDeleteResponse
}

type ShortcodeDeleteResponse struct {
	PowerpackDeleteResponse
}

type TollfreeDeleteResponse struct {
	PowerpackDeleteResponse
}

type PowerpackDeleteParams struct {
	UnrentNumbers bool `json:"unrent_numbers,omitempty"`
}

type NumberRemoveParams struct {
	Unrent bool `json:"unrent,omitempty"`
}

type PowerpackAddNumberOptions struct {
	//Service can be 'sms' or 'mms'. Default to 'sms' when not set.
	Service string `json:"service,omitempty" url:"service,omitempty"`
}

type PowerpackFindNumberOptions struct {
	//Service can be 'sms' or 'mms'. Default to 'sms' when not set.
	Service string `json:"service,omitempty" url:"service,omitempty"`
}

type PowerpackAddTollfreeOptions struct {
	//Service can be 'sms' or 'mms'. Default to 'sms' when not set.
	Service string `json:"service,omitempty" url:"service,omitempty"`
}

type PowerpackListShortcodeOptions struct {
	//Service can be 'sms' or 'mms'. Default to 'sms' when not set.
	Service string `json:"service,omitempty" url:"service,omitempty"`
}

type PowerpackListTollfreeOptions struct {
	//Service can be 'sms' or 'mms'. Default to 'sms' when not set.
	Service string `json:"service,omitempty" url:"service,omitempty"`
}

type PowerpackFindShortcodeOptions struct {
	//Service can be 'sms' or 'mms'. Default to 'sms' when not set.
	Service string `json:"service,omitempty" url:"service,omitempty"`
}

type PowerpackFindTollfreeOptions struct {
	//Service can be 'sms' or 'mms'. Default to 'sms' when not set.
	Service string `json:"service,omitempty" url:"service,omitempty"`
}

type PowerpackCreateResponseBody struct {
	ApiID string  `json:"api_id,omitempty"`
	Error *string `json:"error,omitempty" url:"error"`
	Powerpack
}

type PowerpackUpdateResponse struct {
	PowerpackCreateResponseBody
}

//powerpack list
type PowerpackList struct {
	BaseListPPKResponse
	Objects []Powerpack `json:"objects" url:"objects"`
}

type PowerpackListParams struct {
	Limit   int    `url:"limit,omitempty"`
	Offset  int    `url:"offset,omitempty"`
	Service string `url:"service,omitempty"`
}
type PowerpackSearchParam struct {
	Starts_with  string `json:"starts_with,omitempty" url:"starts_with,omitempty"`
	Country_iso2 string `json:"country_iso2,omitempty" url:"country_iso2,omitempty"`
	Type         string `json:"type,omitempty" url:"type,omitempty"`
	Limit        string `json:"limit,omitempty" url:"limit,omitempty"`
	Offset       string `json:"offset,omitempty" url:"offset,omitempty"`
	Service      string `json:"service,omitempty" url:"service,omitempty"`
}
type PowerpackPhoneResponseBody struct {
	BaseListPPKResponse
	Objects []NumberPoolResponse `json:"objects" url:"objects"`
}

type PPKMeta struct {
	Previous   *string
	Next       *string
	TotalCount int `json:"total_count" url:"api_id"`
	Offset     int
	Limit      int
}

type BaseListPPKResponse struct {
	ApiID string  `json:"api_id" url:"api_id"`
	Meta  PPKMeta `json:"meta" url:"meta"`
}

type BuyPhoneNumberParam struct {
	Number       string `json:"number,omitempty"`
	Country_iso2 string `json:"country_iso,omitempty"`
	Type         string `json:"type,omitempty"`
	Region       string `json:"region,omitempty"`
	Pattern      string `json:"pattern,omitempty"`
	Service      string `json:"service,omitempty"`
}

type RentNumber struct {
	Rent    string `json:"rent,omitempty"`
	Service string `json:"service,omitempty"`
}

// returns the List.. of all powerpack
func (service *PowerpackService) List(params PowerpackListParams) (response *PowerpackList, err error) {
	req, err := service.client.NewRequest("GET", params, "Powerpack/")
	if err != nil {
		return
	}
	response = &PowerpackList{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *PowerpackService) Get(powerpackUUID string) (response *PowerpackService, err error) {
	req, err := service.client.NewRequest("GET", nil, "Powerpack/%s/", powerpackUUID)
	if err != nil {
		return
	}
	resp := &PowerpackCreateResponseBody{}
	err = service.client.ExecuteRequest(req, resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	response = &PowerpackService{}
	response.client = service.client
	response.Powerpack = resp.Powerpack
	return
}

func (service *PowerpackService) Create(params PowerackCreateParams) (response *PowerpackCreateResponseBody, err error) {
	req, err := service.client.NewRequest("POST", params, "Powerpack")
	if err != nil {
		return
	}
	response = &PowerpackCreateResponseBody{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *PowerpackService) Update(params PowerackUpdateParams) (response *PowerpackUpdateResponse, err error) {
	uuid := service.Powerpack.UUID
	req, err := service.client.NewRequest("POST", params, "Powerpack/%s", uuid)
	if err != nil {
		return
	}
	response = &PowerpackUpdateResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *PowerpackService) Delete(params PowerpackDeleteParams) (response *PowerpackDeleteResponse, err error) {
	uuid := service.Powerpack.UUID
	req, err := service.client.NewRequest("DELETE", params, "Powerpack/%s", uuid)
	if err != nil {
		return
	}
	response = &PowerpackDeleteResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *PowerpackService) List_numbers(params PowerpackSearchParam) (response *PowerpackPhoneResponseBody, err error) {
	numberpool_path := service.Powerpack.NumberPoolUUID
	uriSegments := strings.Split(numberpool_path, "/")
	req, err := service.client.NewRequest("GET", params, "NumberPool/%s/Number/", uriSegments[5])
	if err != nil {
		return
	}
	response = &PowerpackPhoneResponseBody{}
	err = service.client.ExecuteRequest(req, response)
	return

}

func (service *PowerpackService) Count_numbers(params PowerpackSearchParam) (count int, err error) {
	numberpool_path := service.Powerpack.NumberPoolUUID
	uriSegments := strings.Split(numberpool_path, "/")
	req, err := service.client.NewRequest("GET", params, "NumberPool/%s/Number/", uriSegments[5])
	if err != nil {
		return
	}
	response := &PowerpackPhoneResponseBody{}
	err = service.client.ExecuteRequest(req, response)
	count = response.BaseListPPKResponse.Meta.TotalCount
	return count, nil
}

func (service *PowerpackService) Find_numbers(number string) (response *NumberResponse, err error) {
	params := PowerpackFindNumberOptions{}
	response, err = service.FindNumbersWithOptions(number, params)
	return
}

func (service *PowerpackService) FindNumbersWithOptions(number string, params PowerpackFindNumberOptions) (response *NumberResponse, err error) {
	numberpool_path := service.Powerpack.NumberPoolUUID
	uriSegments := strings.Split(numberpool_path, "/")
	req, err := service.client.NewRequest("GET", params, "NumberPool/%s/Number/%s/", uriSegments[5], number)
	if err != nil {
		return
	}
	response = &NumberResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *PowerpackService) Add_number(number string) (response *NumberResponse, err error) {
	params := PowerpackAddNumberOptions{}
	response, err = service.AddNumberWithOptions(number, params)
	return
}

func (service *PowerpackService) AddNumberWithOptions(number string, params PowerpackAddNumberOptions) (response *NumberResponse, err error) {
	numberpool_path := service.Powerpack.NumberPoolUUID
	uriSegments := strings.Split(numberpool_path, "/")
	req, err := service.client.NewRequest("POST", params, "NumberPool/%s/Number/%s", uriSegments[5], number)
	if err != nil {
		return
	}
	response = &NumberResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *PowerpackService) Add_tollfree(tollfree string) (response *NumberResponse, err error) {
	params := PowerpackAddTollfreeOptions{}
	response, err = service.AddTollfreeWithOptions(tollfree, params)
	return
}

func (service *PowerpackService) AddTollfreeWithOptions(tollfree string, params PowerpackAddTollfreeOptions) (response *NumberResponse, err error) {
	numberpoolUUID := service.Powerpack.NumberPoolUUID
	uriSegments := strings.Split(numberpoolUUID, "/")
	req, err := service.client.NewRequest("POST", params, "NumberPool/%s/Tollfree/%s", uriSegments[5], tollfree)
	if err != nil {
		return
	}
	response = &NumberResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *PowerpackService) Remove_number(number string, param NumberRemoveParams) (response *NumberDeleteResponse, err error) {
	numberpoolUUID := service.Powerpack.NumberPoolUUID
	uriSegments := strings.Split(numberpoolUUID, "/")
	req, err := service.client.NewRequest("DELETE", param, "NumberPool/%s/Number/%s", uriSegments[5], number)
	if err != nil {
		return
	}
	response = &NumberDeleteResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *PowerpackService) Remove_tollfree(tollfree string, param NumberRemoveParams) (response *TollfreeDeleteResponse, err error) {
	numberpoolUUID := service.Powerpack.NumberPoolUUID
	uriSegments := strings.Split(numberpoolUUID, "/")
	req, err := service.client.NewRequest("DELETE", param, "NumberPool/%s/Tollfree/%s", uriSegments[5], tollfree)
	if err != nil {
		return
	}
	response = &TollfreeDeleteResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *PowerpackService) Remove_shortcode(shortcode string) (response *ShortcodeDeleteResponse, err error) {
	numberpoolUUID := service.Powerpack.NumberPoolUUID
	uriSegments := strings.Split(numberpoolUUID, "/")
	req, err := service.client.NewRequest("DELETE", nil, "NumberPool/%s/Shortcode/%s", uriSegments[5], shortcode)
	if err != nil {
		return
	}
	response = &ShortcodeDeleteResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *PowerpackService) List_shortcodes() (response *ShortCodeResponse, err error) {
	params := PowerpackListShortcodeOptions{}
	response, err = service.ListShortcodesWithOptions(params)
	return
}

func (service *PowerpackService) ListShortcodesWithOptions(params PowerpackListShortcodeOptions) (response *ShortCodeResponse, err error) {
	numberpool_path := service.Powerpack.NumberPoolUUID
	uriSegments := strings.Split(numberpool_path, "/")
	req, err := service.client.NewRequest("GET", params, "NumberPool/%s/Shortcode", uriSegments[5])
	if err != nil {
		return
	}
	response = &ShortCodeResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *PowerpackService) List_tollfree() (response *TollfreeResponse, err error) {
	params := PowerpackListTollfreeOptions{}
	response, err = service.ListTollfreeWithOptions(params)
	return
}

func (service *PowerpackService) ListTollfreeWithOptions(params PowerpackListTollfreeOptions) (response *TollfreeResponse, err error) {
	numberpoolUUID := service.Powerpack.NumberPoolUUID
	uriSegments := strings.Split(numberpoolUUID, "/")
	req, err := service.client.NewRequest("GET", params, "NumberPool/%s/Tollfree", uriSegments[5])
	if err != nil {
		return
	}
	response = &TollfreeResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *PowerpackService) Find_shortcode(shortcode string) (response *FindShortCodeResponse, err error) {
	params := PowerpackFindShortcodeOptions{}
	response, err = service.FindShortcodeWithOptions(shortcode, params)
	return
}

func (service *PowerpackService) FindShortcodeWithOptions(shortcode string, params PowerpackFindShortcodeOptions) (response *FindShortCodeResponse, err error) {
	numberpool_path := service.Powerpack.NumberPoolUUID
	uriSegments := strings.Split(numberpool_path, "/")
	req, err := service.client.NewRequest("GET", params, "NumberPool/%s/Shortcode/%s/", uriSegments[5], shortcode)
	if err != nil {
		return
	}
	response = &FindShortCodeResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *PowerpackService) Find_tollfree(tollfree string) (response *FindTollfreeResponse, err error) {
	params := PowerpackFindTollfreeOptions{}
	response, err = service.FindTollfreeWithOptions(tollfree, params)
	return
}

func (service *PowerpackService) FindTollfreeWithOptions(tollfree string, params PowerpackFindTollfreeOptions) (response *FindTollfreeResponse, err error) {
	numberpoolUUID := service.Powerpack.NumberPoolUUID
	uriSegments := strings.Split(numberpoolUUID, "/")
	req, err := service.client.NewRequest("GET", params, "NumberPool/%s/Tollfree/%s/", uriSegments[5], tollfree)
	if err != nil {
		return
	}
	response = &FindTollfreeResponse{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *PowerpackService) Buy_add_number(phoneParam BuyPhoneNumberParam) (response *NumberResponse, err error) {
	numberpool_path := service.Powerpack.NumberPoolUUID
	uriSegments := strings.Split(numberpool_path, "/")
	payload := RentNumber{
		Rent: "true",
	}

	payload.Service = phoneParam.Service
	number := phoneParam.Number
	if number != "" {
		req, err := service.client.NewRequest("POST", payload, "NumberPool/%s/Number/%s", uriSegments[5], number)
		if err != nil {
			panic(err)
		}
		response = &NumberResponse{}
		err = service.client.ExecuteRequest(req, response)
	} else {
		Type := phoneParam.Type
		region := phoneParam.Region
		pattern := phoneParam.Pattern
		countryiso := phoneParam.Country_iso2
		serviceType := phoneParam.Service
		params := PhoneNumberListParams{
			Type:       Type,
			Region:     region,
			Pattern:    pattern,
			CountryISO: countryiso,
			Services:   serviceType,
		}
		responsephoneNo, err := service.client.PhoneNumbers.List(params)
		if err != nil {
			panic(err)
		}
		if len(responsephoneNo.Objects) < 1 {
			response = &NumberResponse{}
			return response, nil
		}

		req, err := service.client.NewRequest("POST", payload, "NumberPool/%s/Number/%s", uriSegments[5], responsephoneNo.Objects[0].Number)
		if err != nil {
			panic(err)
		}
		response = &NumberResponse{}
		err = service.client.ExecuteRequest(req, response)

	}
	return
}
