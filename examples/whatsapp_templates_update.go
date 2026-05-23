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

	resp, err := client.WhatsappTemplates.Update(
		"your_waba_id",
		"your_template_id",
		plivo.WhatsappTemplateUpdateParams{
			Category: "UTILITY",
			Components: []plivo.WhatsappTemplateComponent{
				{
					Type: "BODY",
					Parameters: []plivo.Parameter{
						{
							Type: "text",
							Text: "Updated body text",
						},
					},
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("Error updating WhatsApp template: %s\n", err)
		return
	}
	fmt.Printf("Template updated: %+v\n", resp)
}