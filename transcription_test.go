package plivo

import (
	"encoding/json"
	"errors"
	"testing"
)

// without queryParam passed
func TestTranscriptionService_GetRecordingTranscription(t *testing.T) {
	expectResponse("getRecordingTranscription.json", 200)

	if _, err := client.Transcription.GetRecordingTranscription(GetRecordingTranscriptionRequest{
		TranscriptionID: "e12d05fe-6979-485c-83dc-9276114dba3b",
	}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Transcription.GetRecordingTranscription(GetRecordingTranscriptionRequest{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl
}

// with queryParam
func TestTranscriptionService_GetRecordingTranscriptionWithParam(t *testing.T) {
	expectResponse("getRecordingTranscription.json", 200)

	if _, err := client.Transcription.GetRecordingTranscription(GetRecordingTranscriptionRequest{
		TranscriptionID:   "e12d05fe-6979-485c-83dc-9276114dba3b",
		TranscriptionType: "raw",
	}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Transcription.GetRecordingTranscription(GetRecordingTranscriptionRequest{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl
}

func TestTranscriptionService_CreateRecordingTranscription(t *testing.T) {
	expectResponse("createRecordingTranscription.json", 200)
	var response map[string]interface{}
	if err := json.Unmarshal([]byte(expectedResponse), &response); err != nil {
		t.Fatalf("failed to unmarshal expected response: %v", err)
	}
	if _, err := client.Transcription.CreateRecordingTranscription(RecordingTranscriptionRequest{
		RecordingID: "e12d05fe-6979-485c-83dc-9276114dba3b",
	}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Transcription.CreateRecordingTranscription(RecordingTranscriptionRequest{})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl
}
