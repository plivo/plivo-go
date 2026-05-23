package plivo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRcsAssistantEventService_Create(t *testing.T) {
	expectResponse("rcsAssistantEventCreateResponse.json", 200)
	assert := require.New(t)
	resp, err := client.RcsAssistantEvents.Create(RcsAssistantEventCreateParams{})
	assert.NotNil(resp)
	assert.Nil(err)

	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.RcsAssistantEvents.Create(RcsAssistantEventCreateParams{})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "POST", "RCS/AssistantEvents")
}