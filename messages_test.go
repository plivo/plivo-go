package plivo

import (
	"errors"
	"testing"
)

func TestMessageService_List(t *testing.T) {
	expectResponse("messageListResponse.json", 200)

	if _, err := client.Messages.List(MessageListParams{}); err != nil {
		panic(err)
	}
	cl := client.httpClient
	client.httpClient = nil
	if _, err := client.Messages.List(MessageListParams{}); err == nil {
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Message")
}

func TestMessageService_Get(t *testing.T) {
	expectResponse("messageGetResponse.json", 200)
	uuid := "5b40a428-bfc7-4daf-9d06-726c558bf3b8"

	if _, err := client.Messages.Get(uuid); err != nil {
		panic(err)
	}
	cl := client.httpClient
	client.httpClient = nil
	if _, err := client.Messages.Get(uuid); err == nil {
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Message/%s", uuid)
}

func TestMessageService_GetMedia(t *testing.T) {
	expectResponse("mmsMediaGetResponse.json", 200)
	uuid := "f734eeec-e59f-11e9-89dc-0242ac110003"
	mediaID := "91f8f9ce-e236-474f-a2b8-f75d9c35988a"

	if _, err := client.Messages.GetMedia(GetMediaParams{
		MessageUUID: uuid,
		MediaID:     mediaID,
	}); err != nil {
		panic(err)
	}
	cl := client.httpClient
	client.httpClient = nil
	if _, err := client.Messages.GetMedia(GetMediaParams{
		MessageUUID: uuid,
		MediaID:     mediaID,
	}); err == nil {
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Message/%s/Media/%s", uuid, mediaID)
}

func TestMessageService_ListMedia(t *testing.T) {
	expectResponse("mmsMediaListResponse.json", 200)
	uuid := "5b40a428-bfc7-4daf-9d06-726c558bf3b8"

	if _, err := client.Messages.ListMedia(uuid); err != nil {
		panic(err)
	}
	cl := client.httpClient
	client.httpClient = nil
	if _, err := client.Messages.ListMedia(uuid); err == nil {
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Message/%s/Media", uuid)
}

func TestMessageService_Create(t *testing.T) {
	expectResponse("messageSendResponse.json", 202)

	if _, err := client.Messages.Create(MessageCreateParams{
		Src:  "+911231231230",
		Dst:  "+911231231231",
		Text: "Testing",
	}); err != nil {
		panic(err)
	}
	cl := client.httpClient
	client.httpClient = nil
	if _, err := client.Messages.Create(MessageCreateParams{
		Src:  "+911231231230",
		Dst:  "+911231231231",
		Text: "Testing",
	}); err == nil {
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "POST", "Message")
}
