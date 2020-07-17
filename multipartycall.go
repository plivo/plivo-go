package plivo

import "github.com/sirupsen/logrus"

type MultiPartyCallService struct {
	client *Client
}

type MultiPartyCallAddParticipant struct {
	role                     string
	friendlyName             string
	mpcUuid                  string
	from                     string
	to                       string
	callUuid                 string
	callStatusCallbackUrl    string
	callStatusCallbackMethod string
	sipHeaders               string
	confirmKey               string
	confirmKeySoundUrl       string
	confirmKeySoundMethod    string
	dialMusic                string
	ringTimeout              int64
	maxDuration              int64
	maxParticipants          int64
	waitMusicUrl             string
	waitMusicMethod          string
	agentHoldMusicUrl        string
	agentHoldMusicMethod     string
	customerHoldMusicUrl     string
	customerHoldMusicMethod  string
	recordingCallbackUrl     string
	recordingCallbackMethod  string
	statusCallbackUrl        string
	statusCallbackMethod     string
	onExitActionUrl          string
	onExitActionMethod       string
	record                   bool
	recordFileFormat         string
	statusCallbackEvents     string
	stayAlone                bool
	coachMode                bool
	mute                     bool
	hold                     bool
	startMpcOnEnter          bool
	endMpcOnExit             bool
	relayDtmfInputs          bool
	enterSound               string
	enterSoundMethod         string
	exitSound                string
	exitSoundMethod          string
}

type MultiPartyCallListParams struct {
	subAccount           string
	friendlyName         string
	status               string
	terminationCauseCode string
	endTimeGt            string
	endTimeGte           string
	endTimeLt            string
	endTimeLte           string
	endTime              string
	creationTimeGt       string
	creationTimeGte      string
	creationTimeLt       string
	creationTimeLte      string
	creationTime         string
	limit                string
	offset               string
}

type MultiPartyCallBasicParams struct {
	mpcUuid      string
	friendlyName string
}

type MultiPartyCallListParticipantParams struct {
	MultiPartyCallBasicParams
	callUuid string
}

type MultiPartyCallParticipantParams struct {
	MultiPartyCallBasicParams
	memberId string
}

type MultiPartyCallUpdateParticipantParams struct {
	MultiPartyCallParticipantParams
	coachMode bool
	mute      bool
	hold      bool
}

type MultiPartyCallStartRecordingParams struct {
	MultiPartyCallBasicParams
	fileFormat           string
	statusCallbackUrl    string
	statusCallbackMethod string
}

type MultiPartyCallListResponse struct {
	ApiID   string                      `json:"api_id" url:"api_id"`
	Meta    *Meta                       `json:"meta" url:"meta"`
	Objects []*MultiPartyCallListParams `json:"objects" url:"objects"`
}

type MultiPartyCallGetResponse struct {
	ApiID                string `json:"api_id" url:"api_id"`
	BilledAmount         string `json:"api_id" url:"api_id"`
	BilledDuration       string `json:"api_id" url:"api_id"`
	CreationTime         string `json:"api_id" url:"api_id"`
	Duration             string `json:"api_id" url:"api_id"`
	EndTime              string `json:"api_id" url:"api_id"`
	FriendlyName         string `json:"api_id" url:"api_id"`
	MpcUuid              string `json:"api_id" url:"api_id"`
	Participants         string `json:"api_id" url:"api_id"`
	Recording            string `json:"api_id" url:"api_id"`
	ResourceUri          string `json:"api_id" url:"api_id"`
	StartTime            string `json:"api_id" url:"api_id"`
	Status               string `json:"api_id" url:"api_id"`
	StayAlone            string `json:"api_id" url:"api_id"`
	SubAccount           string `json:"api_id" url:"api_id"`
	TerminationCause     string `json:"api_id" url:"api_id"`
	TerminationCauseCode string `json:"api_id" url:"api_id"`
}

