package plivo

import (
	"testing"

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

func TestCreateWhatsappTemplatePass(t *testing.T) {
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
				Type: "body",
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

func TestCreateWhatsappInteractiveList(t *testing.T) {
	interactiveData := `{
		"type": "list",
		"header": {
			"type": "text",
			"text": "Welcome to Plivo"
		},
		"body": {
			"text": "You can review the list of rewards we offer"
		},
		"footer": {
			"text": "Yours Truly"
		},
		"action": {
			"buttons": [{
				"title": "Click here",
				"id": "bt1j1k2jkjk"
			}],
			"sections": [{
				"title": "SECTION_1_TITLE",
				"rows": [{
					"id": "SECTION_1_ROW_1_ID",
					"title": "SECTION_1_ROW_1_TITLE",
					"description": "SECTION_1_ROW_1_DESCRIPTION"
				}, {
					"id": "SECTION_1_ROW_2_ID",
					"title": "SECTION_1_ROW_2_TITLE",
					"description": "SECTION_1_ROW_2_DESCRIPTION"
				}]
			}, {
				"title": "SECTION_2_TITLE",
				"rows": [{
					"id": "SECTION_2_ROW_1_ID",
					"title": "SECTION_2_ROW_1_TITLE",
					"description": "SECTION_2_ROW_1_DESCRIPTION"
				}, {
					"id": "SECTION_2_ROW_2_ID",
					"title": "SECTION_2_ROW_2_TITLE",
					"description": "SECTION_2_ROW_2_DESCRIPTION"
				}]
			}]
		}
	}`
	txt := "Welcome to Plivo"
	expectedInteractive := Interactive{
		Type: "list",
		Header: &Header{
			Type: "text",
			Text: &txt,
		},
		Body: &Body{
			Text: "You can review the list of rewards we offer",
		},
		Footer: &Footer{
			Text: "Yours Truly",
		},
		Action: &Action{
			Button: []*Buttons{
				{
					ID:    "bt1j1k2jkjk",
					Title: "Click here",
				},
			},
			Section: []*Section{
				{
					Title: "SECTION_1_TITLE",
					Row: []*Row{
						{
							ID:          "SECTION_1_ROW_1_ID",
							Title:       "SECTION_1_ROW_1_TITLE",
							Description: "SECTION_1_ROW_1_DESCRIPTION",
						},
						{
							ID:          "SECTION_1_ROW_2_ID",
							Title:       "SECTION_1_ROW_2_TITLE",
							Description: "SECTION_1_ROW_2_DESCRIPTION",
						},
					},
				},
				{
					Title: "SECTION_2_TITLE",
					Row: []*Row{
						{
							ID:          "SECTION_2_ROW_1_ID",
							Title:       "SECTION_2_ROW_1_TITLE",
							Description: "SECTION_2_ROW_1_DESCRIPTION",
						},
						{
							ID:          "SECTION_2_ROW_2_ID",
							Title:       "SECTION_2_ROW_2_TITLE",
							Description: "SECTION_2_ROW_2_DESCRIPTION",
						},
					},
				},
			},
		},
	}

	interactive, _ := CreateWhatsappInteractive(interactiveData)
	assert.Equal(t, expectedInteractive, interactive)
}

func TestCreateWhatsappInteractiveReply(t *testing.T) {
	interactiveData := `{
        "type": "reply",
        "header": {
            "type": "media",
            "media": "https://media.geeksforgeeks.org/wp-content/uploads/20190712220639/ybearoutput-300x225.png"
        },
        "body": {
            "text": "Make your selection"
        },
        "action": {
            "buttons": [
                {
                    "title": "Click here",
                    "id": "bt1"
                },
                {
                    "title": "Know More",
                    "id": "bt2"
                },
                {
                    "title": "Request Callback",
                    "id": "bt3"
                }
            ]
        }
    }`
	mediaLink := "https://media.geeksforgeeks.org/wp-content/uploads/20190712220639/ybearoutput-300x225.png"
	expectedInteractive := Interactive{
		Type: "reply",
		Header: &Header{
			Type:  "media",
			Media: &mediaLink,
		},
		Body: &Body{
			Text: "Make your selection",
		},
		Action: &Action{
			Button: []*Buttons{
				{
					ID:    "bt1",
					Title: "Click here",
				},
				{
					ID:    "bt2",
					Title: "Know More",
				},
				{
					ID:    "bt3",
					Title: "Request Callback",
				},
			},
		},
	}

	interactive, _ := CreateWhatsappInteractive(interactiveData)
	assert.Equal(t, expectedInteractive, interactive)
}

func TestCreateWhatsappInteractiveCTA(t *testing.T) {
	interactiveData := `{
        "type": "cta_url",
        "header": {
            "type": "media",
            "media": "https://media.geeksforgeeks.org/wp-content/uploads/20190712220639/ybearoutput-300x225.png"
        },
        "body": {
            "text": "Know More"
        },
        "action": {
            "buttons": [
                {
                    "title": "Click here",
                    "id": "bt1",
                    "cta_url": "https://plivo.com"
                }
            ]
        }
    }`

	mediaLink := "https://media.geeksforgeeks.org/wp-content/uploads/20190712220639/ybearoutput-300x225.png"

	expectedInteractive := Interactive{
		Type: "cta_url",
		Header: &Header{
			Type:  "media",
			Media: &mediaLink,
		},
		Body: &Body{
			Text: "Know More",
		},
		Action: &Action{
			Button: []*Buttons{
				{
					ID:     "bt1",
					Title:  "Click here",
					CTAURL: "https://plivo.com",
				},
			},
		},
	}

	interactive, err := CreateWhatsappInteractive(interactiveData)
	assert.NoError(t, err)
	assert.Equal(t, expectedInteractive, interactive)
}

func TestCreateWhatsappLocationPass(t *testing.T) {
	location_data := `{
        "latitude": "37.483307",
        "longitude": "122.148981",
        "name": "Pablo Morales",
        "address": "1 Hacker Way, Menlo Park, CA 94025"
    }`
	location := Location{
		Latitude:  "37.483307",
		Longitude: "122.148981",
		Name:      "Pablo Morales",
		Address:   "1 Hacker Way, Menlo Park, CA 94025",
	}
	locaitonCreated, _ := CreateWhatsappLocation(location_data)
	assert.Equal(t, location, locaitonCreated)
}
