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

	resp, err := client.WhatsappTemplates.List(
		"your_waba_id",
		plivo.WhatsappTemplateListParams{
			Limit:  20,
			Offset: 0,
		},
	)
	if err != nil {
		fmt.Printf("Error listing WhatsApp templates: %s\n", err)
		return
	}
	fmt.Printf("Templates: %+v\n", resp)
}