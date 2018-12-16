package main

import (
	"fmt"
	"plivo-go"
)


// Initialize the following params with corresponding values to trigger resources

const authId  = "auth_id"
const authToken = "auth_token"
const phloId = "phlo_id"
const nodeId = "node_id"

/* Initialize phloClient using plivo.NewPhloClient("authId, "authToken") as shown below
  For example to trigger phloGet invoke testPhloGetter() in main()

 Sample expected output for PhloGet resource is

	{
		BaseResource:{client:xxx}
		PhloId:xxxx-xxx-xxxx-xxx-xxx
		Name:xxx+xxx+flow
		CreatedOn:2018-11-20 04:46:20.307414+00:00
	}
*/

func main() {
	testPhloGetter()
}

func testPhloGetter() {
	phloClient,err := plivo.NewPhloClient(authId, authToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := phloClient.Phlos.Get(phloId)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)
}

func testNodeGetter() {
	phloClient,err := plivo.NewPhloClient(authId, authToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get(phloId)
	if err != nil {
		panic(err)
	}
	response, err := phloGet.Node(nodeId)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)
}

func testPhloRunWithParams() {
	phloClient,err := plivo.NewPhloClient(authId, authToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get(phloId)
	if err != nil {
		panic(err)
	}
	//pass corresponding from and to values
	type params map[string]interface{}
	response, err := phloGet.Run(params{
		"from": "111111111",
		"to": "2222222222",
	})

	if (err != nil) {
		println(err)
	}
	fmt.Printf("Response: %#v\n", response)
}

func testPhloRunWithoutParams() {
	phloClient,err := plivo.NewPhloClient(authId, authToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get(phloId)
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
	phloClient,err := plivo.NewPhloClient(authId, authToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	phloGet, err := phloClient.Phlos.Get(phloId)
	if err != nil {
		panic(err)
	}
	response, err := phloGet.MultiPartyCall(nodeId)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)
}




