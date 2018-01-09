package plivo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplicationService_Create(t *testing.T) {
	expectResponse("applicationCreateResponse.json", 202)

	if _, err := client.Applications.Create(ApplicationCreateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Applications.Create(ApplicationCreateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Application")
}

func TestApplicationService_Get(t *testing.T) {
	expectResponse("applicationGetResponse.json", 202)
	appId := "appId"

	application, err := client.Applications.Get(appId)
	assert.Equal(t, application.ID(), application.AppID)
	if err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err = client.Applications.Get(appId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Application/%s", appId)
}

func TestApplicationService_Update(t *testing.T) {
	expectResponse("applicationModifyResponse.json", 202)
	appId := "appId"

	if _, err := client.Applications.Update(appId, ApplicationUpdateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Applications.Update(appId, ApplicationUpdateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Application/%s", appId)
}

func TestApplicationService_Delete(t *testing.T) {
	expectResponse("", 204)
	appId := "appId"

	if err := client.Applications.Delete(appId); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	err := client.Applications.Delete(appId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Application/%s", appId)
}

func TestApplicationService_List(t *testing.T) {
	expectResponse("applicationListResponse.json", 200)

	if _, err := client.Applications.List(ApplicationListParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Applications.List(ApplicationListParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Application")
}
