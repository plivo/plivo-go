package main

import (
	"fmt"
	"plivo-go"
)

// Initialize phloClient using plivo.NewPhloClient("authId, "authToken") as shown below
// To trigger phlo resource you can call the corresponding resource in main
// For example to trigger phloGet call testPhloGetter() in main()
// Sample expected output for PhloGet resource is
/*
	{
		BaseResource:{client:xxx}
		PhloId:xxxx-xxx-xxxx-xxx-xxx
		Name:xxx+xxx+flow
		CreatedOn:2018-11-20 04:46:20.307414+00:00
	}
*/

func main() {
	testNodeGetter()
}

func testPhloGetter() {
	phloClient,err := plivo.NewPhloClient("auth_id", "auth_token", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := phloClient.Phlos.Get("phlo_id")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)
}

func testNodeGetter() {
	phloClient,err := plivo.NewPhloClient("auth_id", "auth_token", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get("phlo_id")
	if err != nil {
		panic(err)
	}
	response, err := phloGet.Node("node_id")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)
}

func testPhloRunWithParams() {
	phloClient,err := plivo.NewPhloClient("auth_id", "auth_token", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get("phlo_id")
	if err != nil {
		panic(err)
	}
	type params map[string]interface{}
	response, err := phloGet.Run(params{"from": "origin", "to": "destination"})
	if (err != nil) {
		println(err)
	}
	fmt.Printf("Response: %#v\n", response)
}

func testPhloRunWithoutParams() {
	phloClient,err := plivo.NewPhloClient("auth_id", "auth_token", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get("phlo_id")
	if err != nil {
		panic(err)
	}
	response, err := phloGet.Run(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)

}

func testMulitPartyCallGetter() {
	phloClient,err := plivo.NewPhloClient("auth_id", "auth_token", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get("phlo_id")
	if err != nil {
		panic(err)
	}
	response, err := phloGet.MultiPartyCall("node_id")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)
}




