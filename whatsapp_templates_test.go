package plivo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWhatsappTemplateService_Create(t *testing.T) {
	expectResponse("whatsappTemplateCreateResponse.json", 200)
	assert := require.New(t)
	wabaID := "test-waba-id"
	resp, err := client.WhatsappTemplates.Create(wabaID, WhatsappTemplateCreateParams{
		Name:     "test_template",
		Category: "MARKETING",
		Language: "en_US",
	})
	assert.NotNil(resp)
	assert.Nil(err)

	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.WhatsappTemplates.Create(wabaID, WhatsappTemplateCreateParams{})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "WhatsApp/Template/%s/", wabaID)
}

func TestWhatsappTemplateService_Update(t *testing.T) {
	expectResponse("whatsappTemplateUpdateResponse.json", 200)
	assert := require.New(t)
	wabaID := "test-waba-id"
	templateID := "test-template-id"
	resp, err := client.WhatsappTemplates.Update(wabaID, templateID, WhatsappTemplateUpdateParams{
		Category: "UTILITY",
	})
	assert.NotNil(resp)
	assert.Nil(err)

	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.WhatsappTemplates.Update(wabaID, templateID, WhatsappTemplateUpdateParams{})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "WhatsApp/Template/%s/%s/", wabaID, templateID)
}

func TestWhatsappTemplateService_Get(t *testing.T) {
	expectResponse("whatsappTemplateGetResponse.json", 200)
	assert := require.New(t)
	wabaID := "test-waba-id"
	templateID := "test-template-id"
	resp, err := client.WhatsappTemplates.Get(wabaID, templateID)
	assert.NotNil(resp)
	assert.Nil(err)

	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.WhatsappTemplates.Get(wabaID, templateID)
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "GET", "WhatsApp/Template/%s/%s/", wabaID, templateID)
}

func TestWhatsappTemplateService_List(t *testing.T) {
	expectResponse("whatsappTemplateListResponse.json", 200)
	assert := require.New(t)
	wabaID := "test-waba-id"
	resp, err := client.WhatsappTemplates.List(wabaID, WhatsappTemplateListParams{})
	assert.NotNil(resp)
	assert.Nil(err)

	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.WhatsappTemplates.List(wabaID, WhatsappTemplateListParams{})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "GET", "WhatsApp/Template/%s/", wabaID)
}

func TestWhatsappTemplateService_Delete(t *testing.T) {
	expectResponse("whatsappTemplateDeleteResponse.json", 204)
	assert := require.New(t)
	wabaID := "test-waba-id"
	templateID := "test-template-id"
	err := client.WhatsappTemplates.Delete(wabaID, templateID, WhatsappTemplateDeleteParams{
		Name: "test_template",
	})
	assert.Nil(err)

	cl := client.httpClient
	client.httpClient = nil
	err = client.WhatsappTemplates.Delete(wabaID, templateID, WhatsappTemplateDeleteParams{Name: "test_template"})
	assert.NotNil(err)
	client.httpClient = cl

	assertRequest(t, "DELETE", "WhatsApp/Template/%s/%s/", wabaID, templateID)
}