package plivo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecordingService_List(t *testing.T) {
	expectResponse("RecordingListResponse.json", 200)

	if _, err := client.Recordings.List(RecordingListParams{}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Recordings.List(RecordingListParams{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Recording")
}

func TestRecordingService_Get(t *testing.T) {
	expectResponse("RecordingGetResponse.json", 200)
	RecordingId := "RecordingId"

	recording, err := client.Recordings.Get(RecordingId)
	assert.Equal(t, recording.ID(), recording.RecordingID)
	if err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err = client.Recordings.Get(RecordingId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Recording/%s", RecordingId)
}
func TestRecordingService_Delete(t *testing.T) {
	expectResponse("", 204)
	RecordingId := "RecordingId"

	if err := client.Recordings.Delete(RecordingId); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	err := client.Recordings.Delete(RecordingId)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "DELETE", "Recording/%s", RecordingId)
}
