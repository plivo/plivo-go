package plivo

import (
	"errors"
	"testing"
)

func TestMultiPartyCallCall_Post(t *testing.T) {
	expectResponse("MultiPartyCallResponse.json", 200)
	testPhloId := "PhloId"
	testNodeId := "NodeId"
	phloGet, _ := phloClient.Phlos.Get(testPhloId)
	multiPartyCallGet, err := phloGet.MultiPartyCall(testNodeId)
	payload := PhloMultiPartyCallActionPayload{"testAction", "testTo", "testRole", "testSource"}

	cl := phloClient.httpClient
	phloClient.httpClient = nil
	expectResponse("MultiPartyCallCallResponse.json", 200)

	_, err = multiPartyCallGet.Call(payload)
	if err == nil {
		phloClient.httpClient = cl
		panic(errors.New("error expected"))
	}

	phloClient.httpClient = cl
	_, err = multiPartyCallGet.Call(payload)
	assertPhloRequest(t, "POST", "phlo/%s/%s/%s", phloGet.PhloId, nodeType, multiPartyCallGet.NodeID)
}



func TestMultiPartyCallWarmTransfer_Post(t *testing.T) {
	expectResponse("MultiPartyCallResponse.json", 200)
	testPhloId := "PhloId"
	testNodeId := "NodeId"
	phloGet, _ := phloClient.Phlos.Get(testPhloId)
	multiPartyCallGet, err := phloGet.MultiPartyCall(testNodeId)
	payload := PhloMultiPartyCallActionPayload{"testAction", "testTo", "testRole", "testSource"}

	cl := phloClient.httpClient
	phloClient.httpClient = nil
	expectResponse("MultiPartyCallWarmTransferResponse.json", 200)

	_, err = multiPartyCallGet.WarmTransfer(payload)
	if err == nil {
		phloClient.httpClient = cl
		panic(errors.New("error expected"))
	}

	phloClient.httpClient = cl
	_, err = multiPartyCallGet.WarmTransfer(payload)
	assertPhloRequest(t, "POST", "phlo/%s/%s/%s", phloGet.PhloId, nodeType, multiPartyCallGet.NodeID)
}

func TestMultiPartyCallColdTransfer_Post(t *testing.T) {
	expectResponse("MultiPartyCallResponse.json", 200)
	testPhloId := "PhloId"
	testNodeId := "NodeId"
	phloGet, _ := phloClient.Phlos.Get(testPhloId)
	multiPartyCallGet, err := phloGet.MultiPartyCall(testNodeId)
	payload := PhloMultiPartyCallActionPayload{"testAction", "testTo", "testRole", "testSource"}

	cl := phloClient.httpClient
	phloClient.httpClient = nil
	expectResponse("MultiPartyCallColdTransferResponse.json", 200)

	_, err = multiPartyCallGet.ColdTransfer(payload)
	if err == nil {
		phloClient.httpClient = cl
		panic(errors.New("error expected"))
	}

	phloClient.httpClient = cl
	_, err = multiPartyCallGet.ColdTransfer(payload)
	assertPhloRequest(t, "POST", "phlo/%s/%s/%s", phloGet.PhloId, nodeType, multiPartyCallGet.NodeID)
}

func TestMultiPartyCallMemberAbortTransfer_Post(t *testing.T) {
	expectResponse("MultiPartyCallResponse.json", 200)
	testPhloId := "PhloId"
	testNodeId := "NodeId"
	testMemberAdress := "testMemnber"
	phloGet, _ := phloClient.Phlos.Get(testPhloId)
	multiPartyCallGet, err := phloGet.MultiPartyCall(testNodeId)


	cl := phloClient.httpClient
	phloClient.httpClient = nil
	expectResponse("multiPartyCallMemberActionResponse.json", 200)

	_, err = multiPartyCallGet.Member(testMemberAdress).AbortTransfer()
	if err == nil {
		phloClient.httpClient = cl
		panic(errors.New("error expected"))
	}

	phloClient.httpClient = cl
	_, err = multiPartyCallGet.Member(testMemberAdress).AbortTransfer()
	assertPhloRequest(t, "POST", "phlo/%s/%s/%s/members/%s", phloGet.PhloId, nodeType, multiPartyCallGet.NodeID, testMemberAdress)
}

