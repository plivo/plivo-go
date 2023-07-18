package plivo

import "testing"
import (
	"github.com/stretchr/testify/assert"
)

func TestNumbers(t *testing.T) {
	num1 := "+911231231230"
	num2 := "+913213213210"
	joined := "+911231231230<+913213213210"
	assert.Equal(t, joined, Numbers(num1, num2))
}

func TestHeaders(t *testing.T) {
	headers := map[string]string{
		"X-PH-Test1": "value1",
		"X-PH-Test2": "value2",
	}
	encoded := Headers(headers)
	// Go iterates over maps in random order
	assert.Contains(t, []string{"X-PH-Test1=value1,X-PH-Test2=value2", "X-PH-Test2=value2,X-PH-Test1=value1"}, encoded)
}

func TestComputeSignature(t *testing.T) {
	assert.Equal(t,
		"EJEt0ELanhr8hjMPIJnLNLex0dE=",
		ComputeSignature("MAXXXXXXXXXXXXXXXXXX", "http://foo.com/answer/", map[string]string{
			"CallUUID": "97ceeb52-58b6-11e1-86da-77300b68f8bb",
			"Duration": "300",
		}))
}

func TestValidateSignature(t *testing.T) {
	assert.Equal(t, true,
		ValidateSignature("MAXXXXXXXXXXXXXXXXXX", "http://foo.com/answer/", map[string]string{
			"CallUUID": "97ceeb52-58b6-11e1-86da-77300b68f8bb",
			"Duration": "300",
		},
			"EJEt0ELanhr8hjMPIJnLNLex0dE="))
}

func TestComputeSignatureEncoding(t *testing.T) {
	assert.Equal(t, "n3Xfo4u+vRFyl3gsH8B0qDUIK5g=",
		ComputeSignature("MAXXXXXXXXXXXXXXXXXX", "http://foo.com/answer/", map[string]string{
			"a": "1 2",
		}))
}

func TestComputeSignatureFail(t *testing.T) {
	assert.Equal(t, false,
		ValidateSignature("MAXXXXXXXXXXXXXXXXXX", "http://foo.com/answer/", map[string]string{
			"CallUUID": "97ceeb52-58b6-11e1-86da-77300b6b8f8bb",
			"Duration": "300",
		},
			"EJEt0ELanhr8hjMPIJnLNLex0dE="))
}

func TestValidateSignatureV2Pass(t *testing.T) {
	assert.Equal(t, true,
		ValidateSignatureV2(
			"https://answer.url",
			"12345",
			"ehV3IKhLysWBxC1sy8INm0qGoQYdYsHwuoKjsX7FsXc=",
			"my_auth_token",
		),
	)
}

func TestValidateSignatureV3Pass1(t *testing.T) {
	assert.Equal(t, true,
		ValidateSignatureV3(
			"https://plivobin.non-prod.plivops.com/api/v1/validate_signature03.xml/?a=b&c=d",
			"31627761595286130198",
			"POST",
			"k7Pusd4OxCIjR5IfA9iedDNu/h/gbdYqdzG/MiYtd1c=",
			"Y2Q2ZDgxZmY5YWRiOTI5YmQ1Njg0MTAxZWIyOTc4",
			map[string]string{
				"Direction":       "outbound",
				"From":            "19792014278",
				"ALegUUID":        "3e82ae9d-2c78-4d85-b1a4-6eae7dbafb36",
				"CallStatus":      "in-progress",
				"BillRate":        "0.002",
				"ParentAuthID":    "MANWVLYTK4ZWU1YTY4QA",
				"To":              "sip:PlivoSignature382029104058171078704104@phone-qa.voice.plivodev.com",
				"ALegRequestUUID": "3e82ae9d-2c78-4d85-b1a4-6eae7dbafb36",
				"CallUUID":        "3e82ae9d-2c78-4d85-b1a4-6eae7dbafb36",
				"RequestUUID":     "3e82ae9d-2c78-4d85-b1a4-6eae7dbafb36",
				"SIP-H-To":        "<sip:PlivoSignature382029104058171078704104@52.9.11.55;transport=udp>;tag=1",
				"SessionStart":    "2020-04-08 11:34:33.238707",
				"Event":           "StartApp",
			},
		),
	)
}

