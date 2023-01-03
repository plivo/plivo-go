package plivo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCallService_Create(t *testing.T) {
	expectResponse("CallCreateResponse.json", 202)

	if _, err := client.Calls.Create(CallCreateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Calls.Create(CallCreateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Call")
}

func TestCallService_Update(t *testing.T) {
	expectResponse("CallUpdateResponse.json", 202)
	CallId := "CallId"

	if _, err := client.Calls.Update(CallId, CallUpdateParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Calls.Update(CallId, CallUpdateParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Call/%s", CallId)
}

func TestCallService_List(t *testing.T) {
	expectResponse("CallListResponse.json", 200)

	if _, err := client.Calls.List(CallListParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Calls.List(CallListParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Call")
}

func TestCallService_Get(t *testing.T) {
	expectResponse("CallGetResponse.json", 200)
	CallId := "CallId"

	call, err := client.Calls.Get(CallId)
	assert.Equal(t, call.ID(), call.CallUUID)
	if err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err = client.Calls.Get(CallId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Call/%s", CallId)
}

func TestCallService_Delete(t *testing.T) {
	expectResponse("", 204)
	CallId := "CallId"

	if err := client.Calls.Delete(CallId); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	err := client.Calls.Delete(CallId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Call/%s", CallId)
}

func TestLiveCallService_List(t *testing.T) {
	expectResponse("LiveCallListGetResponse.json", 200)

	if _, err := client.LiveCalls.IDList(LiveCallFilters{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.LiveCalls.IDList(LiveCallFilters{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Call")
}

func TestLiveCallService_Get(t *testing.T) {
	expectResponse("LiveCallGetResponse.json", 200)
	LiveCallId := "LiveCallId"

	liveCall, err := client.LiveCalls.Get(LiveCallId)
	assert.Equal(t, liveCall.ID(), liveCall.CallUUID)
	if err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err = client.LiveCalls.Get(LiveCallId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Call/%s", LiveCallId)
}

func TestCallService_Record(t *testing.T) {
	expectResponse("liveCallRecordCreateResponse.json", 202)
	CallID := "CallId"

	if _, err := client.Calls.Record(CallID, CallRecordParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Calls.Record(CallID, CallRecordParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Call/%s/Record", CallID)
}

func TestCallService_StopRecording(t *testing.T) {
	expectResponse("", 204)
	CallID := "CallId"

	if err := client.Calls.StopRecording(CallID); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	err := client.Calls.StopRecording(CallID)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Call/%s/Record", CallID)
}

func TestCallService_Speak(t *testing.T) {
	expectResponse("liveCallSpeakCreateResponse.json", 202)
	CallID := "CallId"

	if _, err := client.Calls.Speak(CallID, CallSpeakParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Calls.Speak(CallID, CallSpeakParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Call/%s/Speak", CallID)
}

func TestCallService_StopSpeaking(t *testing.T) {
	expectResponse("", 204)
	CallID := "CallId"

	if err := client.Calls.StopSpeaking(CallID); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	err := client.Calls.StopSpeaking(CallID)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Call/%s/Speak", CallID)
}

func TestCallService_Play(t *testing.T) {
	expectResponse("liveCallPlayCreateResponse.json", 202)
	CallID := "CallId"

	if _, err := client.Calls.Play(CallID, CallPlayParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Calls.Play(CallID, CallPlayParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Call/%s/Play", CallID)
}

func TestCallService_StopPlaying(t *testing.T) {
	expectResponse("", 204)
	CallID := "CallId"

	if err := client.Calls.StopPlaying(CallID); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	err := client.Calls.StopPlaying(CallID)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Call/%s/Play", CallID)
}

func TestCallService_SendDigits(t *testing.T) {
	expectResponse("liveCallDtmfCreateResponse.json", 202)
	CallID := "CallId"

	if _, err := client.Calls.SendDigits(CallID, CallDTMFParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Calls.SendDigits(CallID, CallDTMFParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Call/%s/DTMF", CallID)
}

func TestCallService_CancelRequest(t *testing.T) {
	expectResponse("", 204)
	RequestID := "RequestID"

	if err := client.Calls.CancelRequest(RequestID); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	err := client.Calls.CancelRequest(RequestID)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Request/%s", RequestID)
}

func TestCallService_CreateMaskingSession(t *testing.T) {
	expectResponse("createMaskingSessionResponse.json", 200)

	if _, err := client.Calls.CreateMaskingSession(CreateMaskingSessionParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Calls.CreateMaskingSession(CreateMaskingSessionParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Session")
}

func TestCallService_DeleteMaskingSession(t *testing.T) {
	expectResponse("deleteMaskingSessionResponse.json", 204)
	SessionUuid := "15e4256c-be01-475c-9a69-95cf65bbed71"

	if _, err := client.Calls.DeleteMaskingSession(SessionUuid); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Calls.DeleteMaskingSession(SessionUuid)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Session/%s", SessionUuid)
}
