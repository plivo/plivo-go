package plivo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifyService_List(t *testing.T) {
	expectResponse("verifySessionListResponse.json", 200)
	assert := require.New(t)
	resp, err := client.VerifySession.List(SessionListParams{})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.Sessions[0].SessionUUID)
	assert.NotNil(resp.Sessions)
	assert.NotNil(resp.Meta)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.VerifySession.List(SessionListParams{})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl
	assertRequest(t, "GET", "Verify/Session")
}

func TestVerifyService_Get(t *testing.T) {
	expectResponse("verifySessionGetResponse.json", 200)
	uuid := "4124e518-a8c9-4feb-8cff-d86636ba9234"
	assert := require.New(t)
	resp, err := client.VerifySession.Get(uuid)
	assert.NotNil(resp)
	assert.Nil(err)
	assert.Equal(resp.SessionUUID, uuid)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.VerifySession.Get(uuid)
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "GET", "Verify/Session/%s", uuid)
}

func TestVerifyService_Create(t *testing.T) {
	expectResponse("verifySessionSendResponse.json", 202)
	assert := require.New(t)
	resp, err := client.VerifySession.Create(SessionCreateParams{
		Recipient: "9089789099",
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.APIID)
	assert.NotEmpty(resp.SessionUUID)
	assert.Equal(resp.Error, "")
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.VerifySession.Create(SessionCreateParams{
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
	resp, err := client.VerifySession.Validate(SessionValidationParams{
		OTP: "9089789",
	}, "5b40a428-bfc7-4daf-9d06-726c558bf3b8")
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.APIID)
	assert.Equal(resp.Error, "")
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.VerifySession.Validate(SessionValidationParams{
		OTP: "9089789",
	}, "5b40a428-bfc7-4daf-9d06-726c558bf3b8")
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "Verify/Session/%s", "5b40a428-bfc7-4daf-9d06-726c558bf3b8")
}

func TestVerifyService_Get_INRFields(t *testing.T) {
	expectResponse("verifySessionGetInrResponse.json", 200)
	assert := require.New(t)
	resp, err := client.VerifySession.Get("4124e518-a8c9-4feb-8cff-d86636ba9234")
	assert.Nil(err)
	assert.Equal("0.00500", resp.Charges.TotalCharge)
	assert.Equal("0.41500", resp.Charges.TotalChargeINR)
	assert.Len(resp.Charges.AttemptCharges, 1)
	assert.Equal("0.00500", resp.Charges.AttemptCharges[0].Charge)
	assert.Equal("0.41500", resp.Charges.AttemptCharges[0].ChargeINR)
	assert.Len(resp.AttemptDetails, 1)
	assert.NotNil(resp.AttemptDetails[0].Amount)
	assert.InDelta(0.005, *resp.AttemptDetails[0].Amount, 1e-9)
	assert.NotNil(resp.AttemptDetails[0].AmountINR)
	assert.InDelta(0.415, *resp.AttemptDetails[0].AmountINR, 1e-9)
}

func TestVerifyService_List_INRFields(t *testing.T) {
	expectResponse("verifySessionListInrResponse.json", 200)
	assert := require.New(t)
	resp, err := client.VerifySession.List(SessionListParams{})
	assert.Nil(err)
	assert.Len(resp.Sessions, 1)
	session := resp.Sessions[0]
	assert.Equal("0.41500", session.Charges.TotalChargeINR)
	assert.Equal("0.41500", session.Charges.AttemptCharges[0].ChargeINR)
	assert.NotNil(session.AttemptDetails[0].Amount)
	assert.InDelta(0.005, *session.AttemptDetails[0].Amount, 1e-9)
	assert.NotNil(session.AttemptDetails[0].AmountINR)
	assert.InDelta(0.415, *session.AttemptDetails[0].AmountINR, 1e-9)
}

func TestVerifyService_List_USDAccountHasNoINRFields(t *testing.T) {
	expectResponse("verifySessionListResponse.json", 200)
	assert := require.New(t)
	resp, err := client.VerifySession.List(SessionListParams{})
	assert.Nil(err)
	session := resp.Sessions[0]
	assert.NotEmpty(session.AttemptDetails)
	assert.NotEmpty(session.Charges.AttemptCharges)
	assert.Empty(session.Charges.TotalChargeINR)
	assert.Empty(session.Charges.AttemptCharges[0].ChargeINR)
	assert.NotNil(session.AttemptDetails[0].Amount)
	assert.InDelta(0.0, *session.AttemptDetails[0].Amount, 1e-9)
	assert.Nil(session.AttemptDetails[0].AmountINR)
}
