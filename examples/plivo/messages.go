package main

import (
	"encoding/json"
	"fmt"
	"github.com/plivo/plivo-go"
)

// Initialize the AuthId and AuthToken parameters in credentials.go
// To trigger Message resource methods invoke corresponding helper method in main.go


//Example to create messsage
//The following source, destination and text are mandatory params to send a message

func TestMessageCreate(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Messages.Create(
		plivo.MessageCreateParams{
			Src: "from_number",
			Dst: "to_number",
			Text: "Test Message",
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

// Example to get a  messsage
// Pass appropriate messageUUID to get details of a message
func TestMessageGet(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Messages.Get(
		"messageUUID",
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to list all messsages
func TestMessageList(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Messages.List(
		plivo.MessageListParams{
			Limit: 5,
			Offset: 0,
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}