package plivo

type TokenService struct {
	client *Client
}

type Token struct {
	Iss            string `json:"iss,omitempty" url:"iss,omitempty"`
	Sub            string `json:"sub,omitempty" url:"sub,omitempty"`
	Exp            int64  `json:"exp,omitempty" url:"exp,omitempty"`
	Nbf            int64  `json:"nbf,omitempty" url:"nbf,omitempty"`
	Incoming_allow bool   `json:"incoming_allow,omitempty" url:"incoming_allow,omitempty"`
	Outgoing_allow bool   `json:"outgoing_allow,omitempty" url:"outgoing_allow,omitempty"`
	App            string `json:"app,omitempty" url:"app,omitempty"`
}

type TokenCreateParams struct {
	// Required parameters.
	Iss string `json:"iss,omitempty" url:"iss,omitempty"`
	// Optional parameters.
	Sub            string `json:"sub,omitempty" url:"sub,omitempty"`
	Exp            int64  `json:"exp,omitempty" url:"exp,omitempty"`
	Nbf            int64  `json:"nbf,omitempty" url:"nbf,omitempty"`
	Incoming_allow bool   `json:"incoming_allow,omitempty" url:"incoming_allow,omitempty"`
	Outgoing_allow bool   `json:"outgoing_allow,omitempty" url:"outgoing_allow,omitempty"`
	App            string `json:"app,omitempty" url:"app,omitempty"`
}

// Stores response for creating a token.
type TokenCreateResponse struct {
	Token string `json:"token" url:"token"`
	ApiID string `json:"api_id" url:"api_id"`
}

func (service *TokenService) Create(params TokenCreateParams) (response *TokenCreateResponse, err error) {
	per := make(map[string]interface{})
	per["voice"] = make(map[string]interface{})

	if params.incoming_allow {
		per["voice"].(map[string]interface{})["incoming_allow"] = params.incoming_allow
	}

	if params.Outgoing_allow {
		per["voice"].(map[string]interface{})["outgoing_allow"] = params.Outgoing_allow
	}

	req, err := service.client.NewRequest("POST", params, "JWT/Token")
	if err != nil {
		return
	}

	response = &TokenCreateResponse{}
	err = service.client.ExecuteRequest(req, response, isVoiceRequest())
	return
}
