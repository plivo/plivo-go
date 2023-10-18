package plivo

import (
	"errors"
	"testing"
)

func TestVerifyCallerIdService_InitiateVerify(t *testing.T) {
	expectResponse("initiateVerifyResponse.json", 200)

	if _, err := client.VerifyCallerId.InitiateVerify(InitiateVerify{
		PhoneNumber: "917708772011",
		Alias:       "test",
		Channel:     "call",
	}); err != nil {
		panic(err)
	}
	cl := client.httpClient
	client.httpClient = nil
	_, err := client.VerifyCallerId.InitiateVerify(InitiateVerify{
		PhoneNumber: "917708772011",
		Alias:       "test",
		Channel:     "call",
	})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "VerifiedCallerId")
}

func TestVerifyCallerIdService_VerifyCallerID(t *testing.T) {
	expectResponse("verifyCallerIDResponse.json", 200)

	if _, err := client.VerifyCallerId.VerifyCallerID("eeab1477-e59b-4821-9e61-fd5847c2a5db", "610534"); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.VerifyCallerId.VerifyCallerID("eeab1477-e59b-4821-9e61-fd5847c2a5db", "610534")
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "VerifiedCallerId/Verification/eeab1477-e59b-4821-9e61-fd5847c2a5db")
}

func TestVerifyCallerIdService_DeleteVerifiedCallerID(t *testing.T) {
	expectResponse("", 204)

	if err := client.VerifyCallerId.DeleteVerifiedCallerID("917708772010"); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	err := client.VerifyCallerId.DeleteVerifiedCallerID("917708772010")
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "VerifiedCallerId/917708772010")
}

func TestVerifyCallerIdService_GetVerifiedCallerID(t *testing.T) {
	expectResponse("getVerifiedCallerIDResponse.json", 200)

	if _, err := client.VerifyCallerId.GetVerifiedCallerID("917708772010"); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.VerifyCallerId.GetVerifiedCallerID("917708772010")
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "VerifiedCallerId/917708772010")
}

func TestVerifyCallerIdService_UpdateVerifiedCallerID(t *testing.T) {
	expectResponse("updateVerifiedCallerIDResponse.json", 204)

	if _, err := client.VerifyCallerId.UpdateVerifiedCallerID("917708772010", UpdateVerifiedCallerIDParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.VerifyCallerId.UpdateVerifiedCallerID("917708772010", UpdateVerifiedCallerIDParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "VerifiedCallerId/917708772010")
}

func TestVerifyCallerIdService_ListVerifiedCallerID(t *testing.T) {
	expectResponse("listVerifiedCallerIDResponse.json", 204)

	if _, err := client.VerifyCallerId.ListVerifiedCallerID(ListVerifiedCallerIdParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.VerifyCallerId.ListVerifiedCallerID(ListVerifiedCallerIdParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "VerifiedCallerId")
}
