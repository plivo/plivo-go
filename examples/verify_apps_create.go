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

	resp, err := client.VerifyApps.Create(plivo.VerifyAppCreateParams{
		Name:      "My Verify App",
		OtpLength: 6,
		OtpExpiry: 5,
		BrandName: "My Brand",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("App UUID: %s\n", resp.AppUUID)
	fmt.Printf("Message: %s\n", resp.Message)
}