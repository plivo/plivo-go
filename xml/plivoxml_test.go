package xml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleXML(t *testing.T) {
	assert.Equal(t, "<Response><Message>Plivo®</Message></Response>", ResponseElement{
		Contents: []interface{}{
			new(MessageElement).SetContents("Plivo®"),
		},
	}.String())
}

func TestComplexXML(t *testing.T) {
	assert.Equal(t, "<Response><Message src=\"123\" dst=\"321\" type=\"sms\" callbackUrl=\"url\" callbackMethod=\"POST\">Plivo®</Message><Record action=\"action\" method=\"POST\" fileFormat=\"fileformat\" redirect=\"true\" timeout=\"10\" maxLength=\"100\" playBeep=\"true\" finishOnKey=\"key\" recordSession=\"true\" callbackUrl=\"url\" callbackMethod=\"POST\"></Record><Dial action=\"action\" method=\"POST\" hangupOnStar=\"true\" timeLimit=\"10\" timeout=\"10\" callbackUrl=\"url\" callbackMethod=\"POST\" redirect=\"true\" digitsMatch=\"123\" sipHeaders=\"h=1\"><Number sendDigits=\"123\" sendOnPreanswer=\"true\">123</Number><User sendDigits=\"123\" sendOnPreanswer=\"true\" sipHeaders=\"header=1\">user</User></Dial><DTMF async=\"true\">123</DTMF><Wait length=\"10\" silence=\"true\" minSilence=\"10\" beep=\"true\"></Wait><PreAnswer><Speak voice=\"MAN\" language=\"en\" loop=\"10\">text</Speak></PreAnswer><GetDigits action=\"action\" timeout=\"10\" digitTimeout=\"10\" finishOnKey=\"#\" numDigits=\"10\" retries=\"10\" redirect=\"true\" playBeep=\"true\" validDigits=\"123\" invalidDigitsSound=\"sound\"></GetDigits><Conference muted=\"true\" enterSound=\"sound\" exitSound=\"sound\" startConferenceOnEnter=\"true\" endConferenceOnExit=\"true\" stayAlone=\"true\" waitSound=\"sound\" maxMembers=\"10\" record=\"true\" recordFileFormat=\"rff\" timeLimit=\"10\" hangupOnStar=\"true\" action=\"action\" method=\"method\" callbackUrl=\"url\" callbackMethod=\"method\" digitsMatch=\"123\" floorEvent=\"true\" redirect=\"true\" relayDTMF=\"true\">name</Conference><Redirect method=\"POST\">url</Redirect><Play loop=\"10\">url</Play><Hangup reason=\"test\" schedule=\"10\"></Hangup></Response>", ResponseElement{
		Contents: []interface{}{
			new(MessageElement).SetSrc("123").SetDst("321").SetType("sms").SetCallbackMethod("POST").SetCallbackUrl("url").SetContents("Plivo®"),
			new(RecordElement).SetCallbackMethod("POST").SetCallbackUrl("url").SetAction("action").SetFileFormat("fileformat").SetFinishOnKey("key").SetRecordSession(true).SetRedirect(true).SetPlayBeep(true).SetTimeout(10).SetMaxLength(100).SetMethod("POST"),
			new(DialElement).SetContents([]interface{}{
				new(NumberElement).SetContents("123").SetSendOnPreanswer(true).SetSendDigits("123"),
				new(UserElement).SetContents("user").SetSendOnPreanswer(true).SetSendDigits("123").SetSipHeaders("header=1"),
			}).SetTimeLimit(10).SetMethod("POST").SetHangupOnStar(true).SetDigitsMatch("123").SetCallbackMethod("POST").SetCallbackUrl("url").SetRedirect(true).SetAction("action").SetTimeout(10).SetSipHeaders("h=1"),
			new(DTMFElement).SetContents("123").SetAsync(true),
			new(WaitElement).SetBeep(true).SetLength(10).SetMinSilence(10).SetSilence(true),
			new(PreAnswerElement).SetContents([]interface{}{
				new(SpeakElement).SetContents("text").SetLanguage("en").SetLoop(10).SetVoice("MAN"),
			}),
			new(GetDigitsElement).SetPlayBeep(true).SetRedirect(true).SetFinishOnKey("#").SetAction("action").SetDigitTimeout(10).SetNumDigits(10).SetRetries(10).SetTimeout(10).SetValidDigits("123").SetInvalidDigitsSound("sound").SetContents([]interface{}{}),
			new(ConferenceElement).SetAction("action").SetRedirect(true).SetCallbackUrl("url").SetCallbackMethod("method").SetContents("name").SetDigitsMatch("123").SetEndConferenceOnExit(true).SetEnterSound("sound").SetExitSound("sound").SetFloorEvent(true).SetHangupOnStar(true).SetMaxMembers(10).SetMethod("method").SetMuted(true).SetRecord(true).SetRecordFileFormat("rff").SetRelayDTMF(true).SetStartConferenceOnEnter(true).SetStayAlone(true).SetTimeLimit(10).SetWaitSound("sound"),
			new(RedirectElement).SetMethod("POST").SetContents("url"),
			new(PlayElement).SetContents("url").SetLoop(10),
			new(HangupElement).SetReason("test").SetSchedule(10),
		},
	}.String())
}