type MultiPartyCallAddParticipantResponse struct {
	ApiID       string `json:"api_id" url:"api_id"`
	Message     string `json:"message,omitempty" url:"message,omitempty"`
	RequestUuid string `json:"request_uuid" url:"request_uuid"`
}

func (service *MultiPartyCallService) List(params MultiPartyCallListParams) (response *MultiPartyCallListResponse, err error) {
	req, err := service.client.NewRequest("GET", params, "MultiPartyCall")
	if err != nil {
		return
	}
	response = &MultiPartyCallListResponse{}
	err = service.client.ExecuteRequest(req, response, isVoiceRequest())
	return
}

func (service *MultiPartyCallService) Get(params MultiPartyCallBasicParams) (response *MultiPartyCallGetResponse, err error) {
	mpcId := MakeMPCId(params.mpcUuid, params.friendlyName)
	req, err := service.client.NewRequest("GET", nil, "MultiPartyCall/%s", mpcId)
	if err != nil {
		return
	}
	response = &MultiPartyCallGetResponse{}
	err = service.client.ExecuteRequest(req, response, isVoiceRequest())
	return
}

func (service *MultiPartyCallService) AddParticipant(params MultiPartyCallAddParticipant) (response *MultiPartyCallAddParticipantResponse, err error) {
	if (params.from != "" || params.to != "") && params.callUuid != "" {
		logrus.Fatal("cannot specify call_uuid when (from, to) is provided")
	}
	if params.from == "" && params.to == "" && params.callUuid == "" {
		logrus.Fatal("specify either call_uuid or (from, to)")
	}
	if params.callUuid == "" && (params.from == "" || params.to == "") {
		logrus.Fatal("specify (from, to) when not adding an existing call_uuid to multi party participant")
	}
	req, err := service.client.NewRequest("POST", params, "MultiPartyCall/%s/Participant", params.friendlyName)
	if err != nil {
		return
	}
	response = &MultiPartyCallAddParticipantResponse{}
	err = service.client.ExecuteRequest(req, response, isVoiceRequest())
	return
}

func (service *MultiPartyCallService) Start(params MultiPartyCallBasicParams) (response *MultiPartyCallAddParticipantResponse, err error) {
	mpcId := MakeMPCId(params.mpcUuid, params.friendlyName)
	req, err := service.client.NewRequest("POST", map[string]string{"status": "active"}, "MultiPartyCall/%s", mpcId)
	if err != nil {
		return
	}
	response = &MultiPartyCallAddParticipantResponse{}
	err = service.client.ExecuteRequest(req, response, isVoiceRequest())
	return
}

func (service *MultiPartyCallService) Stop(params MultiPartyCallBasicParams) (response *MultiPartyCallAddParticipantResponse, err error) {
	mpcId := MakeMPCId(params.mpcUuid, params.friendlyName)
	req, err := service.client.NewRequest("DELETE", nil, "MultiPartyCall/%s", mpcId)
	if err != nil {
		return
	}
	response = &MultiPartyCallAddParticipantResponse{}
	err = service.client.ExecuteRequest(req, response, isVoiceRequest())
	return
}

func (service *MultiPartyCallService) StartRecording(params MultiPartyCallStartRecordingParams) (response *MultiPartyCallAddParticipantResponse, err error) {
	mpcId := MakeMPCId(params.mpcUuid, params.friendlyName)
	req, err := service.client.NewRequest("POST", params, "MultiPartyCall/%s/Record", mpcId)
	if err != nil {
		return
	}
	response = &MultiPartyCallAddParticipantResponse{}
	err = service.client.ExecuteRequest(req, response, isVoiceRequest())
	return
}

func (service *MultiPartyCallService) StopRecording(params MultiPartyCallBasicParams) (response *MultiPartyCallAddParticipantResponse, err error) {
	mpcId := MakeMPCId(params.mpcUuid, params.friendlyName)
	req, err := service.client.NewRequest("DELETE", nil, "MultiPartyCall/%s/Record", mpcId)
	if err != nil {
		return
	}
	response = &MultiPartyCallAddParticipantResponse{}
	err = service.client.ExecuteRequest(req, response, isVoiceRequest())
	return
}

