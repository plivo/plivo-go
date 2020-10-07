package plivo

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLookupGet(t *testing.T) {
	expectResponse("LookupGetResponse.json", http.StatusOK)
	assert := require.New(t)

	number := "+14154305555"
	infoType := "carrier"
	resp, err := client.Lookup.Get(number, LookupParams{
		Type: infoType,
	})
	assert.Nil(err)
	assert.NotNil(resp)
	assert.NotNil(resp.Country)
	assert.NotNil(resp.Carrier)
	assert.NotNil(resp.Format)
	assert.Equal(number, resp.PhoneNumber)
	assert.Equal(number, resp.Format.E164)
	assert.Equal(fmt.Sprintf("/v1/Lookup/Number/%s?type=%s", number, infoType), resp.ResourceURI)

	assertBaseRequest(t, http.MethodGet, "/v1/Lookup/Number/%s?type=%s", number, infoType)
}

func TestLookupGetError(t *testing.T) {
	expectResponse("LookupGetError.json", http.StatusNotFound)
	assert := require.New(t)

	number := "xxxxxxxxxxx"
	infoType := "carrier"
	resp, err := client.Lookup.Get(number, LookupParams{
		Type: infoType,
	})
	assert.NotNil(err)
	assert.Nil(resp)

	errResp, ok := err.(*LookupError)
	assert.True(ok)
	assert.Equal(404, errResp.ErrorCode)
	assert.Equal("Phone number is invalid or does not exist.", errResp.Message)

	assertBaseRequest(t, http.MethodGet, "/v1/Lookup/Number/%s?type=%s", number, infoType)
}

func TestLookupGetValidation(t *testing.T) {
	assert := require.New(t)

	resp, err := client.Lookup.Get("+14154305555", LookupParams{
		Type: "", // empty type string
	})
	assert.NotNil(err)
	assert.Equal("Type must be set in params", err.Error())
	assert.Nil(resp)
}
