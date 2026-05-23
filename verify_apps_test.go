package plivo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifyAppService_Create(t *testing.T) {
	expectResponse("verifyAppCreateResponse.json", 201)
	assert := require.New(t)
	resp, err := client.VerifyApps.Create(VerifyAppCreateParams{
		Name:      "TestApp",
		BrandName: "TestBrand",
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	assert.NotEmpty(resp.AppUUID)

	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.VerifyApps.Create(VerifyAppCreateParams{
		Name: "TestApp",
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

func TestVerifyAppService_ListTemplates(t *testing.T) {
	expectResponse("verifyAppListTemplatesResponse.json", 200)
	assert := require.New(t)
	resp, err := client.VerifyApps.ListTemplates()
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotNil(resp.Templates)

	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.VerifyApps.ListTemplates()
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "GET", "Verify/App/templates")
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
		Name: "UpdatedApp",
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	assert.NotEmpty(resp.AppUUID)

	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.VerifyApps.Update(appUUID, VerifyAppUpdateParams{
		Name: "UpdatedApp",
	})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "Verify/App/%s", appUUID)
}

func TestVerifyAppService_Delete(t *testing.T) {
	expectResponse("verifyAppDeleteResponse.json", 204)
	appUUID := "test-app-uuid-1234"
	assert := require.New(t)
	err := client.VerifyApps.Delete(appUUID)
	assert.Nil(err)

	cl := client.httpClient
	client.httpClient = nil
	err = client.VerifyApps.Delete(appUUID)
	assert.NotNil(err)
	client.httpClient = cl

	assertRequest(t, "DELETE", "Verify/App/%s", appUUID)
}