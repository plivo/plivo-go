package plivo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMedia_List(t *testing.T) {
	expectResponse("MediaListResponse.json", 200)
	assert := require.New(t)
	resp, err := client.Media.List(MediaListParams{Limit: 0, Offset: 20})
	assert.NotNil(resp)
	assert.Nil(err)
	assert.NotEmpty(resp.Media[0].MediaID)
	assert.NotNil(resp.Meta)
	cl := client.httpClient
	client.httpClient = nil
	resp, err = client.Media.List(MediaListParams{Limit: 0, Offset: 20})
	assert.NotNil(err)
	assert.Nil(resp)
	client.httpClient = cl

	assertRequest(t, "GET", "Media")
}

func TestMedia_Get(t *testing.T) {
	expectResponse("MediaGetResponse.json", 200)
	MediaID := "98854bc5-ea05-4837-a301-0272523e6156"
	assert := require.New(t)
	media, err := client.Media.Get(MediaID)
	assert.NotNil(media)
	assert.Nil(err)
	assert.Equal(MediaID, media.MediaID)
	assert.NotEmpty(media.ContentType)

	cl := client.httpClient
	client.httpClient = nil
	media, err = client.Media.Get(MediaID)
	assert.NotNil(err)
	assert.Nil(media)
	client.httpClient = cl

	assertRequest(t, "GET", "Media/%s", MediaID)
}
