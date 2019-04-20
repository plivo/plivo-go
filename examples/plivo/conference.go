package main

import (
	"fmt"
	"github.com/plivo/plivo-go"
	"encoding/json"
)


// Initialize the AuthId and AuthToken parameters in go
// To trigger Conference resource methods invoke corresponding helper method in main()
// Initialize the following the conferenceName and conferenceMemberId to trigger appropriate helper methods


// 	To build and run conference.go
//  cd  plivo-go/examples/plivo
//  go run conference.go credentials.go


const conferenceId  = "My Conf Room"
const conferenceMemberId = "27697"



func main(){
	testGetConference()
}

//Example to details of a particular conference.
//Pass conferenceId
func testGetConference(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Conferences.Get(
		"My Conf Room",
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to record conference
//Pass conference name and ConferenceRecordParams
func testRecordConference(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Conferences.Record(
		conferenceId,
		plivo.ConferenceRecordParams{
			30,
			"mp3",
			"auto",
			"transcription_url",
			"POST",
			"call_back_url",
			"GET",
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to stop recording a conference
//Pass conference name

func testStopConferenceRecord(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	err = client.Conferences.RecordStop(
		conferenceId,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Stopped successfully.")
}

//Example to delete conference
//Pass correponding conference name

func  testDeleteConference() {
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	err = client.Conferences.Delete(
		conferenceId,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted successfully.")
}

//Example to delete all conferences
//Pass correponding

func testDeleteAllConference() {
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	err = client.Conferences.DeleteAll()
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted all conferences successfully.")
}

//Example to list all conferences
//Pass correponding  list all conferences

func testConferenceList(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Conferences.IDList(
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to hangup a particular member present in a conference.
//Pass member id to hang-up

func testConferenceMemberHangup() {
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	_, err = client.Conferences.MemberHangup(
		conferenceId,
		conferenceMemberId,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted successfully.")
}

//Example to disconnect a member in a conference..
//Pass member id  and conference name params to hang-up

func testConferenceMemberKick(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Conferences.MemberKick(
		conferenceId,
		conferenceMemberId,
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to mute members in a conference.
//Pass member id  and conference name params

func testConferenceMemberMute(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Conferences.MemberMute(
		conferenceId,
		conferenceMemberId,
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to unmute members in a conference.
//Pass member id  and conference name params

func testConferenceMemberUnMute(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Conferences.MemberUnmute(
		conferenceId,
		conferenceMemberId,
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to deaf a particular member in the conference
//Pass member id  and conference name params

func testConferenceMemberDeaf(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Conferences.MemberDeaf(
		conferenceId,
		conferenceMemberId,
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to undeaf a particular member in conference
//Pass member id  and conference name params

func testConferenceMemberUnDeaf(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Conferences.MemberUndeaf(
		conferenceId,
		conferenceMemberId,
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to play music to  member in a conference
//Pass member id  and conference name and audio url params

func testConferenceMemberPlay(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Conferences.MemberPlay(
		conferenceId,
		conferenceMemberId,
		"https://s3.amazonaws.com/plivocloud/Trumpet.mp3",
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to stop play music to  member in a conference
//Pass member id  and conference name params

func  testConferenceMemberPlayStop() {
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Conferences.MemberPlayStop(
		conferenceId,
		conferenceMemberId,
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to play text to  member in a conference
//Pass member id  and conference name params

func testConferenceMemberSpeak(){
	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Conferences.MemberSpeak(
		conferenceId,
		conferenceMemberId,
		plivo.ConferenceMemberSpeakParams{
			Text: "Hello World!",
			Voice:"MAN",
			Language:"US English ",
		},
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}

//Example to stop play text to  member in a conference
//Pass member id  and conference name params

func testConferenceMemberSpeakStop(){

	client, err := plivo.NewClient(AuthId, AuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := client.Conferences.MemberSpeakStop(
		conferenceId,
		conferenceMemberId,
	)
	if err != nil {
		panic(err)
	}
	responseJSON, _ := json.Marshal(response)
	fmt.Println("Response : "+string(responseJSON))
}