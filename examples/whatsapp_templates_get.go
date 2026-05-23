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

	resp, err := client.WhatsappTemplates.Get("your_waba_id", "your_template_id")
	if err != nil {
		fmt.Printf("Error retrieving WhatsApp template: %s\n", err)
		return
	}
	fmt.Printf("Template: %+v\n", resp)
}