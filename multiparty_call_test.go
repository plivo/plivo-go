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
	payload := MultiPartyCallActionPayload{"testAction", "testTo", "testRole", "testSource"}

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
	payload := MultiPartyCallActionPayload{"testAction", "testTo", "testRole", "testSource"}

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
	payload := MultiPartyCallActionPayload{"testAction", "testTo", "testRole", "testSource"}

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
