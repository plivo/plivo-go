package plivo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyAppTemplateService_List(t *testing.T) {
	expectResponse("verifyAppTemplateListResponse.json", 200)
	assert := assert.New(t)
	resp, err := client.VerifyAppTemplates.List()
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	assert.NotNil(resp.Templates)

	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.VerifyAppTemplates.List()
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "GET", "Verify/App/templates/")
}