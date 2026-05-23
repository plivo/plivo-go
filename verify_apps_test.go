package plivo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifyAppService_Create(t *testing.T) {
	expectResponse("verifyAppCreateResponse.json", 201)
	assert := require.New(t)
	resp, err := client.VerifyApps.Create(VerifyAppCreateParams{
		Name: "Test App",
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	assert.NotEmpty(resp.AppUUID)
	assert.Equal(resp.Error, "")

	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.VerifyApps.Create(VerifyAppCreateParams{
		Name: "Test App",
	})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "Verify/App")
}

func TestVerifyAppService_List(t *testing.T) {
	expectResponse("verifyAppListResponse.json", 200)
	assert := require.New(t)
	resp, err := client.VerifyApps.List(VerifyAppListParams{})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotNil(resp.VerifyApps)

	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.VerifyApps.List(VerifyAppListParams{})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "GET", "Verify/App")
}

func TestVerifyAppService_Get(t *testing.T) {
	expectResponse("verifyAppGetResponse.json", 200)
	appUUID := "test-app-uuid-1234"
	assert := require.New(t)
	resp, err := client.VerifyApps.Get(appUUID)
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotNil(resp.VerifyApp)

	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.VerifyApps.Get(appUUID)
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "GET", "Verify/App/%s", appUUID)
}

func TestVerifyAppService_Update(t *testing.T) {
	expectResponse("verifyAppUpdateResponse.json", 200)
	appUUID := "test-app-uuid-1234"
	assert := require.New(t)
	resp, err := client.VerifyApps.Update(appUUID, VerifyAppUpdateParams{
		Name: "Updated App",
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	assert.NotEmpty(resp.AppUUID)
	assert.Equal(resp.Error, "")

	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.VerifyApps.Update(appUUID, VerifyAppUpdateParams{
		Name: "Updated App",
	})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "Verify/App/%s", appUUID)
}

func TestVerifyAppService_Delete(t *testing.T) {
	expectResponse("verifyAppDeleteResponse.json", 200)
	appUUID := "test-app-uuid-1234"
	assert := require.New(t)
	resp, err := client.VerifyApps.Delete(appUUID)
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	assert.Equal(resp.Error, "")

	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.VerifyApps.Delete(appUUID)
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "DELETE", "Verify/App/%s", appUUID)
}