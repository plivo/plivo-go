package main

import (
	"fmt"
	"plivo-go"
)

// Initialize the following member_id with corresponding value to trigger action on member resource

const memberId  = "member_id"

/*Initialize phloClient using plivo.NewPhloClient("authId, "authToken") as shown below
 For example to trigger CallAbortTransfer invoke testMultiPartyCallAbortTransfer() in main()
 Sample expected output for testMultiPartyCallAbortTransfer resource is

	{
	  "ApiID":"7ebd8daa-d216-495a-ab94-7c71a776be28",
	  "Error":""

	}
*/

func main() {
	testMultiPartyCallAbortTransfer()
}

func testMultiPartyCallAbortTransfer() {
	phloClient,err := plivo.NewPhloClient(authId, authToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get(phloId)
	if err != nil {
		panic(err)
	}
	MPCGet, err := phloGet.MultiPartyCall(nodeId)
	if err != nil {
		panic(err)
	}
	response, err := MPCGet.Member(memberId).AbortTransfer()
	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Printf("Response: %#v\n", response)

}

func testMultiPartyCallMemberHold() {
	phloClient,err := plivo.NewPhloClient(authId, authToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get(phloId)
	if err != nil {
		panic(err)
	}
	MPCGet, err := phloGet.MultiPartyCall(nodeId)
	if err != nil {
		panic(err)
	}
	response, err := MPCGet.Member(memberId).Hold()
	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Printf("Response: %#v\n", response)

}

func testMultiPartyCallMemberUnHold() {
	phloClient,err := plivo.NewPhloClient(authId, authToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get(phloId)
	if err != nil {
		panic(err)
	}
	MPCGet, err := phloGet.MultiPartyCall(nodeId)
	if err != nil {
		panic(err)
	}
	response, err := MPCGet.Member(memberId).UnHold()
	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Printf("Response: %#v\n", response)

}

func testMultiPartyCallMemberVoiceMailDrop() {
	phloClient,err := plivo.NewPhloClient(authId, authToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get(phloId)
	if err != nil {
		panic(err)
	}
	MPCGet, err := phloGet.MultiPartyCall(nodeId)
	if err != nil {
		panic(err)
	}
	response, err := MPCGet.Member(memberId).VoiceMailDrop()
	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Printf("Response: %#v\n", response)

}

func testMultiPartyCallMemberHangUp() {
	phloClient,err := plivo.NewPhloClient(authId, authToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get(phloId)
	if err != nil {
		panic(err)
	}
	MPCGet, err := phloGet.MultiPartyCall(nodeId)
	if err != nil {
		panic(err)
	}
	response, err := MPCGet.Member(memberId).HangUp()
	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Printf("Response: %#v\n", response)

}

func testMultiPartyCallMemberResume() {
	phloClient,err := plivo.NewPhloClient(authId, authToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get(phloId)
	if err != nil {
		panic(err)
	}
	MPCGet, err := phloGet.MultiPartyCall(nodeId)
	if err != nil {
		panic(err)
	}
	response, err := MPCGet.Member(memberId).ResumeCall()
	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Printf("Response: %#v\n", response)

}