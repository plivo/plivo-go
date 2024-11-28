package main

import (
	"encoding/json"
	"fmt"
	"github.com/plivo/plivo-go"
)

// Initialize the AuthId and AuthToken parameters in credentials.go
// To trigger Call resource methods invoke corresponding helper method in main.go


//Example to create a Call
//Use CallCreateParams to create a Call

func TestCreateCall(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Calls.Create(
		plivo.CallCreateParams{
			From: "2222222222", // The phone number to be used as the caller id
			To : "1111111111", // The phone numer to which the call will be placed
			AnswerURL: "https://s3.amazonaws.com/static.plivo.com/answer.xml",// answer_url is the URL invoked by Plivo when the outbound call is answered
															// and contains instructions telling Plivo what to do with the call
		     AnswerMethod: "GET", //The method used to call the answer_url

		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to get a Call
//Pass corresponding callID to get details of a call

func TestGetCall(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Calls.Get(
		"10f0cb68-7533-XXXXX-acb5-87ceac29ee48",
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to update a Call
//Pass corresponding callID and update CallUpdateParams to update details of a call

func TestUpdateCallDetails(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Calls.Update(
		"10f0cb68-7533-45ed-acb5-87ceac29ee48",
		plivo.CallUpdateParams{
			Legs: "ALEG",
			AlegURL:"https://example.com/connect/", //URL to transfer for aleg
			 AlegMethod: "GET", // method to invoke the aleg_url
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to delete a Call
//Pass corresponding callID to delete details of a call

func TestCallDelete(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	err = client.Calls.Delete(
		"10f0cb68-7533-45ed-acb5-87ceac29ee48",
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Call deleted successfully.")
}

//Example to list all Calls
//Pass offset and limit values to limit the number of call details

func TestGetCallList(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Calls.List(
		plivo.CallListParams{
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

//Example to get Live call details
//Pass corresponding liveCallID to get details of a liveCall

func TestLiveCallGet(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.LiveCalls.Get(
		"LiveCallID", // The ID of the call
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to get Live call lsit
//Pass corresponding liveCallID to get details of a liveCall

func TestLiveCallIdList(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.LiveCalls.IDList(
		plivo.LiveCallFilters{
			"Inbound",
			"123456789",
			"123456789",
			"live",
		})
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

// Example to get retrieve details of a specific queued call.
// Pass caller_id param

func TestLiveQueuedCallGet(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.QueuedCalls.Get("QueueCallID")
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

// Example to retrieve details of all queued calls
// Pass caller_id param

func TestLiveQueuedCallIdList(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.QueuedCalls.IDList()
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to start recording a call.
// Pass caller_id param

func TestCallRecord() {
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Calls.Record(
		"3a2e4c90-dcee-4931-8a59-f123ab507e60",
		plivo.CallRecordParams{
			30,
			"mp3",
			"auto",
			"transcription_url",
			"POST",
			"call_bacl_url",
			"GET",

		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)
}

//Example to stop call record
// Pass caller_id param

func TestStopCallRecord() {
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	err = client.Calls.StopRecording("3a2e4c90-dcee-4931-8a59-f123ab507e60")
	if err != nil {
		panic(err)
	}
	fmt.Println("Stopped recording successfully.")
}

// Example to call
// Pass caller_id param

func TestCallSpeak(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Calls.Speak(
		"3a2e4c90-dcee-4931-8a59-f123",
		plivo.CallSpeakParams{
			Text: "Hello Test",
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)
}

// Example to stop call speak
// Pass caller_id param

func TestStopCallSpeaking(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	err = client.Calls.StopSpeaking(
		"3a2e4c90-dcee-4931-8a59-f123",
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Stopped call successfully.")

}

// Example to Test call play
// Pass caller_id param

func TestCallPlay() {
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Calls.Play(
		"3a2e4c90-dcee-4931-8a59-f123",
		plivo.CallPlayParams{
			URLs: "https://Tests.mp3",
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)
}

// Example to stop call playing
// Pass caller_id param

func TestStopCallPlaying(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	err = client.Calls.StopPlaying(
		"3a2e4c90-dcee-4931-8a59-f123",
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted successfully.")
}

// Example to send digits
// Pass caller_id param

func TestSendDigits() {
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Calls.SendDigits(
		"3a2e4c90-dcee-4931-8a59-f123",
		plivo.CallDTMFParams{
			Digits: "123",
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)
}

// Example to cancel call request
// Pass caller_id param

func TestCallCancelRequest() {
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	err = client.Calls.CancelRequest("3a2e4c90-dcee-4931-8a59-f123")

	if(err!=nil){
		fmt.Println(err)
	}

	fmt.Println("Deleted successfully.")
}