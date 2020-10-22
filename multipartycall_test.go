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

	if _, err := client.MultiPartyCall.Get(MultiPartyCallBasicParams{MpcUuid: "b1e37275-b8e2-42cd-ae63-fffcc54a50b5"}); err != nil {
		panic(err)
	}

	assertRequest(t, "GET", "MultiPartyCall/%s", "uuid_b1e37275-b8e2-42cd-ae63-fffcc54a50b5")
}

func TestMPCService_AddParticipant(t *testing.T) {
	expectResponse("MPCAddParticipantResponse.json", 201)

	if response, err := client.MultiPartyCall.AddParticipant(MultiPartyCallBasicParams{FriendlyName: "thanos"}, MultiPartyCallAddParticipantParams{To: "917013835803", From: "918888888888", Role: "agent"}); err != nil {
		panic(err)
	} else {
		logrus.Info(response)
	}

	assertRequest(t, "POST", "MultiPartyCall/%s/Participant", "name_thanos")
}

func TestMPCService_Start(t *testing.T) {
	expectResponse("", 204)

	if err := client.MultiPartyCall.Start(MultiPartyCallBasicParams{MpcUuid: "ebacced2-21ab-466d-9df4-67339991761b"}); err != nil {
		panic(err)
	}

	assertRequest(t, "POST", "MultiPartyCall/%s", "uuid_ebacced2-21ab-466d-9df4-67339991761b")
}

func TestMPCService_Stop(t *testing.T) {
	expectResponse("", 204)

	if err := client.MultiPartyCall.Stop(MultiPartyCallBasicParams{MpcUuid: "ebacced2-21ab-466d-9df4-67339991761b"}); err != nil {
		panic(err)
	}

	assertRequest(t, "DELETE", "MultiPartyCall/%s", "uuid_ebacced2-21ab-466d-9df4-67339991761b")
}

func TestMPCService_StartRecord(t *testing.T) {
	expectResponse("MPCStartRecordResponse.json", 200)

	if _, err := client.MultiPartyCall.StartRecording(MultiPartyCallBasicParams{MpcUuid: "ebacced2-21ab-466d-9df4-67339991761b"}, MultiPartyCallStartRecordingParams{FileFormat: "wav", StatusCallbackUrl: "https://www.google.com", StatusCallbackMethod: "GET"}); err != nil {
		panic(err)
	}

	assertRequest(t, "POST", "MultiPartyCall/%s/Record", "uuid_ebacced2-21ab-466d-9df4-67339991761b")
}

func TestMPCService_StopRecord(t *testing.T) {
	expectResponse("", 204)
	if err := client.MultiPartyCall.StopRecording(MultiPartyCallBasicParams{MpcUuid: "ebacced2-21ab-466d-9df4-67339991761b"}); err != nil {
		panic(err)
	}

	assertRequest(t, "DELETE", "MultiPartyCall/%s/Record", "uuid_ebacced2-21ab-466d-9df4-67339991761b")
}

func TestMPCService_ResumeRecord(t *testing.T) {
	expectResponse("", 204)
	if err := client.MultiPartyCall.ResumeRecording(MultiPartyCallBasicParams{MpcUuid: "ebacced2-21ab-466d-9df4-67339991761b"}); err != nil {
		panic(err)
	}

	assertRequest(t, "POST", "MultiPartyCall/%s/Record/Resume", "uuid_ebacced2-21ab-466d-9df4-67339991761b")
}

func TestMPCService_PauseRecord(t *testing.T) {
	expectResponse("", 204)

	if err := client.MultiPartyCall.PauseRecording(MultiPartyCallBasicParams{MpcUuid: "ebacced2-21ab-466d-9df4-67339991761b"}); err != nil {
		panic(err)
	}

	assertRequest(t, "POST", "MultiPartyCall/%s/Record/Pause", "uuid_ebacced2-21ab-466d-9df4-67339991761b")
}

func TestMPCService_ListParticipants(t *testing.T) {
	expectResponse("MPCListParticipantsResponse.json", 200)

	if _, err := client.MultiPartyCall.ListParticipants(MultiPartyCallBasicParams{MpcUuid: "ebacced2-21ab-466d-9df4-67339991761b"}, MultiPartyCallListParticipantParams{CallUuid: "call_uuid"}); err != nil {
		panic(err)
	}

	assertRequest(t, "GET", "MultiPartyCall/%s/Participant", "uuid_ebacced2-21ab-466d-9df4-67339991761b")
}

func TestMPCService_UpdateParticipant(t *testing.T) {
	expectResponse("MPCUpdateParticipantResponse.json", 202)
	coachMode := false
	mute := true
	hold := false
	if _, err := client.MultiPartyCall.UpdateParticipant(MultiPartyCallParticipantParams{MpcUuid: "ebacced2-21ab-466d-9df4-67339991761b", ParticipantId: "209"}, MultiPartyCallUpdateParticipantParams{CoachMode: &coachMode, Mute: &mute, Hold: &hold}); err != nil {
		panic(err)
	}

	assertRequest(t, "POST", "MultiPartyCall/%s/Participant/%s", "uuid_ebacced2-21ab-466d-9df4-67339991761b", "209")
}

func TestMPCService_KickParticipant(t *testing.T) {
	expectResponse("", 204)
	if err := client.MultiPartyCall.KickParticipant(MultiPartyCallParticipantParams{MpcUuid: "ebacced2-21ab-466d-9df4-67339991761b", ParticipantId: "209"}); err != nil {
		panic(err)
	}

	assertRequest(t, "DELETE", "MultiPartyCall/%s/Participant/%s", "uuid_ebacced2-21ab-466d-9df4-67339991761b", "209")
}

func TestMPCService_GetParticipant(t *testing.T) {
	expectResponse("MPCGetParticipantResponse.json", 200)

	if _, err := client.MultiPartyCall.GetParticipant(MultiPartyCallParticipantParams{MpcUuid: "ebacced2-21ab-466d-9df4-67339991761b", ParticipantId: "209"}); err != nil {
		panic(err)
	}

	assertRequest(t, "GET", "MultiPartyCall/%s/Participant/%s", "uuid_ebacced2-21ab-466d-9df4-67339991761b", "209")
}
