package plivo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMessageService_List(t *testing.T) {
	expectResponse("messageListResponse.json", 200)
	assert := require.New(t)
	resp, err := client.Messages.List(MessageListParams{})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotNil(resp.Objects)
	assert.NotEmpty(resp.Objects[0].MessageUUID)

	assert.Equal(resp.Objects[0].RequesterIP, "192.168.1.1")
	assert.Equal(resp.Objects[19].RequesterIP, "192.168.1.20")

	assert.Equal(resp.Objects[0].DLTEntityID, "1111")
	assert.Equal(resp.Objects[0].DLTTemplateID, "2222")
	assert.Equal(resp.Objects[0].DLTTemplateCategory, "promotional")

	assert.Equal(resp.Objects[19].DLTEntityID, "")
	assert.Equal(resp.Objects[19].DLTTemplateID, "")
	assert.Equal(resp.Objects[19].DLTTemplateCategory, "")

	assert.Equal(resp.Objects[0].ConversationID, "0079")
	assert.Equal(resp.Objects[0].ConversationOrigin, "marketing")
	assert.Equal(resp.Objects[0].ConversationExpirationTimestamp, "2023-08-03 23:02:00+05:30")

	assert.Equal(resp.Objects[19].ConversationID, "")
	assert.Equal(resp.Objects[19].ConversationOrigin, "")
	assert.Equal(resp.Objects[19].ConversationExpirationTimestamp, "")

	assert.NotNil(resp.Meta)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Messages.List(MessageListParams{})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl
	assertRequest(t, "GET", "Message")
}

func TestMessageService_Get(t *testing.T) {
	expectResponse("messageGetResponse.json", 200)
	uuid := "5b40a428-bfc7-4daf-9d06-726c558bf3b8"
	assert := require.New(t)
	resp, err := client.Messages.Get(uuid)
	assert.NotNil(resp)
	assert.Nil(err)
	assert.Equal(resp.MessageUUID, uuid)
	assert.Equal(resp.RequesterIP, "192.168.1.1")

	assert.Equal(resp.DLTEntityID, "1234")
	assert.Equal(resp.DLTTemplateID, "5678")
	assert.Equal(resp.DLTTemplateCategory, "service_implicit")

	assert.Equal(resp.ConversationID, "9876")
	assert.Equal(resp.ConversationOrigin, "utility")
	assert.Equal(resp.ConversationExpirationTimestamp, "2023-08-03 23:02:00+05:30")

	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Messages.Get(uuid)
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "GET", "Message/%s", uuid)
}

func TestListMedia(t *testing.T) {
	expectResponse("mmsMediaListResponse.json", 200)
	uuid := "f734eeec-e59f-11e9-89dc-0242ac110003"
	assert := require.New(t)
	resp, err := client.Messages.ListMedia(uuid)
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.Objects[0].MediaID)
	assert.Equal(resp.Objects[0].MessageUUID, uuid)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Messages.ListMedia(uuid)
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "GET", "Message/%s/Media/", uuid)
}

func TestMessageService_Create(t *testing.T) {
	expectResponse("messageSendResponse.json", 202)
	assert := require.New(t)
	resp, err := client.Messages.Create(MessageCreateParams{
		Src:  "+911231231230",
		Dst:  "+911231231231",
		Text: "Testing",
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	assert.NotEmpty(resp.MessageUUID)
	assert.Equal(resp.Error, "")
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Messages.Create(MessageCreateParams{
		Src:  "+911231231230",
		Dst:  "+911231231231",
		Text: "Testing",
	})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "Message")
}

func TestMessageService_Get_INRFields(t *testing.T) {
	expectResponse("messageGetInrResponse.json", 200)
	assert := require.New(t)
	resp, err := client.Messages.Get("5b40a428-bfc7-4daf-9d06-726c558bf3b8")
	assert.Nil(err)
	assert.Equal("0.00500", resp.TotalAmount)
	assert.Equal("0.41500", resp.TotalAmountINR)
	assert.Equal("0.00500", resp.TotalRate)
	assert.Equal("0.41500", resp.TotalRateINR)
	assert.Equal("0.00200", resp.CarrierFees)
	assert.Equal("0.16600", resp.CarrierFeesINR)
	assert.Equal("0.00100", resp.CarrierFeesRate)
	assert.Equal("0.08300", resp.CarrierFeesRateINR)
}

func TestMessageService_List_INRFields(t *testing.T) {
	expectResponse("messageListInrResponse.json", 200)
	assert := require.New(t)
	resp, err := client.Messages.List(MessageListParams{})
	assert.Nil(err)
	assert.Len(resp.Objects, 1)
	assert.Equal("0.41500", resp.Objects[0].TotalAmountINR)
	assert.Equal("0.00500", resp.Objects[0].TotalRate)
	assert.Equal("0.41500", resp.Objects[0].TotalRateINR)
	assert.Equal("0.16600", resp.Objects[0].CarrierFeesINR)
	assert.Equal("0.08300", resp.Objects[0].CarrierFeesRateINR)
}

func TestMessageService_Get_USDAccountHasNoINRFields(t *testing.T) {
	expectResponse("messageGetResponse.json", 200)
	assert := require.New(t)
	resp, err := client.Messages.Get("5b40a428-bfc7-4daf-9d06-726c558bf3b8")
	assert.Nil(err)
	assert.Equal("0.00250", resp.TotalAmount)
	assert.Equal("0.00250", resp.TotalRate)
	assert.Empty(resp.TotalAmountINR)
	assert.Empty(resp.TotalRateINR)
	assert.Empty(resp.CarrierFeesINR)
	assert.Empty(resp.CarrierFeesRateINR)
}
