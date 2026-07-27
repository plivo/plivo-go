package plivo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTollFreeRequestVerificationServiceCreate(t *testing.T) {
	expectResponse("TollFreeRequestVerificationCreateResponse.json", 201)

	createParams := TollfreeVerificationCreateParams{
		TermsAndConditionsLink: "https://example.com/terms",
		PrivacyPolicyLink:      "https://example.com/privacy",
		OptinMessage:           "Reply YES to opt in",
		HelpMessage:            "Reply HELP for help",
	}
	if _, err := client.TollFreeRequestVerification.Create(createParams); err != nil {
		panic(err)
	}

	body := string(requestBody)
	assert.Contains(t, body, `"terms_and_conditions_link":"https://example.com/terms"`)
	assert.Contains(t, body, `"privacy_policy_link":"https://example.com/privacy"`)
	assert.Contains(t, body, `"optin_message":"Reply YES to opt in"`)
	assert.Contains(t, body, `"help_message":"Reply HELP for help"`)

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

	updateParams := TollfreeVerificationUpdateParams{
		TermsAndConditionsLink: "https://example.com/terms",
		PrivacyPolicyLink:      "https://example.com/privacy",
		OptinMessage:           "Reply YES to opt in",
		HelpMessage:            "Reply HELP for help",
	}
	if _, err := client.TollFreeRequestVerification.Update(RequestId, updateParams); err != nil {
		panic(err)
	}

	body := string(requestBody)
	assert.Contains(t, body, `"terms_and_conditions_link":"https://example.com/terms"`)
	assert.Contains(t, body, `"privacy_policy_link":"https://example.com/privacy"`)
	assert.Contains(t, body, `"optin_message":"Reply YES to opt in"`)
	assert.Contains(t, body, `"help_message":"Reply HELP for help"`)

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
