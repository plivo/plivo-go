package main

import (
	"fmt"
	"github.com/plivo/plivo-go"
)

func main() {
	client, err := plivo.NewClient("MADCHANDRESH02TANK06", "OTljNmVmOGVkNGZhNjJlOWIyMWM0ZDI0ZjQwZDdk", &plivo.ClientOptions{})
	if err != nil {
		fmt.Print("Error", err.Error())
		return
	}
	response, err := client.Calls.Create(
		plivo.CallCreateParams{
			From:         "+14151234567",
			To:           "sip:ajaydjv902035012689385@phone-qa.voice.plivodev.com",
			AnswerURL:    "https://s3.amazonaws.com/static.plivo.com/answer.xml",
			AnswerMethod: "GET",
		},
	)
	if err != nil {
		fmt.Print("Error", err.Error())
		return
	}
	fmt.Printf("Response: %#v\n", response)
}
