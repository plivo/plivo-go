package plivo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifySessionService_Create(t *testing.T) {
	expectResponse("verifySessionCreateResponse.json", 200)
	assert := require.New(t)
	resp, err := client.VerifySessions.Create(VerifySessionCreateParams{
		BrandName:  "TestBrand",
		CodeLength: 6,
		Text:       "Your OTP is <otp>",
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)

	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.VerifySessions.Create(VerifySessionCreateParams{
		BrandName:  "TestBrand",
		CodeLength: 6,
		Text:       "Your OTP is <otp>",
	})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "Verify/Session")
}