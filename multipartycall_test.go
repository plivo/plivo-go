package plivo

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestMPCService_List(t *testing.T) {
	expectResponse("MPCListResponse.json", 200)

	if _, err := client.MultiPartyCall.List(MultiPartyCallListParams{}); err != nil {
		panic(err)
	}

	assertRequest(t, "GET", "MultiPartyCall")
}

func TestMPCService_Get(t *testing.T) {
	expectResponse("MPCGetResponse.json", 200)

	if _, err := client.MultiPartyCall.Get(MultiPartyCallBasicParams{"b1e37275-b8e2-42cd-ae63-fffcc54a50b5", "thanos_test"}); err != nil {
		panic(err)
	}

	assertRequest(t, "GET", "MultiPartyCall/%s", "uuid_b1e37275-b8e2-42cd-ae63-fffcc54a50b5")
}

func TestMPCService_AddParticipant(t *testing.T) {
	expectResponse("MPCAddParticipantResponse.json", 201)

	if response, err := client.MultiPartyCall.AddParticipant(MultiPartyCallAddParticipant{to: "917013835803", from: "918888888888", role: "agent", friendlyName: "name_thanos"}); err != nil {
		panic(err)
	} else {
		logrus.Info(response)
	}

	assertRequest(t, "POST", "MultiPartyCall/%s/Participant", "name_thanos")
}

func TestMPCService_Start(t *testing.T) {
	//expectResponse("MPCStartResponse.json", 200)

	if _, err := client.MultiPartyCall.Start(MultiPartyCallBasicParams{"ebacced2-21ab-466d-9df4-67339991761b", "thanos_test"}); err != nil {
		panic(err)
	}

	assertRequest(t, "POST", "MultiPartyCall/%s", "uuid_ebacced2-21ab-466d-9df4-67339991761b")
}

func TestMPCService_Stop(t *testing.T) {
	//expectResponse("MPCStopResponse.json", 200)

	if _, err := client.MultiPartyCall.Stop(MultiPartyCallBasicParams{"ebacced2-21ab-466d-9df4-67339991761b", "thanos_test"}); err != nil {
		panic(err)
	}

	assertRequest(t, "DELETE", "MultiPartyCall/%s", "uuid_ebacced2-21ab-466d-9df4-67339991761b")
}

func TestMPCService_StartRecord(t *testing.T) {
	expectResponse("MPCStartRecordResponse.json", 200)

	if _, err := client.MultiPartyCall.StartRecording(MultiPartyCallStartRecordingParams{MultiPartyCallBasicParams{"ebacced2-21ab-466d-9df4-67339991761b", "thanos_test"}, "wav", "https://www.google.com", "GET"}); err != nil {
		panic(err)
	}

	assertRequest(t, "POST", "MultiPartyCall/%s/Record", "uuid_ebacced2-21ab-466d-9df4-67339991761b")
}

func TestMPCService_StopRecord(t *testing.T) {
	if _, err := client.MultiPartyCall.StopRecording(MultiPartyCallBasicParams{"ebacced2-21ab-466d-9df4-67339991761b", "name_thanos"}); err != nil {
		panic(err)
	}

	assertRequest(t, "DELETE", "MultiPartyCall/%s/Record", "uuid_ebacced2-21ab-466d-9df4-67339991761b")
}

func TestMPCService_ResumeRecord(t *testing.T) {
	if _, err := client.MultiPartyCall.ResumeRecording(MultiPartyCallBasicParams{"ebacced2-21ab-466d-9df4-67339991761b", "name_thanos"}); err != nil {
		panic(err)
	}

	assertRequest(t, "POST", "MultiPartyCall/%s/Record/Resume", "uuid_ebacced2-21ab-466d-9df4-67339991761b")
}

func TestMPCService_PauseRecord(t *testing.T) {
	//expectResponse("MPCPauseRecordResponse.json", 200)

	if _, err := client.MultiPartyCall.PauseRecording(MultiPartyCallBasicParams{"ebacced2-21ab-466d-9df4-67339991761b", "name_thanos"}); err != nil {
		panic(err)
	}

	assertRequest(t, "POST", "MultiPartyCall/%s/Record/Pause", "uuid_ebacced2-21ab-466d-9df4-67339991761b")
}

func TestMPCService_ListParticipants(t *testing.T) {
	expectResponse("MPCListParticipantsResponse.json", 200)

	if _, err := client.MultiPartyCall.ListParticipants(MultiPartyCallListParticipantParams{MultiPartyCallBasicParams{"ebacced2-21ab-466d-9df4-67339991761b", "name_thanos"}, "call_uuid"}); err != nil {
		panic(err)
	}

	assertRequest(t, "GET", "MultiPartyCall/%s/Participant", "uuid_ebacced2-21ab-466d-9df4-67339991761b")
}

func TestMPCService_UpdateParticipant(t *testing.T) {
	expectResponse("MPCUpdateParticipantResponse.json", 200)

	if _, err := client.MultiPartyCall.UpdateParticipant(MultiPartyCallUpdateParticipantParams{MultiPartyCallParticipantParams{MultiPartyCallBasicParams{"ebacced2-21ab-466d-9df4-67339991761b", "name_thanos"}, "209"}, false, false, true}); err != nil {
		panic(err)
	}

	assertRequest(t, "POST", "MultiPartyCall/%s/Participant/%s", "uuid_ebacced2-21ab-466d-9df4-67339991761b", "209")
}

func TestMPCService_KickParticipant(t *testing.T) {
	if _, err := client.MultiPartyCall.KickParticipant(MultiPartyCallParticipantParams{MultiPartyCallBasicParams{"ebacced2-21ab-466d-9df4-67339991761b", "name_thanos"}, "209"}); err != nil {
		panic(err)
	}

	assertRequest(t, "DELETE", "MultiPartyCall/%s/Participant/%s", "uuid_ebacced2-21ab-466d-9df4-67339991761b", "209")
}

func TestMPCService_GetParticipant(t *testing.T) {
	expectResponse("MPCGetParticipantResponse.json", 200)

	if _, err := client.MultiPartyCall.GetParticipant(MultiPartyCallParticipantParams{MultiPartyCallBasicParams{"ebacced2-21ab-466d-9df4-67339991761b", "name_thanos"}, "209"}); err != nil {
		panic(err)
	}

	assertRequest(t, "GET", "MultiPartyCall/%s/Participant/%s", "uuid_ebacced2-21ab-466d-9df4-67339991761b", "209")
}
