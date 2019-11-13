package xml

import (
	"encoding/xml"

	"github.com/sirupsen/logrus"
)

type ResponseElement struct {
	XMLName  xml.Name      `xml:"Response"`
	Contents []interface{} `xml:",xmlvalue"`
}

func (element ResponseElement) String() string {
	bytes, err := xml.Marshal(element)
	if err != nil {
		logrus.Errorln(err.Error())
	}
	return string(bytes)
}

type ConferenceElement struct {
	Contents string `xml:",innerxml"`

	Muted *bool `xml:"muted,attr"`

	EnterSound *string `xml:"enterSound,attr"`

	ExitSound *string `xml:"exitSound,attr"`

	StartConferenceOnEnter *bool `xml:"startConferenceOnEnter,attr"`

	EndConferenceOnExit *bool `xml:"endConferenceOnExit,attr"`

	StayAlone *bool `xml:"stayAlone,attr"`

	WaitSound *string `xml:"waitSound,attr"`

	MaxMembers *int `xml:"maxMembers,attr"`

	Record *bool `xml:"record,attr"`

	RecordFileFormat *string `xml:"recordFileFormat,attr"`

	TimeLimit *int `xml:"timeLimit,attr"`

	HangupOnStar *bool `xml:"hangupOnStar,attr"`

	Action *string `xml:"action,attr"`

	Method *string `xml:"method,attr"`

	CallbackUrl *string `xml:"callbackUrl,attr"`

	CallbackMethod *string `xml:"callbackMethod,attr"`

	DigitsMatch *string `xml:"digitsMatch,attr"`

	FloorEvent *bool `xml:"floorEvent,attr"`

	Redirect *bool `xml:"redirect,attr"`

	RelayDTMF *bool `xml:"relayDTMF,attr"`

	XMLName xml.Name `xml:"Conference"`
}

func (e ConferenceElement) SetMuted(value bool) ConferenceElement {
	e.Muted = &value
	return e
}

func (e ConferenceElement) SetEnterSound(value string) ConferenceElement {
	e.EnterSound = &value
	return e
}

func (e ConferenceElement) SetExitSound(value string) ConferenceElement {
	e.ExitSound = &value
	return e
}

func (e ConferenceElement) SetStartConferenceOnEnter(value bool) ConferenceElement {
	e.StartConferenceOnEnter = &value
	return e
}

func (e ConferenceElement) SetEndConferenceOnExit(value bool) ConferenceElement {
	e.EndConferenceOnExit = &value
	return e
}

func (e ConferenceElement) SetStayAlone(value bool) ConferenceElement {
	e.StayAlone = &value
	return e
}

func (e ConferenceElement) SetWaitSound(value string) ConferenceElement {
	e.WaitSound = &value
	return e
}

func (e ConferenceElement) SetMaxMembers(value int) ConferenceElement {
	e.MaxMembers = &value
	return e
}

func (e ConferenceElement) SetRecord(value bool) ConferenceElement {
	e.Record = &value
	return e
}

func (e ConferenceElement) SetRecordFileFormat(value string) ConferenceElement {
	e.RecordFileFormat = &value
	return e
}

func (e ConferenceElement) SetTimeLimit(value int) ConferenceElement {
	e.TimeLimit = &value
	return e
}

func (e ConferenceElement) SetHangupOnStar(value bool) ConferenceElement {
	e.HangupOnStar = &value
	return e
}

func (e ConferenceElement) SetAction(value string) ConferenceElement {
	e.Action = &value
	return e
}

func (e ConferenceElement) SetMethod(value string) ConferenceElement {
	e.Method = &value
	return e
}

func (e ConferenceElement) SetCallbackUrl(value string) ConferenceElement {
	e.CallbackUrl = &value
	return e
}

func (e ConferenceElement) SetCallbackMethod(value string) ConferenceElement {
	e.CallbackMethod = &value
	return e
}

func (e ConferenceElement) SetDigitsMatch(value string) ConferenceElement {
	e.DigitsMatch = &value
	return e
}

func (e ConferenceElement) SetFloorEvent(value bool) ConferenceElement {
	e.FloorEvent = &value
	return e
}

func (e ConferenceElement) SetRedirect(value bool) ConferenceElement {
	e.Redirect = &value
	return e
}

func (e ConferenceElement) SetRelayDTMF(value bool) ConferenceElement {
	e.RelayDTMF = &value
	return e
}

