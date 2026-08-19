// Example for Call create
package main

import (
	"fmt"
	"github.com/plivo/plivo-go/v7"
)

func main() {
	client, err := plivo.NewClient("MAYJI2ZJDIMTVIODIWYJ", "OTI2NWFmNGI4MmZjZjZkOTQ0YjNkYTQwMzY2ZDJl", &plivo.ClientOptions{})
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
