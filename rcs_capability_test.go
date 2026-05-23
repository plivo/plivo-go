package plivo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRcsCapabilityService_Check(t *testing.T) {
	expectResponse("rcsCapabilityCheckResponse.json", 200)
	assert := require.New(t)
	resp, err := client.RcsCapability.Check(RcsCapabilityParams{
		PhoneNumber: "+14151234567",
	})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.ApiID)
	assert.NotEmpty(resp.PhoneNumber)

	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.RcsCapability.Check(RcsCapabilityParams{
		PhoneNumber: "+14151234567",
	})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "GET", "RCS/Capability/")
}