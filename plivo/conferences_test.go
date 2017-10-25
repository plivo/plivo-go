package plivo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConferenceService_List(t *testing.T) {
	expectResponse("ConferenceListResponse.json", 200)

	if _, err := client.Conferences.IDList(); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Conferences.IDList()
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Conference")
}

func TestConferenceService_Get(t *testing.T) {
	expectResponse("ConferenceGetResponse.json", 200)
	ConferenceId := "ConferenceId"

	conference, err := client.Conferences.Get(ConferenceId)
	assert.Equal(t, conference.ID(), conference.ConferenceName)
	if err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err = client.Conferences.Get(ConferenceId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Conference/%s", ConferenceId)
}

func TestConferenceService_Delete(t *testing.T) {
	expectResponse("", 204)
	ConferenceId := "ConferenceId"

	if err := client.Conferences.Delete(ConferenceId); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	err := client.Conferences.Delete(ConferenceId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Conference/%s", ConferenceId)
}

func TestConferenceService_DeleteAll(t *testing.T) {
	expectResponse("", 204)

	if err := client.Conferences.DeleteAll(); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	err := client.Conferences.DeleteAll()
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Conference")
}

func TestConferenceService_MemberHangup(t *testing.T) {
	expectResponse("ConferenceMemberDeleteResponse.json", 204)
	ConferenceId := "ConferenceId"
	MemberId := "MemberId"

	if _, err := client.Conferences.MemberHangup(ConferenceId, MemberId); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Conferences.MemberHangup(ConferenceId, MemberId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Conference/%s/Member/%s", ConferenceId, MemberId)
}

func TestConferenceService_MemberKick(t *testing.T) {
	expectResponse("ConferenceMemberKickCreateResponse.json", 204)
	ConferenceId := "ConferenceId"
	MemberId := "MemberId"

	if _, err := client.Conferences.MemberKick(ConferenceId, MemberId); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Conferences.MemberKick(ConferenceId, MemberId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Conference/%s/Member/%s/Kick", ConferenceId, MemberId)
}

func TestConferenceService_MemberMute(t *testing.T) {
	expectResponse("ConferenceMemberMuteCreateResponse.json", 204)
	ConferenceId := "ConferenceId"
	MemberId := "MemberId"

	if _, err := client.Conferences.MemberMute(ConferenceId, MemberId); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Conferences.MemberMute(ConferenceId, MemberId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Conference/%s/Member/%s/Mute", ConferenceId, MemberId)
}

func TestConferenceService_MemberUnmute(t *testing.T) {
	expectResponse("", 204)
	ConferenceId := "ConferenceId"
	MemberId := "MemberId"

	if _, err := client.Conferences.MemberUnmute(ConferenceId, MemberId); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Conferences.MemberUnmute(ConferenceId, MemberId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Conference/%s/Member/%s/Mute", ConferenceId, MemberId)
}

func TestConferenceService_MemberDeaf(t *testing.T) {
	expectResponse("ConferenceMemberDeafCreateResponse.json", 204)
	ConferenceId := "ConferenceId"
	MemberId := "MemberId"

	if _, err := client.Conferences.MemberDeaf(ConferenceId, MemberId); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Conferences.MemberDeaf(ConferenceId, MemberId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Conference/%s/Member/%s/Deaf", ConferenceId, MemberId)
}

func TestConferenceService_MemberUndeaf(t *testing.T) {
	expectResponse("", 204)
	ConferenceId := "ConferenceId"
	MemberId := "MemberId"

	if _, err := client.Conferences.MemberUndeaf(ConferenceId, MemberId); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Conferences.MemberUndeaf(ConferenceId, MemberId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Conference/%s/Member/%s/Deaf", ConferenceId, MemberId)
}

func TestConferenceService_MemberSpeak(t *testing.T) {
	expectResponse("ConferenceMemberSpeakCreateResponse.json", 204)
	ConferenceId := "ConferenceId"
	MemberId := "MemberId"

	if _, err := client.Conferences.MemberSpeak(ConferenceId, MemberId, ConferenceMemberSpeakParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Conferences.MemberSpeak(ConferenceId, MemberId, ConferenceMemberSpeakParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Conference/%s/Member/%s/Speak", ConferenceId, MemberId)
}

func TestConferenceService_MemberSpeakStop(t *testing.T) {
	expectResponse("ConferenceMemberSpeakDeleteResponse.json", 204)
	ConferenceId := "ConferenceId"
	MemberId := "MemberId"

	if _, err := client.Conferences.MemberSpeakStop(ConferenceId, MemberId); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Conferences.MemberSpeakStop(ConferenceId, MemberId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Conference/%s/Member/%s/Speak", ConferenceId, MemberId)
}

func TestConferenceService_MemberPlay(t *testing.T) {
	expectResponse("ConferenceMemberPlayCreateResponse.json", 204)
	ConferenceId := "ConferenceId"
	MemberId := "MemberId"

	if _, err := client.Conferences.MemberPlay(ConferenceId, MemberId, "url"); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Conferences.MemberPlay(ConferenceId, MemberId, "url")
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Conference/%s/Member/%s/Play", ConferenceId, MemberId)
}

func TestConferenceService_MemberPlayStop(t *testing.T) {
	expectResponse("ConferenceMemberPlayDeleteResponse.json", 204)
	ConferenceId := "ConferenceId"
	MemberId := "MemberId"

	if _, err := client.Conferences.MemberPlayStop(ConferenceId, MemberId); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Conferences.MemberPlayStop(ConferenceId, MemberId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Conference/%s/Member/%s/Play", ConferenceId, MemberId)
}
