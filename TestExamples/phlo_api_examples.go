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

func testMultiPartyCallAbortTransfer() {
	phloGet, _ := phloClient.Phlos.Get("2e802f94-cc03-4e35-b4c8-5a0328544135")
	MPCGet, _ := phloGet.MultiPartyCall("b18848dd-e308-4758-bc5c-fb9c274d5fb7")
	callPayload := plivo.MultiPartyCallActionPayload{"abort_transfer",
		"919441456465",
		"agent",
		"919503364736"}

	callResponse, _ := MPCGet.AbortTransfer(callPayload)
	fmt.Println(callResponse)
}

func testMultiPartyCallMemberHold() {
	phloGet, _ := phloClient.Phlos.Get("2e802f94-cc03-4e35-b4c8-5a0328544135")
	MPCGet, _ := phloGet.MultiPartyCall("b18848dd-e308-4758-bc5c-fb9c274d5fb7")
	memberUpdateHold, err := MPCGet.Member("919704583677").Hold()
	fmt.Println(memberUpdateHold)
	if (err != nil) {
		fmt.Println(err)
	}
}

func testMultiPartyCallMemberUnHold() {
	phloGet, _ := phloClient.Phlos.Get("2e802f94-cc03-4e35-b4c8-5a0328544135")
	MPCGet, _ := phloGet.MultiPartyCall("b18848dd-e308-4758-bc5c-fb9c274d5fb7")
	memberUpdateUnHold, err := MPCGet.Member("919704583677").UnHold()
	fmt.Println(memberUpdateUnHold)
	if (err != nil) {
		fmt.Println(err)
	}
}

func testMultiPartyCallMemberVoiceMailDrop() {
	phloGet, _ := phloClient.Phlos.Get("2e802f94-cc03-4e35-b4c8-5a0328544135")
	MPCGet, _ := phloGet.MultiPartyCall("b18848dd-e308-4758-bc5c-fb9c274d5fb7")
	memberUpdateVoiceMailDrop, err := MPCGet.Member("919704583677").VoiceMailDrop()
	fmt.Println(memberUpdateVoiceMailDrop)
	if (err != nil) {
		fmt.Println(err)
	}
}

func testMultiPartyCallMemberHangUp() {
	phloGet, _ := phloClient.Phlos.Get("2e802f94-cc03-4e35-b4c8-5a0328544135")
	MPCGet, _ := phloGet.MultiPartyCall("b18848dd-e308-4758-bc5c-fb9c274d5fb7")
	memberUpdateHangUp, err := MPCGet.Member("919704583677").HangUp()
	fmt.Println(memberUpdateHangUp)
	if (err != nil) {
		fmt.Println(err)
	}
}

func testMultiPartyCallMemberResume() {
	phloGet, _ := phloClient.Phlos.Get("2e802f94-cc03-4e35-b4c8-5a0328544135")
	MPCGet, _ := phloGet.MultiPartyCall("b18848dd-e308-4758-bc5c-fb9c274d5fb7")
	memberUpdateResume, err := MPCGet.Member("919704583677").ResumeCall()
	fmt.Println(memberUpdateResume)
	if (err != nil) {
		fmt.Println(err)
	}
}
