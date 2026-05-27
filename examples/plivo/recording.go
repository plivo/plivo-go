package main

import (
	"encoding/json"
	"fmt"
	"github.com/plivo/plivo-go"
)

// Initialize the AuthId and AuthToken parameters in credentials.go
// To trigger Recording resource methods invoke corresponding helper method in main.go

const recordingId="1ca34b00-XXXXX-XXXX-06bcf6c57c65'"

// Example to get details of a recording
// Pass recordingID to fetch the details
func TestRecordingGet(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Recordings.Get(
		recordingId,
	)
	if err != nil {
		panic(err)
	}

	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

// Example to delete  a recording
// Pass recordingID to delete

func TestRecordingDelete(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	err = client.Recordings.Delete(
		recordingId,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted Recording successfully ")
}


// Example to list all recordings
func TestRecordingList(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Recordings.List(
		plivo.RecordingListParams{
			Offset: 0,
			Limit: 5,
		},
	)
	if err != nil {
		panic(err)
	}

	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}