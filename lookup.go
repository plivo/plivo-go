package plivo

import (
	"encoding/json"
	"errors"
)

// NOTE: All of Plivo's APIs are in a single Go package. Unfortunately,
// this imposes the limitation that each struct type has to be unique across
// all of Plivo's product APIs.

type Country struct {
	Name     string `json:"name"`
	CodeISO2 string `json:"code_iso2"`
	CodeISO3 string `json:"code_iso3"`
}

type NumberFormat struct {
	E164          string `json:"e164"`
	National      string `json:"national"`
	International string `json:"international"`
	RFC3966       string `json:"rfc3966"`
}

type Carrier struct {
	MobileCountryCode string `json:"mobile_country_code"`
	MobileNetworkCode string `json:"mobile_network_code"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	Ported            bool   `json:"ported"`
}

// LookupResponse is the success response returned by Plivo Lookup API.
type LookupResponse struct {
	ApiID       string        `json:"api_id"`
	PhoneNumber string        `json:"phone_number"`
	Country     *Country      `json:"country"`
	Format      *NumberFormat `json:"format"`
	Carrier     *Carrier      `json:"carrier"`
	ResourceURI string        `json:"resource_uri"`
}

type LookupService struct {
	client *Client
}

// LookupParams is the input parameters for Plivo Lookup API.
type LookupParams struct {
	Type string `url:"type"`
}

// Get looks up a phone number using Plivo Lookup API.
func (s *LookupService) Get(number string, params LookupParams) (*LookupResponse, error) {
	if params.Type == "" {
		return nil, errors.New("Type must be set in params")
	}

	req, err := s.client.BaseClient.NewRequest("GET", params, "v1/Lookup/Number/%s", number)
	if err != nil {
		return nil, err
	}

	resp := new(LookupResponse)
	if err := s.client.ExecuteRequest(req, resp); err != nil {
		return nil, s.newError(err.Error())
	}

	return resp, nil
}

// LookupError is the error response returned by Plivo Lookup API.
type LookupError struct {
	respBody  string
	ApiID     string `json:"api_id"`
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

// Error returns the raw json/text response from Lookup API.
func (e *LookupError) Error() string {
	return e.respBody
}

func (s *LookupService) newError(body string) *LookupError {
	resp := &LookupError{
		respBody: body,
	}

	_ = json.Unmarshal([]byte(body), resp)
	return resp
}