func (e ConferenceElement) SetContents(value string) ConferenceElement {
	e.Contents = value
	return e
}

type DialElement struct {
	Contents []interface{} `xml:",innerxml"`

	Action *string `xml:"action,attr"`

	Method *string `xml:"method,attr"`

	HangupOnStar *bool `xml:"hangupOnStar,attr"`

	TimeLimit *int `xml:"timeLimit,attr"`

	Timeout *int `xml:"timeout,attr"`

	CallerID *string `xml:"callerId,attr"`

	CallerName *string `xml:"callerName,attr"`

	ConfirmSound *string `xml:"confirmSound,attr"`

	ConfirmKey *string `xml:"confirmKey,attr"`

	DialMusic *string `xml:"dialMusic,attr"`

	CallbackUrl *string `xml:"callbackUrl,attr"`

	CallbackMethod *string `xml:"callbackMethod,attr"`

	Redirect *bool `xml:"redirect,attr"`

	DigitsMatch *string `xml:"digitsMatch,attr"`

	DigitsMatchBLeg *string `xml:"digitsMatchBLeg,attr"`

	SipHeaders *string `xml:"sipHeaders,attr"`

	XMLName xml.Name `xml:"Dial"`
}

func (e DialElement) SetAction(value string) DialElement {
	e.Action = &value
	return e
}

func (e DialElement) SetMethod(value string) DialElement {
	e.Method = &value
	return e
}

func (e DialElement) SetHangupOnStar(value bool) DialElement {
	e.HangupOnStar = &value
	return e
}

func (e DialElement) SetTimeLimit(value int) DialElement {
	e.TimeLimit = &value
	return e
}

func (e DialElement) SetTimeout(value int) DialElement {
	e.Timeout = &value
	return e
}

func (e DialElement) SetCallerID(value string) DialElement {
	e.CallerID = &value
	return e
}

func (e DialElement) SetCallerName(value string) DialElement {
	e.CallerName = &value
	return e
}

func (e DialElement) SetConfirmSound(value string) DialElement {
	e.ConfirmSound = &value
	return e
}

func (e DialElement) SetConfirmKey(value string) DialElement {
	e.ConfirmKey = &value
	return e
}

func (e DialElement) SetDialMusic(value string) DialElement {
	e.DialMusic = &value
	return e
}

func (e DialElement) SetCallbackUrl(value string) DialElement {
	e.CallbackUrl = &value
	return e
}

func (e DialElement) SetCallbackMethod(value string) DialElement {
	e.CallbackMethod = &value
	return e
}

func (e DialElement) SetRedirect(value bool) DialElement {
	e.Redirect = &value
	return e
}

func (e DialElement) SetDigitsMatch(value string) DialElement {
	e.DigitsMatch = &value
	return e
}

func (e DialElement) SetDigitsMatchBLeg(value string) DialElement {
	e.DigitsMatchBLeg = &value
	return e
}

func (e DialElement) SetSipHeaders(value string) DialElement {
	e.SipHeaders = &value
	return e
}

func (e DialElement) SetContents(value []interface{}) DialElement {
	e.Contents = value
	return e
}

type NumberElement struct {
	Contents string `xml:",innerxml"`

	SendDigits *string `xml:"sendDigits,attr"`

	SendOnPreanswer *bool `xml:"sendOnPreanswer,attr"`

	XMLName xml.Name `xml:"Number"`
}

func (e NumberElement) SetSendDigits(value string) NumberElement {
	e.SendDigits = &value
	return e
}

func (e NumberElement) SetSendOnPreanswer(value bool) NumberElement {
	e.SendOnPreanswer = &value
	return e
}

func (e NumberElement) SetContents(value string) NumberElement {
	e.Contents = value
	return e
}

type UserElement struct {
	Contents string `xml:",innerxml"`

	SendDigits *string `xml:"sendDigits,attr"`

	SendOnPreanswer *bool `xml:"sendOnPreanswer,attr"`

	SipHeaders *string `xml:"sipHeaders,attr"`

	XMLName xml.Name `xml:"User"`
}

func (e UserElement) SetSendDigits(value string) UserElement {
	e.SendDigits = &value
	return e
}

func (e UserElement) SetSendOnPreanswer(value bool) UserElement {
	e.SendOnPreanswer = &value
	return e
}

