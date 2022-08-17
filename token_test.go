package plivo

import (
	"errors"
	_ "github.com/stretchr/testify/require"
	"testing"
)

func TestTokenService_Create(t *testing.T) {
	expectResponse("TokenCreateResponse.json", 202)

	if _, err := client.Token.Create(TokenCreateParams{Iss: "MAMDVLZJY2ZGY5MWU1ZJ", Sub: "kowshik"}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Token.Create(TokenCreateParams{Iss: "MAMDVLZJY2ZGY5MWU1Z", Sub: "kowshik"})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "JWT/Token")
}
