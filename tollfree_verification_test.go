package plivo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTollFreeRequestVerificationService_Create(t *testing.T) {
	expectResponse("TollFreeRequestVerificationCreateResponse.json", 201)

	if _, err := client.TollFreeRequestVerification.Create(TollFreeCreateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.TollFreeRequestVerification.Create(TollFreeCreateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "TollfreeVerification")
}

func TestTollFreeRequestVerificationService_Update(t *testing.T) {
	expectResponse("TollFreeRequestVerificationUpdateResponse.json", 202)
	RequestId := "RequestId"

	if _, err := client.TollFreeRequestVerification.Update(RequestId, TollFreeUpdateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.TollFreeRequestVerification.Update(RequestId, TollFreeUpdateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "TollfreeVerification/%s", RequestId)
}

func TestTollFreeRequestVerificationService_List(t *testing.T) {
	expectResponse("TollFreeRequestVerificationListResponse.json", 200)

	if _, err := client.TollFreeRequestVerification.List(TollFreeListParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.TollFreeRequestVerification.List(TollFreeListParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "TollfreeVerification")
}

func TestTollFreeRequestVerificationService_Get(t *testing.T) {
	expectResponse("TollFreeRequestVerificationGetResponse.json", 200)
	RequestId := "RequestId"

	TollFreeRequestVerificationRequest, err := client.TollFreeRequestVerification.Get(RequestId)
	assert.Equal(t, TollFreeRequestVerificationRequest.ID(), TollFreeRequestVerificationRequest.UUID)
	if err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err = client.TollFreeRequestVerification.Get(RequestId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "TollfreeVerification/%s", RequestId)
}

func TestTollFreeRequestVerificationService_Delete(t *testing.T) {
	expectResponse("", 204)
	RequestId := "RequestId"

	if err := client.TollFreeRequestVerification.Delete(RequestId); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	err := client.TollFreeRequestVerification.Delete(RequestId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "TollfreeVerification/%s", RequestId)
}
