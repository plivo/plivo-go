package plivo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndpointService_Create(t *testing.T) {
	expectResponse("endpointCreateResponse.json", 202)

	if _, err := client.Endpoints.Create(EndpointCreateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Endpoints.Create(EndpointCreateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Endpoint")
}

func TestEndpointService_Update(t *testing.T) {
	expectResponse("endpointUpdateResponse.json", 202)
	endpointId := "endpointId"

	if _, err := client.Endpoints.Update(endpointId, EndpointUpdateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Endpoints.Update(endpointId, EndpointUpdateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Endpoint/%s", endpointId)
}

func TestEndpointService_List(t *testing.T) {
	expectResponse("endpointListResponse.json", 200)

	if _, err := client.Endpoints.List(EndpointListParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Endpoints.List(EndpointListParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Endpoint")
}

func TestEndpointService_Get(t *testing.T) {
	expectResponse("endpointGetResponse.json", 200)
	endpointId := "endpointId"

	endpoint, err := client.Endpoints.Get(endpointId)
	assert.Equal(t, endpoint.ID(), endpoint.EndpointID)
	if err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err = client.Endpoints.Get(endpointId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Endpoint/%s", endpointId)
}
func TestEndpointService_Delete(t *testing.T) {
	expectResponse("", 204)
	endpointId := "endpointId"

	if err := client.Endpoints.Delete(endpointId); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	err := client.Endpoints.Delete(endpointId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Endpoint/%s", endpointId)
}