func (e UserElement) SetSipHeaders(value string) UserElement {
	e.SipHeaders = &value
	return e
}

func (e UserElement) SetContents(value string) UserElement {
	e.Contents = value
	return e
}

type DTMFElement struct {
	Contents string `xml:",innerxml"`

	Async *bool `xml:"async,attr"`

	XMLName xml.Name `xml:"DTMF"`
}

func (e DTMFElement) SetAsync(value bool) DTMFElement {
	e.Async = &value
	return e
}

func (e DTMFElement) SetContents(value string) DTMFElement {
	e.Contents = value
	return e
}

type GetDigitsElement struct {
	Contents []interface{} `xml:",innerxml"`

	Action *string `xml:"action,attr"`

	Method *string `xml:"method,attr"`

	Timeout *int `xml:"timeout,attr"`

	DigitTimeout *int `xml:"digitTimeout,attr"`

	FinishOnKey *string `xml:"finishOnKey,attr"`

	NumDigits *int `xml:"numDigits,attr"`

	Retries *int `xml:"retries,attr"`

	Redirect *bool `xml:"redirect,attr"`

	PlayBeep *bool `xml:"playBeep,attr"`

	ValidDigits *string `xml:"validDigits,attr"`

	InvalidDigitsSound *string `xml:"invalidDigitsSound,attr"`

	Log *bool `xml:"log,attr"`

	XMLName xml.Name `xml:"GetDigits"`
}

func (e GetDigitsElement) SetAction(value string) GetDigitsElement {
	e.Action = &value
	return e
}

func (e GetDigitsElement) SetMethod(value string) GetDigitsElement {
	e.Method = &value
	return e
}

func (e GetDigitsElement) SetTimeout(value int) GetDigitsElement {
	e.Timeout = &value
	return e
}

func (e GetDigitsElement) SetDigitTimeout(value int) GetDigitsElement {
	e.DigitTimeout = &value
	return e
}

func (e GetDigitsElement) SetFinishOnKey(value string) GetDigitsElement {
	e.FinishOnKey = &value
	return e
}

func (e GetDigitsElement) SetNumDigits(value int) GetDigitsElement {
	e.NumDigits = &value
	return e
}

func (e GetDigitsElement) SetRetries(value int) GetDigitsElement {
	e.Retries = &value
	return e
}

func (e GetDigitsElement) SetRedirect(value bool) GetDigitsElement {
	e.Redirect = &value
	return e
}

func (e GetDigitsElement) SetPlayBeep(value bool) GetDigitsElement {
	e.PlayBeep = &value
	return e
}

func (e GetDigitsElement) SetValidDigits(value string) GetDigitsElement {
	e.ValidDigits = &value
	return e
}

func (e GetDigitsElement) SetInvalidDigitsSound(value string) GetDigitsElement {
	e.InvalidDigitsSound = &value
	return e
}

func (e GetDigitsElement) SetLog(value bool) GetDigitsElement {
	e.Log = &value
	return e
}

func (e GetDigitsElement) SetContents(value []interface{}) GetDigitsElement {
	e.Contents = value
	return e
}

type GetInputElement struct {
	Contents []interface{} `xml:",innerxml"`

	Action *string `xml:"action,attr"`

	Method *string `xml:"method,attr"`

	InputType *string `xml:"inputType,attr"`

	ExecutionTimeout *int `xml:"executionTimeout,attr"`

	DigitEndTimeout *int `xml:"digitEndTimeout,attr"`

	SpeechEndTimeout *int `xml:"speechEndTimeout,attr"`

	FinishOnKey *string `xml:"finishOnKey,attr"`

	NumDigits *int `xml:"numDigits,attr"`

	SpeechModel *string `xml:"speechModel,attr"`

	Hints *string `xml:"hints,attr"`

	Language *string `xml:"language,attr"`

	InterimSpeechResultsCallback *string `xml:"interimSpeechResultsCallback,attr"`

	InterimSpeechResultsCallbackMethod *string `xml:"interimSpeechResultsCallbackMethod,attr"`

	Redirect *bool `xml:"redirect,attr"`

	Log *bool `xml:"log,attr"`

	ProfanityFilter *bool `xml:"profanityFilter,attr"`

	XMLName xml.Name `xml:"GetInput"`
}

func (e GetInputElement) SetAction(value string) GetInputElement {
	e.Action = &value
	return e
}

