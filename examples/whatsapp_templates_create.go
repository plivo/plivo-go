package main

import (
	"fmt"

	"github.com/plivo/plivo-go/v7"
)

func main() {
	client, err := plivo.NewClient("YOUR_AUTH_ID", "YOUR_AUTH_TOKEN", &plivo.ClientOptions{})
	if err != nil {
		fmt.Printf("Error initializing client: %s\n", err)
		return
	}

	resp, err := client.WhatsappTemplates.Create(
		"your_waba_id",
		plivo.WhatsappTemplateCreateParams{
			Name:                "sample_template",
			Category:            "MARKETING",
			Language:            "en_US",
			AllowCategoryChange: true,
			Components: []plivo.WhatsappTemplateComponent{
				{
					Type: "BODY",
					Parameters: []plivo.Parameter{
						{
							Type: "text",
							Text: "Hello World",
						},
					},
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("Error creating WhatsApp template: %s\n", err)
		return
	}
	fmt.Printf("Template created: %+v\n", resp)
}