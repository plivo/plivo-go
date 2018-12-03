package main

import (
	"fmt"
	"plivo-go"
)

func testMultiPartyCallCall() {
	phloGet, _ := phloClient.Phlos.Get("2e802f94-cc03-4e35-b4c8-5a0328544135")
	MPCGet, _ := phloGet.MultiPartyCall("b18848dd-e308-4758-bc5c-fb9c274d5fb7")
	callPayload := plivo.MultiPartyCallActionPayload{"call",
		"919704583677",
		"agent",
		"919503364736"}

	callResponse, _ := MPCGet.Call(callPayload)
	fmt.Println(callResponse)
}

func testMultiPartyCallWarmTransfer() {
	phloGet, _ := phloClient.Phlos.Get("2e802f94-cc03-4e35-b4c8-5a0328544135")
	MPCGet, _ := phloGet.MultiPartyCall("b18848dd-e308-4758-bc5c-fb9c274d5fb7")
	callPayload := plivo.MultiPartyCallActionPayload{"warm_tranfer",
		"919441456465",
		"agent",
		"919503364736"}

	callResponse, _ := MPCGet.WarmTransfer(callPayload)
	fmt.Println(callResponse)
}

func testMultiPartyCallColdTransfer() {
	phloGet, _ := phloClient.Phlos.Get("2e802f94-cc03-4e35-b4c8-5a0328544135")
	MPCGet, _ := phloGet.MultiPartyCall("b18848dd-e308-4758-bc5c-fb9c274d5fb7")
	callPayload := plivo.MultiPartyCallActionPayload{"cold_transfer",
		"919704583677",
		"agent",
		"919503364736"}

	callResponse, _ := MPCGet.ColdTransfer(callPayload)
	fmt.Println(callResponse)
}