package main

import (
	"fmt"
	"plivo-go"
)

var phloClient, _ = plivo.NewPhloClient("MAZMI2NZE5N2EWZDI4MZ", "NzUyYTVhMTY5MDczZmRjNDk1NmI5YTJmNTgwMDI4", &plivo.ClientOptions{})

func main() {
	//testMultiPartyCallCall()
	//testMultiPartyCallMemberHold()
	//testMultiPartyCallMemberUnHold()
	//testMultiPartyCallColdTransfer()
}

func testPhloGetter() {
	phloGet, _ := phloClient.Phlos.Get("2e802f94-cc03-4e35-b4c8-5a0328544135")
	fmt.Println(phloGet)
}

func testNodeGetter() {
	phloGet, _ := phloClient.Phlos.Get("2e802f94-cc03-4e35-b4c8-5a0328544135")
	nodeGet, _ := phloGet.Node("b18848dd-e308-4758-bc5c-fb9c274d5fb7")
	fmt.Println(nodeGet)
}

func testPhloRunWithParams() {
	phloGet, _ := phloClient.Phlos.Get("2e802f94-cc03-4e35-b4c8-5a0328544135")
	type params map[string]interface{}
	phloRunWithParams, err := phloGet.Run(params{"from": "919441456465", "to": "919503364736"})
	if (err != nil) {
		println(err)
	}
	fmt.Println(phloRunWithParams)
}

func testPhloRunWithoutParams() {
	phloGet, _ := phloClient.Phlos.Get("d32c92d3-deb0-4954-af00-25f919c2767f")

	phloRunWithoutParams, err := phloGet.Run(nil)
	if (err != nil) {
		println(err)
	}
	fmt.Println(phloRunWithoutParams)

}

func testMulitPartyCallGetter() {
	phloGet, _ := phloClient.Phlos.Get("2e802f94-cc03-4e35-b4c8-5a0328544135")
	MPCGet, _ := phloGet.MultiPartyCall("b18848dd-e308-4758-bc5c-fb9c274d5fb7")
	fmt.Println(MPCGet)
}




