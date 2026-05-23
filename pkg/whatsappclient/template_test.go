package whatsappclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhatsAppTemplateCreateParams_ApplicationID(t *testing.T) {
	params := WhatsAppTemplateCreateParams{
		ElementName:  "test_template",
		LanguageCode: "en_US",
		Category:     "UTILITY",
		ApplicationID: "test-app-id",
	}
	assert.Equal(t, "test-app-id", params.ApplicationID)
	assert.Equal(t, "test_template", params.ElementName)
	assert.Equal(t, "en_US", params.LanguageCode)
	assert.Equal(t, "UTILITY", params.Category)
}

func TestWhatsAppTemplateCreateParams_ApplicationID_Optional(t *testing.T) {
	params := WhatsAppTemplateCreateParams{
		ElementName:  "test_template",
		LanguageCode: "en_US",
		Category:     "UTILITY",
	}
	assert.Equal(t, "", params.ApplicationID)
}

func TestWhatsAppTemplateUpdateParams_ApplicationID(t *testing.T) {
	params := WhatsAppTemplateUpdateParams{
		Category:      "MARKETING",
		ApplicationID: "updated-app-id",
	}
	assert.Equal(t, "updated-app-id", params.ApplicationID)
	assert.Equal(t, "MARKETING", params.Category)
}

func TestWhatsAppTemplateUpdateParams_ApplicationID_Optional(t *testing.T) {
	params := WhatsAppTemplateUpdateParams{
		Category: "MARKETING",
	}
	assert.Equal(t, "", params.ApplicationID)
}