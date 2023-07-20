package plivo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifyService_List(t *testing.T) {
	expectResponse("verifySessionListResponse.json", 200)
	assert := require.New(t)
	resp, err := client.Verify.List(SessionListParams{})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.Sessions[0].SessionUUID)
	assert.NotNil(resp.Sessions)
	assert.NotNil(resp.Meta)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Verify.List(SessionListParams{})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl
	assertRequest(t, "GET", "Verify/Session")
}

func TestVerifyService_Get(t *testing.T) {
	expectResponse("verifySessionGetResponse.json", 200)
	uuid := "4124e518-a8c9-4feb-8cff-d86636ba9234"
	assert := require.New(t)
	resp, err := client.Verify.Get(uuid)
	assert.NotNil(resp)
	assert.Nil(err)
	assert.Equal(resp.SessionUUID, uuid)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Verify.Get(uuid)
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "GET", "Verify/Session/%s", uuid)
}

func TestVerifyService_Create(t *testing.T) {
	expectResponse("verifySessionSendResponse.json", 202)
	assert := require.New(t)
	resp, err := client.Verify.Create(SessionCreateParams{
		Recipient: "9089789099",
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.APIID)
	assert.NotEmpty(resp.SessionUUID)
	assert.Equal(resp.Error, "")
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Verify.Create(SessionCreateParams{
		Recipient: "9089789099",
	})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "Verify/Session")
}

func TestVerifyService_Validate(t *testing.T) {
	expectResponse("verifySessionValidateResponse.json", 200)
	assert := require.New(t)
	resp, err := client.Verify.Validate(SessionValidationParams{
		OTP: "9089789",
	}, "5b40a428-bfc7-4daf-9d06-726c558bf3b8")
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.APIID)
	assert.Equal(resp.Error, "")
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Verify.Validate(SessionValidationParams{
		OTP: "9089789",
	}, "5b40a428-bfc7-4daf-9d06-726c558bf3b8")
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "Verify/Session/%s", "5b40a428-bfc7-4daf-9d06-726c558bf3b8")
}