func TestMultiPartyCallMemberHangUp_Post(t *testing.T) {
	expectResponse("MultiPartyCallResponse.json", 200)
	testPhloId := "PhloId"
	testNodeId := "NodeId"
	testMemberAdress := "testMemnber"
	phloGet, _ := phloClient.Phlos.Get(testPhloId)
	multiPartyCallGet, err := phloGet.MultiPartyCall(testNodeId)

	cl := phloClient.httpClient
	phloClient.httpClient = nil
	expectResponse("multiPartyCallMemberActionResponse.json", 200)

	_, err = multiPartyCallGet.Member(testMemberAdress).Hold()
	if err == nil {
		phloClient.httpClient = cl
		panic(errors.New("error expected"))
	}

	phloClient.httpClient = cl
	_, err = multiPartyCallGet.Member(testMemberAdress).Hold()
	assertPhloRequest(t, "POST", "phlo/%s/%s/%s/members/%s", phloGet.PhloId, nodeType, multiPartyCallGet.NodeID, testMemberAdress)
}

func TestMultiPartyCallMemberVoiceMailDrop_Post(t *testing.T) {
	expectResponse("MultiPartyCallResponse.json", 200)
	testPhloId := "PhloId"
	testNodeId := "NodeId"
	testMemberAdress := "testMemnber"
	phloGet, _ := phloClient.Phlos.Get(testPhloId)
	multiPartyCallGet, err := phloGet.MultiPartyCall(testNodeId)

	cl := phloClient.httpClient
	phloClient.httpClient = nil
	expectResponse("multiPartyCallMemberActionResponse.json", 200)

	_, err = multiPartyCallGet.Member(testMemberAdress).VoiceMailDrop()
	if err == nil {
		phloClient.httpClient = cl
		panic(errors.New("error expected"))
	}

	phloClient.httpClient = cl
	_, err = multiPartyCallGet.Member(testMemberAdress).VoiceMailDrop()
	assertPhloRequest(t, "POST", "phlo/%s/%s/%s/members/%s", phloGet.PhloId, nodeType, multiPartyCallGet.NodeID, testMemberAdress)
}

func TestMultiPartyCallMemberHold_Post(t *testing.T) {
	expectResponse("MultiPartyCallResponse.json", 200)
	testPhloId := "PhloId"
	testNodeId := "NodeId"
	testMemberAdress := "testMemnber"
	phloGet, _ := phloClient.Phlos.Get(testPhloId)
	multiPartyCallGet, err := phloGet.MultiPartyCall(testNodeId)

	cl := phloClient.httpClient
	phloClient.httpClient = nil
	expectResponse("multiPartyCallMemberActionResponse.json", 200)

	_, err = multiPartyCallGet.Member(testMemberAdress).Hold()
	if err == nil {
		phloClient.httpClient = cl
		panic(errors.New("error expected"))
	}

	phloClient.httpClient = cl
	_, err = multiPartyCallGet.Member(testMemberAdress).Hold()
	assertPhloRequest(t, "POST", "phlo/%s/%s/%s/members/%s", phloGet.PhloId, nodeType, multiPartyCallGet.NodeID, testMemberAdress)
}

func TestMultiPartyCallMemberResumeCall_Post(t *testing.T) {
	expectResponse("MultiPartyCallResponse.json", 200)
	testPhloId := "PhloId"
	testNodeId := "NodeId"
	testMemberAdress := "testMemnber"
	phloGet, _ := phloClient.Phlos.Get(testPhloId)
	multiPartyCallGet, err := phloGet.MultiPartyCall(testNodeId)

	cl := phloClient.httpClient
	phloClient.httpClient = nil
	expectResponse("multiPartyCallMemberActionResponse.json", 200)

	_, err = multiPartyCallGet.Member(testMemberAdress).ResumeCall()
	if err == nil {
		phloClient.httpClient = cl
		panic(errors.New("error expected"))
	}

	phloClient.httpClient = cl
	_, err = multiPartyCallGet.Member(testMemberAdress).ResumeCall()
	assertPhloRequest(t, "POST", "phlo/%s/%s/%s/members/%s", phloGet.PhloId, nodeType, multiPartyCallGet.NodeID, testMemberAdress)
}

func TestMultiPartyCallMemberUnHold_Post(t *testing.T) {
	expectResponse("MultiPartyCallResponse.json", 200)
	testPhloId := "PhloId"
	testNodeId := "NodeId"
	testMemberAdress := "testMemnber"
	phloGet, _ := phloClient.Phlos.Get(testPhloId)
	multiPartyCallGet, err := phloGet.MultiPartyCall(testNodeId)

	cl := phloClient.httpClient
	phloClient.httpClient = nil
	expectResponse("multiPartyCallMemberActionResponse.json", 200)

	_, err = multiPartyCallGet.Member(testMemberAdress).UnHold()
	if err == nil {
		phloClient.httpClient = cl
		panic(errors.New("error expected"))
	}

	phloClient.httpClient = cl
	_, err = multiPartyCallGet.Member(testMemberAdress).UnHold()
	assertPhloRequest(t, "POST", "phlo/%s/%s/%s/members/%s", phloGet.PhloId, nodeType, multiPartyCallGet.NodeID, testMemberAdress)
}
