package plivo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTollFreeRequestVerificationServiceCreate(t *testing.T) {
	expectResponse("TollFreeRequestVerificationCreateResponse.json", 201)

	if _, err := client.TollFreeRequestVerification.Create(TollfreeVerificationCreateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.TollFreeRequestVerification.Create(TollfreeVerificationCreateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "TollfreeVerification")
}

func TestTollFreeRequestVerificationServiceUpdate(t *testing.T) {
	expectResponse("TollFreeRequestVerificationUpdateResponse.json", 202)
	RequestId := "RequestId"

	if _, err := client.TollFreeRequestVerification.Update(RequestId, TollfreeVerificationUpdateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.TollFreeRequestVerification.Update(RequestId, TollfreeVerificationUpdateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "TollfreeVerification/%s", RequestId)
}

func TestTollFreeRequestVerificationServiceList(t *testing.T) {
	expectResponse("TollFreeRequestVerificationListResponse.json", 200)

	if _, err := client.TollFreeRequestVerification.List(TollfreeVerificationListParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.TollFreeRequestVerification.List(TollfreeVerificationListParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "TollfreeVerification")
}

func TestTollFreeRequestVerificationServiceGet(t *testing.T) {
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

func TestTollFreeRequestVerificationServiceDelete(t *testing.T) {
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
