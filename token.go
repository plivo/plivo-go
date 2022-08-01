package plivo

type TokenService struct {
	client *Client
}

type Token struct {
	iss            string `json:"iss,omitempty" url:"iss,omitempty"`
	sub            string `json:"sub,omitempty" url:"sub,omitempty"`
	exp            int64  `json:"exp,omitempty" url:"exp,omitempty"`
	nbf            int64  `json:"nbf,omitempty" url:"nbf,omitempty"`
	incoming_allow bool   `json:"incoming_allow,omitempty" url:"incoming_allow,omitempty"`
	outgoing_allow bool   `json:"outgoing_allow,omitempty" url:"outgoing_allow,omitempty"`
	app            string `json:"app,omitempty" url:"app,omitempty"`
}

type TokenCreateParams struct {
	// Required parameters.
	iss string `json:"iss,omitempty" url:"iss,omitempty"`
	// Optional parameters.
	sub            string `json:"sub,omitempty" url:"sub,omitempty"`
	exp            int64  `json:"exp,omitempty" url:"exp,omitempty"`
	nbf            int64  `json:"nbf,omitempty" url:"nbf,omitempty"`
	incoming_allow bool   `json:"incoming_allow,omitempty" url:"incoming_allow,omitempty"`
	outgoing_allow bool   `json:"outgoing_allow,omitempty" url:"outgoing_allow,omitempty"`
	app            string `json:"app,omitempty" url:"app,omitempty"`
}

// Stores response for creating a token.
type TokenCreateResponse struct {
	Token string `json:"token" url:"token"`
	ApiID string `json:"api_id" url:"api_id"`
}

func (service *TokenService) Create(params TokenCreateParams) (response *TokenCreateResponse, err error) {
	//if incoming_allow is not empty add it to params

	req, err := service.client.NewRequest("POST", params, "JWT/Token")
	if err != nil {
		return
	}

	response = &TokenCreateResponse{}
	err = service.client.ExecuteRequest(req, response, isVoiceRequest())
	return
}
