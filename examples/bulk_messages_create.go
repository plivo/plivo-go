package main

import (
	"fmt"
	"github.com/plivo/plivo-go"
)

func main() {
	client, err := plivo.NewClient("YOUR_AUTH_ID", "YOUR_AUTH_TOKEN", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}

	response, err := client.BulkMessages.Create(plivo.BulkMessageCreateParams{
		Src:  "<sender_id_or_number>",
		Dst:  []string{"<destination_number_1>", "<destination_number_2>"},
		Text: "Hello from Plivo bulk messaging!",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("API ID: %s\n", response.ApiID)
	fmt.Printf("Message UUIDs: %v\n", response.MessageUUID)
	fmt.Printf("Message: %s\n", response.Message)
	if len(response.InvalidNumber) > 0 {
		fmt.Printf("Invalid numbers: %v\n", response.InvalidNumber)
	}
}