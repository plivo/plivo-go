package plivo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPhlo_Get(t *testing.T) {
	expectResponse("PhloGetResoponse.json", 200)
	assert := require.New(t)
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
	assert.Nil(err)
	assertPhloRequest(t, "GET", "phlo/%s", testPhloId)
}

func TestNode_Get(t *testing.T) {
	expectResponse("NodeGetResponse.json", 200)
	assert := require.New(t)
	testPhloId := "PhloId"
	testNodeId := "NodeId"
	phloGet, _ := phloClient.Phlos.Get(testPhloId)

	cl := phloClient.httpClient
	phloClient.httpClient = nil

	_, err := phloGet.Node(testNodeId)
	if err == nil {
		phloClient.httpClient = cl
		panic(errors.New("error expected"))
	}

	phloClient.httpClient = cl
	_, err = phloGet.Node(testNodeId)
	assert.Nil(err)
	assertPhloRequest(t, "GET", "phlo/%s/%s/%s", phloGet.PhloId, nodeType, testNodeId)
}

func TestMultiPartyCall_Get(t *testing.T) {
	expectResponse("MultiPartyCallResponse.json", 200)
	assert := require.New(t)
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
	assert.Nil(err)
	assertPhloRequest(t, "GET", "phlo/%s/%s/%s", phloGet.PhloId, nodeType, testNodeId)
}

func TestPhloRun_Post(t *testing.T) {
	expectResponse("PhloRunResponse.json", 200)
	assert := require.New(t)
	testPhloId := "PhloId"
	phloGet, _ := phloClient.Phlos.Get(testPhloId)

	cl := phloClient.httpClient
	phloClient.httpClient = nil

	//expectResponse("PhloRunResponse.json", 200)
	_, err := phloGet.Run(nil)
	if err == nil {
		phloClient.httpClient = cl
		panic(errors.New("error expected"))
	}

	phloClient.httpClient = cl
	_, err = phloGet.Run(nil)
	assert.Nil(err)
	assertPhloRequest(t, "POST", "account/%s/phlo/%s/", phloClient.AuthId, phloGet.PhloId)
}
