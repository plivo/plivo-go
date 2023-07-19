package plivo

import "time"

type VerifyService struct {
	client *Client
	Session
}

type SessionList struct {
	APIID string `json:"api_id"`
	Meta  struct {
		Limit    int     `json:"limit"`
		Offset   int     `json:"offset"`
		Next     *string `json:"next"`
		Previous *string `json:"previous"`
	} `json:"meta"`
	Sessions []Session `json:"sessions"`
}

type Session struct {
	APIID              string           `json:"api_id,omitempty"`
	SessionUUID        string           `json:"session_uuid,omitempty"`
	AppUUID            string           `json:"app_uuid,omitempty"`
	Alias              string           `json:"alias,omitempty"`
	Recipient          string           `json:"recipient,omitempty"`
	Channel            string           `json:"channel,omitempty"`
	Status             string           `json:"status,omitempty"`
	Count              int              `json:"count,omitempty"`
	RequesterIP        string           `json:"requestor_ip,omitempty"`
	CountryISO         string           `json:"destination_country_iso2,omitempty"`
	DestinationNetwork *string          `json:"destination_network,omitempty"`
	AttemptDetails     []AttemptDetails `json:"attempt_details,omitempty"`
	Charges            Charges          `json:"charges,omitempty"`
	CreatedAt          time.Time        `json:"created_at,omitempty"`
	UpdatedAt          time.Time        `json:"updated_at,omitempty"`
}

type AttemptDetails struct {
	Channel     string    `json:"channel,omitempty"`
	AttemptUUID string    `json:"attempt_uuid,omitempty"`
	Status      string    `json:"status,omitempty"`
	Time        time.Time `json:"time,omitempty"`
}

type Charges struct {
	TotalCharge      string           `json:"total_charge,omitempty"`
	ValidationCharge string           `json:"validation_charge,omitempty"`
	AttemptCharges   []AttemptCharges `json:"attempt_charges,omitempty"`
}

type AttemptCharges struct {
	AttemptUUID string `json:"attempt_uuid,omitempty"`
	Channel     string `json:"channel,omitempty"`
	Charge      string `json:"charge,omitempty"`
}

type SessionCreateParams struct {
	Recipient string `json:"recipient,omitempty"`
	// Optional parameters.
	AppUUID string `json:"app_uuid,omitempty"`
	Channel string `json:"channel,omitempty"`
	URL     string `json:"url,omitempty"`
	Method  string `json:"method,omitempty"`
	Src     string `json:"src,omitempty"`
}

type SessionCreateResponseBody struct {
	APIID       string `json:"api_id,omitempty"`
	Error       string `json:"error,omitempty"`
	Message     string `json:"message,omitempty"`
	SessionUUID string `json:"session_uuid,omitempty"`
}

func (service *VerifyService) Create(params SessionCreateParams) (response *SessionCreateResponseBody, err error) {
	req, err := service.client.NewRequest("POST", params, "Verify/Session")
	if err != nil {
		return
	}
	response = &SessionCreateResponseBody{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *VerifyService) Get(sessionUUID string) (response *Session, err error) {
	req, err := service.client.NewRequest("GET", nil, "Verify/Session/%s", sessionUUID)
	if err != nil {
		return
	}
	response = &Session{}
	err = service.client.ExecuteRequest(req, response)
	return
}

func (service *VerifyService) ListSessions() (response *SessionList, err error) {
	req, err := service.client.NewRequest("GET", nil, "Verify/Session")
	if err != nil {
		return
	}
	response = &SessionList{}
	err = service.client.ExecuteRequest(req, response)
	return
}
