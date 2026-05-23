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

	err = client.WhatsappTemplates.Delete(
		"your_waba_id",
		"your_template_id",
		plivo.WhatsappTemplateDeleteParams{
			Name: "sample_template",
		},
	)
	if err != nil {
		fmt.Printf("Error deleting WhatsApp template: %s\n", err)
		return
	}
	fmt.Println("Template deleted successfully.")
}