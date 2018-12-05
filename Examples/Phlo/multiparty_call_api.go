package main

import (
	"fmt"
	"plivo-go"
)
// Initialize phloClient using plivo.NewPhloClient("authId, "authToken") as shown below
// To trigger phlo resource you can call the corresponding resource in main
// For example to trigger testMultiPartyCallWarmTransfer call testMultiPartyCallWarmTransfer() in main()
// Sample expected output for testMultiPartyCallWarmTransfer resource is
/*

	{
	  "ApiID":"7ebd8daa-d216-495a-ab94-7c71a776be28",
	  "Error":""

	}
*/

func main() {
	testMultiPartyCallWarmTransfer()
}

func testMultiPartyCallCall() {
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
	callPayload := plivo.MultiPartyCallActionPayload{
		"call",
		"destination",
		"agent",
		"origin",
	}

	response, _ := MPCGet.Call(callPayload)
	fmt.Printf("Response: %#v\n", response)
}

func testMultiPartyCallWarmTransfer() {
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
	callPayload := plivo.MultiPartyCallActionPayload{
		"warm_tranfer",
		"destination",
		"agent",
		"origin",
	}
	response, err := MPCGet.WarmTransfer(callPayload)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)
}

func testMultiPartyCallColdTransfer() {
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
	callPayload := plivo.MultiPartyCallActionPayload{
		"cold_transfer",
		"destination",
		"agent",
		"origin",
	}

	response, err := MPCGet.ColdTransfer(callPayload)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)
}