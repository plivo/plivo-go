package plivo

import (
	"errors"
	"testing"
)

func TestPhlo_Get(t *testing.T) {
	expectResponse("PhloGetResoponse.json", 200)
	testPhloId := "PhloId"
	_, err := phloClient.Phlos.Get("")
	if err == nil {
		panic(errors.New("error expected"))
	}

	cl := phloClient.httpClient
	phloClient.httpClient = nil
	_, err = phloClient.Phlos.Get(testPhloId)
	if err == nil {
		phloClient.httpClient = cl
		panic(errors.New("error expected"))
	}
	phloClient.httpClient = cl
	_, err = phloClient.Phlos.Get(testPhloId)
	assertPhloRequest(t, "GET", "phlo/%s", testPhloId)
}

func TestNode_Get(t *testing.T) {
	expectResponse("NodeGetResponse.json", 200)
	testPhloId := "PhloId"
	testNodeId := "NodeId"
	phloGet, _ := phloClient.Phlos.Get(testPhloId)

	_, err := phloGet.Node("")
	if err == nil {
		panic(errors.New("error expected"))
	}

	cl := phloClient.httpClient
	phloClient.httpClient = nil

	_, err = phloGet.Node(testNodeId)
	if err == nil {
		phloClient.httpClient = cl
		panic(errors.New("error expected"))
	}

	phloClient.httpClient = cl
	_, err = phloGet.Node(testNodeId)
	assertPhloRequest(t, "GET", "phlo/%s/%s/%s", phloGet.PhloId, nodeType, testNodeId)
}

func TestMultiPartyCall_Get(t *testing.T) {
	expectResponse("MultiPartyCallResponse.json", 200)
	testPhloId := "PhloId"
	testNodeId := "NodeId"
	phloGet, _ := phloClient.Phlos.Get(testPhloId)

	cl := phloClient.httpClient
	phloClient.httpClient = nil

	_, err := phloGet.MultiPartyCall(testNodeId)
	if err == nil {
		phloClient.httpClient = cl
		panic(errors.New("error expected"))
	}

	phloClient.httpClient = cl
	_, err = phloGet.MultiPartyCall(testNodeId)
	assertPhloRequest(t, "GET", "phlo/%s/%s/%s", phloGet.PhloId, nodeType, testNodeId)
}

func TestPhloRun_Post(t *testing.T) {
	expectResponse("PhloRunResponse.json", 200)
	testPhloId := "PhloId"
	phloGet, _ := phloClient.Phlos.Get(testPhloId)

	cl := phloClient.httpClient
	phloClient.httpClient = nil

	_, err := phloGet.Run(nil)
	if err == nil {
		phloClient.httpClient = cl
		panic(errors.New("error expected"))
	}

	phloClient.httpClient = cl
	_, err = phloGet.Run(nil)
	assertPhloRequest(t, "POST", "account/%s/phlo/%s/", phloClient.AuthId, phloGet.PhloId)
}
