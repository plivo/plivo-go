package main

import (
	"fmt"

	plivo "github.com/plivo/plivo-go"
)

func main() {
	client, err := plivo.NewClient("YOUR_AUTH_ID", "YOUR_AUTH_TOKEN", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}

	resp, err := client.VerifyAppTemplates.List()
	if err != nil {
		panic(err)
	}

	fmt.Printf("API ID: %s\n", resp.ApiID)
	for _, tmpl := range resp.Templates {
		fmt.Printf("Template UUID: %s, Friendly Name: %s, Locale: %s\n",
			tmpl.TemplateUUID, tmpl.FriendlyName, tmpl.Locale)
	}
}