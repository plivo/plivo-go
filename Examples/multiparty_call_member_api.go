package main

import (
	"fmt"
)

func testMultiPartyCallAbortTransfer() {
	phloGet, _ := phloClient.Phlos.Get("2e802f94-cc03-4e35-b4c8-5a0328544135")
	MPCGet, _ := phloGet.MultiPartyCall("b18848dd-e308-4758-bc5c-fb9c274d5fb7")
	memberUpdateAbortTransfer, err := MPCGet.Member("919704583677").AbortTransfer()
	fmt.Println(memberUpdateAbortTransfer)
	if (err != nil) {
		fmt.Println(err)
	}
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