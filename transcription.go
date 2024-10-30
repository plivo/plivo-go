package plivo

type TranscriptionService struct {
	client *Client
}

type GetRecordingTranscriptionRequest struct {
	TranscriptionID   string `json:"transcription_id"`
	TranscriptionType string `json:"type"`
}

type GetRecordingTranscriptionParams struct {
	Type string `url:"type"`
}

type GetRecordingTranscriptionResponse struct {
	APIID               string      `json:"api_id"`
	Cost                float64     `json:"cost"`
	Rate                float64     `json:"rate"`
	RecordingDurationMs float64     `json:"recording_duration_ms"`
	RecordingStartMs    float64     `json:"recording_start_ms"`
	Status              string      `json:"status"`
	Transcription       interface{} `json:"transcription"`
}

func (service *TranscriptionService) GetRecordingTranscription(request GetRecordingTranscriptionRequest) (response *GetRecordingTranscriptionResponse, err error) {
	params := GetRecordingTranscriptionParams{
		Type: request.TranscriptionType,
	}
	req, err := service.client.NewRequest("GET", params, "Transcription/%s", request.TranscriptionID)
	if err != nil {
		return
	}
	response = &GetRecordingTranscriptionResponse{}
	err = service.client.ExecuteRequest(req, response, isVoiceRequest())
	return
}
