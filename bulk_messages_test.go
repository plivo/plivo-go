package plivo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBulkMessageService_Create(t *testing.T) {
	expectResponse("bulkMessageCreateResponse.json", 202)
	assert := require.New(t)
	resp, err := client.BulkMessages.Create(BulkMessageCreateParams{
		Src:  "+911231231230",
		Dst:  []string{"+911231231231", "+911231231232"},
		Text: "Testing bulk",
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	assert.NotEmpty(resp.MessageUUID)
	assert.Equal(resp.Message, "message(s) queued")
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.BulkMessages.Create(BulkMessageCreateParams{
		Src:  "+911231231230",
		Dst:  []string{"+911231231231", "+911231231232"},
		Text: "Testing bulk",
	})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "Message/Bulk")
}