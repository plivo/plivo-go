package main

import (
	"fmt"
	"plivo-go"
)
// Initialize phloClient using plivo.NewPhloClient("authId, "authToken") as shown below
// To trigger phlo resource you can call the corresponding resource in main
// For example to trigger testMultiPartyCallAbortTransfer call testMultiPartyCallAbortTransfer() in main()
// Sample expected output for testMultiPartyCallAbortTransfer resource is
/*

	{
	  "ApiID":"7ebd8daa-d216-495a-ab94-7c71a776be28",
	  "Error":""

	}
*/
func testMultiPartyCallAbortTransfer() {
	phloClient,err := plivo.NewPhloClient("auth_id", "auth_token", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get("phlo_id")
	if err != nil {
		panic(err)
	}
	MPCGet, err := phloGet.MultiPartyCall("node_id")
	if err != nil {
		panic(err)
	}
	response, err := MPCGet.Member("member_id").AbortTransfer()
	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Printf("Response: %#v\n", response)

}

func testMultiPartyCallMemberHold() {
	phloClient,err := plivo.NewPhloClient("auth_id", "auth_token", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get("phlo_id")
	if err != nil {
		panic(err)
	}
	MPCGet, err := phloGet.MultiPartyCall("node_id")
	if err != nil {
		panic(err)
	}
	response, err := MPCGet.Member("member_id").Hold()
	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Printf("Response: %#v\n", response)

}

func testMultiPartyCallMemberUnHold() {
	phloClient,err := plivo.NewPhloClient("auth_id", "auth_token", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get("phlo_id")
	if err != nil {
		panic(err)
	}
	MPCGet, err := phloGet.MultiPartyCall("node_id")
	if err != nil {
		panic(err)
	}
	response, err := MPCGet.Member("member_id").UnHold()
	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Printf("Response: %#v\n", response)

}

func testMultiPartyCallMemberVoiceMailDrop() {
	phloClient,err := plivo.NewPhloClient("auth_id", "auth_token", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get("phlo_id")
	if err != nil {
		panic(err)
	}
	MPCGet, err := phloGet.MultiPartyCall("node_id")
	if err != nil {
		panic(err)
	}
	response, err := MPCGet.Member("member_id").VoiceMailDrop()
	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Printf("Response: %#v\n", response)

}

func testMultiPartyCallMemberHangUp() {
	phloClient,err := plivo.NewPhloClient("auth_id", "auth_token", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get("phlo_id")
	if err != nil {
		panic(err)
	}
	MPCGet, err := phloGet.MultiPartyCall("node_id")
	if err != nil {
		panic(err)
	}
	response, err := MPCGet.Member("member_id").HangUp()
	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Printf("Response: %#v\n", response)

}

func testMultiPartyCallMemberResume() {
	phloClient,err := plivo.NewPhloClient("auth_id", "auth_token", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get("phlo_id")
	if err != nil {
		panic(err)
	}
	MPCGet, err := phloGet.MultiPartyCall("node_id")
	if err != nil {
		panic(err)
	}
	response, err := MPCGet.Member("member_id").ResumeCall()
	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Printf("Response: %#v\n", response)

}