func TestValidateSignatureV3Pass2(t *testing.T) {
	assert.Equal(t, true,
		ValidateSignatureV3(
			"https://plivobin.non-prod.plivops.com/api/v1/validate_signature03.xml/?a=b&c=d",
			"31627761595286130198",
			"GET",
			"UBq8jAtd32wR8EK9VgxbBn4n5rpI/l1H9iN4WfSEHFQ=",
			"Y2Q2ZDgxZmY5YWRiOTI5YmQ1Njg0MTAxZWIyOTc4",
			map[string]string{
				"Direction":       "outbound",
				"From":            "19792014278",
				"ALegUUID":        "3e82ae9d-2c78-4d85-b1a4-6eae7dbafb36",
				"CallStatus":      "in-progress",
				"BillRate":        "0.002",
				"ParentAuthID":    "MANWVLYTK4ZWU1YTY4QA",
				"To":              "sip:PlivoSignature382029104058171078704104@phone-qa.voice.plivodev.com",
				"ALegRequestUUID": "3e82ae9d-2c78-4d85-b1a4-6eae7dbafb36",
				"CallUUID":        "3e82ae9d-2c78-4d85-b1a4-6eae7dbafb36",
				"RequestUUID":     "3e82ae9d-2c78-4d85-b1a4-6eae7dbafb36",
				"SIP-H-To":        "<sip:PlivoSignature382029104058171078704104@52.9.11.55;transport=udp>;tag=1",
				"SessionStart":    "2020-04-08 11:34:33.238707",
				"Event":           "StartApp",
			},
		),
	)
}

func TestValidateSignatureV3Pass3(t *testing.T) {
	assert.Equal(t, true,
		ValidateSignatureV3(
			"https://plivobin.non-prod.plivops.com/api/v1/validate_signature03.xml",
			"31627761595286130198",
			"POST",
			"iAjE5QqI37mbkYe4w3jTMudqEzbDufdqi7sYwTu64e0=",
			"Y2Q2ZDgxZmY5YWRiOTI5YmQ1Njg0MTAxZWIyOTc4",
			map[string]string{
				"Direction":       "outbound",
				"From":            "19792014278",
				"ALegUUID":        "3e82ae9d-2c78-4d85-b1a4-6eae7dbafb36",
				"CallStatus":      "in-progress",
				"BillRate":        "0.002",
				"ParentAuthID":    "MANWVLYTK4ZWU1YTY4QA",
				"To":              "sip:PlivoSignature382029104058171078704104@phone-qa.voice.plivodev.com",
				"ALegRequestUUID": "3e82ae9d-2c78-4d85-b1a4-6eae7dbafb36",
				"CallUUID":        "3e82ae9d-2c78-4d85-b1a4-6eae7dbafb36",
				"RequestUUID":     "3e82ae9d-2c78-4d85-b1a4-6eae7dbafb36",
				"SIP-H-To":        "<sip:PlivoSignature382029104058171078704104@52.9.11.55;transport=udp>;tag=1",
				"SessionStart":    "2020-04-08 11:34:33.238707",
				"Event":           "StartApp",
			},
		),
	)
}

func TestValidateSignatureV3Pass4(t *testing.T) {
	assert.Equal(t, true,
		ValidateSignatureV3(
			"https://plivobin.non-prod.plivops.com/api/v1/validate_signature03.xml",
			"31627761595286130198",
			"GET",
			"i/MQsaQSAd6fiKhOh2qeeeLHZ9faldADSb3/7+Akfbc=",
			"Y2Q2ZDgxZmY5YWRiOTI5YmQ1Njg0MTAxZWIyOTc4",
			map[string]string{
				"Direction":       "outbound",
				"From":            "19792014278",
				"ALegUUID":        "3e82ae9d-2c78-4d85-b1a4-6eae7dbafb36",
				"CallStatus":      "in-progress",
				"BillRate":        "0.002",
				"ParentAuthID":    "MANWVLYTK4ZWU1YTY4QA",
				"To":              "sip:PlivoSignature382029104058171078704104@phone-qa.voice.plivodev.com",
				"ALegRequestUUID": "3e82ae9d-2c78-4d85-b1a4-6eae7dbafb36",
				"CallUUID":        "3e82ae9d-2c78-4d85-b1a4-6eae7dbafb36",
				"RequestUUID":     "3e82ae9d-2c78-4d85-b1a4-6eae7dbafb36",
				"SIP-H-To":        "<sip:PlivoSignature382029104058171078704104@52.9.11.55;transport=udp>;tag=1",
				"SessionStart":    "2020-04-08 11:34:33.238707",
				"Event":           "StartApp",
			},
		),
	)
}

func TestCreateWhatsappTemplatePass(t *testing.T){
	template_data := `{
        "name": "plivo_verification",
        "language": "en_US",
        "components": [
        {
          "type": "body",
          "parameters": [
            {
              "type": "text",
              "text": "J$FpnYnP"
            }
          ]
        },
        {
          "type": "button",
          "sub_type": "url",
          "index": "0",
          "parameters": [
            {
              "type": "text",
              "text": "J$FpnYnP"
            }
          ]
        }
      ]
    }`
	template := Template{
		Name:     "plivo_verification",
		Language: "en_US",
		Components: []Component{
			{
				Type:    "body",
				Parameters: []Parameter{
					{
						Type: "text",
						Text: "J$FpnYnP",
					},
				},
			},
			{
				Type:    "button",
				SubType: "url",
				Index:   "0",
				Parameters: []Parameter{
					{
						Type: "text",
						Text: "J$FpnYnP",
					},
				},
			},
		},
	}
	templateCreated, _ := CreateWhatsappTemplate(template_data)
	assert.Equal(t, template, templateCreated)
}




