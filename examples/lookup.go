package main

import (
	"fmt"
	"log"

	"github.com/plivo/plivo-go"
)

const (
	authId    = ""
	authToken = ""
)

func main() {
	client, err := plivo.NewClient(authId, authToken, &plivo.ClientOptions{})
	if err != nil {
		log.Fatalf("plivo.NewClient() failed: %s", err.Error())
	}

	lookup(client)
}

func lookup(client *plivo.Client) {
	resp, err := client.Lookup.Get("<insert-number-here>",
		plivo.LookupParams{
			Type: "carrier",
		})
	if err != nil {
		if respErr, ok := err.(*plivo.LookupError); ok {
			fmt.Printf("API ID: %s\nError Code: %d\nMessage: %s\n",
				respErr.ApiID, respErr.ErrorCode, respErr.Message)
			return
		}
		log.Fatalf("client.Lookup.Get() failed: %s", err.Error())
	}

	fmt.Printf("%+v\n", resp)
	fmt.Printf("%+v\n", resp.Country)
	fmt.Printf("%+v\n", resp.Format)
	fmt.Printf("%+v\n", resp.Carrier)
}
