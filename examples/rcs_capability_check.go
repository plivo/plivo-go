package main

import (
	"fmt"

	"github.com/plivo/plivo-go/v7"
)

func main() {
	client, err := plivo.NewClient("YOUR_AUTH_ID", "YOUR_AUTH_TOKEN", &plivo.Options{})
	if err != nil {
		panic(err)
	}

	response, err := client.RcsCapability.Check(plivo.RcsCapabilityParams{
		PhoneNumber: "+14151234567",
		AgentUUID:   "your-agent-uuid",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("API ID: %s\n", response.ApiID)
	fmt.Printf("Phone Number: %s\n", response.PhoneNumber)
	fmt.Printf("Is Capable: %v\n", response.IsCapable)
	fmt.Printf("Features: %v\n", response.Features)
	fmt.Printf("Message: %s\n", response.Message)
}