func (service *MultiPartyCallService) PauseRecording(params MultiPartyCallBasicParams) (response *MultiPartyCallAddParticipantResponse, err error) {
	mpcId := MakeMPCId(params.mpcUuid, params.friendlyName)
	req, err := service.client.NewRequest("POST", nil, "MultiPartyCall/%s/Record/Pause", mpcId)
	if err != nil {
		return
	}
	response = &MultiPartyCallAddParticipantResponse{}
	err = service.client.ExecuteRequest(req, response, isVoiceRequest())
	return
}

func (service *MultiPartyCallService) ResumeRecording(params MultiPartyCallBasicParams) (response *MultiPartyCallAddParticipantResponse, err error) {
	mpcId := MakeMPCId(params.mpcUuid, params.friendlyName)
	req, err := service.client.NewRequest("POST", nil, "MultiPartyCall/%s/Record/Resume", mpcId)
	if err != nil {
		return
	}
	response = &MultiPartyCallAddParticipantResponse{}
	err = service.client.ExecuteRequest(req, response, isVoiceRequest())
	return
}

func (service *MultiPartyCallService) ListParticipants(params MultiPartyCallListParticipantParams) (response *MultiPartyCallListResponse, err error) {
	mpcId := MakeMPCId(params.mpcUuid, params.friendlyName)
	req, err := service.client.NewRequest("GET", params, "MultiPartyCall/%s/Participant", mpcId)
	if err != nil {
		return
	}
	response = &MultiPartyCallListResponse{}
	err = service.client.ExecuteRequest(req, response, isVoiceRequest())
	return
}

func (service *MultiPartyCallService) UpdateParticipant(params MultiPartyCallUpdateParticipantParams) (response *MultiPartyCallAddParticipantResponse, err error) {
	mpcId := MakeMPCId(params.mpcUuid, params.friendlyName)
	req, err := service.client.NewRequest("POST", params, "MultiPartyCall/%s/Participant/%s", mpcId, params.memberId)
	if err != nil {
		return
	}
	response = &MultiPartyCallAddParticipantResponse{}
	err = service.client.ExecuteRequest(req, response, isVoiceRequest())
	return
}

func (service *MultiPartyCallService) KickParticipant(params MultiPartyCallParticipantParams) (response *MultiPartyCallAddParticipantResponse, err error) {
	mpcId := MakeMPCId(params.mpcUuid, params.friendlyName)
	req, err := service.client.NewRequest("DELETE", nil, "MultiPartyCall/%s/Participant/%s", mpcId, params.memberId)
	if err != nil {
		return
	}
	response = &MultiPartyCallAddParticipantResponse{}
	err = service.client.ExecuteRequest(req, response, isVoiceRequest())
	return
}

func (service *MultiPartyCallService) GetParticipant(params MultiPartyCallParticipantParams) (response *MultiPartyCallAddParticipantResponse, err error) {
	mpcId := MakeMPCId(params.mpcUuid, params.friendlyName)
	req, err := service.client.NewRequest("GET", nil, "MultiPartyCall/%s/Participant/%s", mpcId, params.memberId)
	if err != nil {
		return
	}
	response = &MultiPartyCallAddParticipantResponse{}
	err = service.client.ExecuteRequest(req, response, isVoiceRequest())
	return
}

func MakeMPCId(mpcUuid string, friendlyName string) string {
	mpcId := ""
	if mpcUuid != "" {
		mpcId = "uuid_" + mpcUuid
	} else if friendlyName != "" {
		mpcId = "name_" + friendlyName
	} else {
		logrus.Fatal("Need to specify a mpc_uuid or name")
	}
	return mpcId
}