func (e GetInputElement) SetMethod(value string) GetInputElement {
	e.Method = &value
	return e
}

func (e GetInputElement) SetInputType(value string) GetInputElement {
	e.InputType = &value
	return e
}

func (e GetInputElement) SetExecutionTimeout(value int) GetInputElement {
	e.ExecutionTimeout = &value
	return e
}

func (e GetInputElement) SetDigitEndTimeout(value int) GetInputElement {
	e.DigitEndTimeout = &value
	return e
}

func (e GetInputElement) SetSpeechEndTimeout(value int) GetInputElement {
	e.SpeechEndTimeout = &value
	return e
}

func (e GetInputElement) SetFinishOnKey(value string) GetInputElement {
	e.FinishOnKey = &value
	return e
}

func (e GetInputElement) SetNumDigits(value int) GetInputElement {
	e.NumDigits = &value
	return e
}

func (e GetInputElement) SetSpeechModel(value string) GetInputElement {
	e.SpeechModel = &value
	return e
}

func (e GetInputElement) SetHints(value string) GetInputElement {
	e.Hints = &value
	return e
}

func (e GetInputElement) SetLanguage(value string) GetInputElement {
	e.Language = &value
	return e
}

func (e GetInputElement) SetInterimSpeechResultsCallback(value string) GetInputElement {
	e.InterimSpeechResultsCallback = &value
	return e
}

func (e GetInputElement) SetInterimSpeechResultsCallbackMethod(value string) GetInputElement {
	e.InterimSpeechResultsCallbackMethod = &value
	return e
}

func (e GetInputElement) SetRedirect(value bool) GetInputElement {
	e.Redirect = &value
	return e
}

func (e GetInputElement) SetLog(value bool) GetInputElement {
	e.Log = &value
	return e
}

func (e GetInputElement) SetProfanityFilter(value bool) GetInputElement {
	e.ProfanityFilter = &value
	return e
}

func (e GetInputElement) SetContents(value []interface{}) GetInputElement {
	e.Contents = value
	return e
}

type HangupElement struct {
	Reason *string `xml:"reason,attr"`

	Schedule *int `xml:"schedule,attr"`

	XMLName xml.Name `xml:"Hangup"`
}

func (e HangupElement) SetReason(value string) HangupElement {
	e.Reason = &value
	return e
}

func (e HangupElement) SetSchedule(value int) HangupElement {
	e.Schedule = &value
	return e
}

type MessageElement struct {
	Contents string `xml:",innerxml"`

	Src *string `xml:"src,attr"`

	Dst *string `xml:"dst,attr"`

	Type *string `xml:"type,attr"`

	CallbackUrl *string `xml:"callbackUrl,attr"`

	CallbackMethod *string `xml:"callbackMethod,attr"`

	XMLName xml.Name `xml:"Message"`
}

func (e MessageElement) SetSrc(value string) MessageElement {
	e.Src = &value
	return e
}

func (e MessageElement) SetDst(value string) MessageElement {
	e.Dst = &value
	return e
}

func (e MessageElement) SetType(value string) MessageElement {
	e.Type = &value
	return e
}

func (e MessageElement) SetCallbackUrl(value string) MessageElement {
	e.CallbackUrl = &value
	return e
}

func (e MessageElement) SetCallbackMethod(value string) MessageElement {
	e.CallbackMethod = &value
	return e
}

func (e MessageElement) SetContents(value string) MessageElement {
	e.Contents = value
	return e
}

type PlayElement struct {
	Contents string `xml:",innerxml"`

	Loop *int `xml:"loop,attr"`

	XMLName xml.Name `xml:"Play"`
}

func (e PlayElement) SetLoop(value int) PlayElement {
	e.Loop = &value
	return e
}

func (e PlayElement) SetContents(value string) PlayElement {
	e.Contents = value
	return e
}

type PreAnswerElement struct {
	Contents []interface{} `xml:",innerxml"`

	XMLName xml.Name `xml:"PreAnswer"`
}

func (e PreAnswerElement) SetContents(value []interface{}) PreAnswerElement {
	e.Contents = value
	return e
}

