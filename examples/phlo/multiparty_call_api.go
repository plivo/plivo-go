package main

import (
	"fmt"
	"plivo-go"
)

// Initialize the following variables with corresponding values to trigger resources

const actionTo = "to_endpoint"
const actionSource = "from_endpoint"

//Use the following multiparty call actions as needed; do not edit the following
const actionRole = "agent"
const actionCall = "call"
const actionWarmTransfer = "warm_tranfer"
const actionColdTransfer = "cold_transfer"

/*Initialize phloClient using plivo.NewPhloClient("authId, "authToken") as shown below
  For example to trigger WarmTransfer invoke testMultiPartyCallWarmTransfer() in main()

  Sample expected output for testMultiPartyCallWarmTransfer resource is
	{
	  "ApiID":"7ebd8daa-d216-495a-ab94-7c71a776be28",
	  "Error":""

	}
*/

func main() {
	testMultiPartyCallWarmTransfer()
}

func testMultiPartyCallCall() {
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
	callPayload := plivo.MultiPartyCallActionPayload{
		actionCall,
		actionTo,
		actionRole,
		actionSource,
	}

	response, _ := MPCGet.Call(callPayload)
	fmt.Printf("Response: %#v\n", response)
}

func testMultiPartyCallWarmTransfer() {
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
	callPayload := plivo.MultiPartyCallActionPayload{
		actionWarmTransfer,
		actionTo,
		actionRole,
		actionSource,
	}
	response, err := MPCGet.WarmTransfer(callPayload)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)
}

func testMultiPartyCallColdTransfer() {
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
	callPayload := plivo.MultiPartyCallActionPayload{
		actionColdTransfer,
		actionTo,
		actionRole,
		actionSource,
	}

	response, err := MPCGet.ColdTransfer(callPayload)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)
}