type RecordElement struct {
	Action *string `xml:"action,attr"`

	Method *string `xml:"method,attr"`

	FileFormat *string `xml:"fileFormat,attr"`

	Redirect *bool `xml:"redirect,attr"`

	Timeout *int `xml:"timeout,attr"`

	MaxLength *int `xml:"maxLength,attr"`

	PlayBeep *bool `xml:"playBeep,attr"`

	FinishOnKey *string `xml:"finishOnKey,attr"`

	RecordSession *bool `xml:"recordSession,attr"`

	StartOnDialAnswer *bool `xml:"startOnDialAnswer,attr"`

	TranscriptionType *string `xml:"transcriptionType,attr"`

	TranscriptionUrl *string `xml:"transcriptionUrl,attr"`

	TranscriptionMethod *string `xml:"transcriptionMethod,attr"`

	CallbackUrl *string `xml:"callbackUrl,attr"`

	CallbackMethod *string `xml:"callbackMethod,attr"`

	XMLName xml.Name `xml:"Record"`
}

func (e RecordElement) SetAction(value string) RecordElement {
	e.Action = &value
	return e
}

func (e RecordElement) SetMethod(value string) RecordElement {
	e.Method = &value
	return e
}

func (e RecordElement) SetFileFormat(value string) RecordElement {
	e.FileFormat = &value
	return e
}

func (e RecordElement) SetRedirect(value bool) RecordElement {
	e.Redirect = &value
	return e
}

func (e RecordElement) SetTimeout(value int) RecordElement {
	e.Timeout = &value
	return e
}

func (e RecordElement) SetMaxLength(value int) RecordElement {
	e.MaxLength = &value
	return e
}

func (e RecordElement) SetPlayBeep(value bool) RecordElement {
	e.PlayBeep = &value
	return e
}

func (e RecordElement) SetFinishOnKey(value string) RecordElement {
	e.FinishOnKey = &value
	return e
}

func (e RecordElement) SetRecordSession(value bool) RecordElement {
	e.RecordSession = &value
	return e
}

func (e RecordElement) SetStartOnDialAnswer(value bool) RecordElement {
	e.StartOnDialAnswer = &value
	return e
}

func (e RecordElement) SetTranscriptionType(value string) RecordElement {
	e.TranscriptionType = &value
	return e
}

func (e RecordElement) SetTranscriptionUrl(value string) RecordElement {
	e.TranscriptionUrl = &value
	return e
}

func (e RecordElement) SetTranscriptionMethod(value string) RecordElement {
	e.TranscriptionMethod = &value
	return e
}

func (e RecordElement) SetCallbackUrl(value string) RecordElement {
	e.CallbackUrl = &value
	return e
}

func (e RecordElement) SetCallbackMethod(value string) RecordElement {
	e.CallbackMethod = &value
	return e
}

type RedirectElement struct {
	Contents string `xml:",innerxml"`

	Method *string `xml:"method,attr"`

	XMLName xml.Name `xml:"Redirect"`
}

func (e RedirectElement) SetMethod(value string) RedirectElement {
	e.Method = &value
	return e
}

func (e RedirectElement) SetContents(value string) RedirectElement {
	e.Contents = value
	return e
}

type SpeakElement struct {
	Contents string `xml:",innerxml"`

	Voice *string `xml:"voice,attr"`

	Language *string `xml:"language,attr"`

	Loop *int `xml:"loop,attr"`

	XMLName xml.Name `xml:"Speak"`
}

func (e SpeakElement) SetVoice(value string) SpeakElement {
	e.Voice = &value
	return e
}

func (e SpeakElement) SetLanguage(value string) SpeakElement {
	e.Language = &value
	return e
}

func (e SpeakElement) SetLoop(value int) SpeakElement {
	e.Loop = &value
	return e
}

func (e SpeakElement) SetContents(value string) SpeakElement {
	e.Contents = value
	return e
}

type WaitElement struct {
	Length *int `xml:"length,attr"`

	Silence *bool `xml:"silence,attr"`

	MinSilence *int `xml:"minSilence,attr"`

	Beep *bool `xml:"beep,attr"`

	XMLName xml.Name `xml:"Wait"`
}

func (e WaitElement) SetLength(value int) WaitElement {
	e.Length = &value
	return e
}

func (e WaitElement) SetSilence(value bool) WaitElement {
	e.Silence = &value
	return e
}

func (e WaitElement) SetMinSilence(value int) WaitElement {
	e.MinSilence = &value
	return e
}

func (e WaitElement) SetBeep(value bool) WaitElement {
	e.Beep = &value
	return e
}
