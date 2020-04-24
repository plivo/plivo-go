package plivo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedia_List(t *testing.T) {
	expectResponse("MediaListResponse.json", 200)

	if _, err := client.Media.List(MediaListParams{Limit: 0, Offset: 20}); err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err := client.Media.List(MediaListParams{Limit: 0, Offset: 20})
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Media")
}

func TestMedia_Get(t *testing.T) {
	expectResponse("MediaGetResponse.json", 200)
	MediaID := "MediaID"

	media, err := client.Media.Get(MediaID)
	assert.Equal(t, MediaID, media.MediaID)
	if err != nil {
		panic(err)
	}

	cl := client.httpClient
	client.httpClient = nil
	_, err = client.Media.Get(MediaID)
	if err == nil {
		client.httpClient = cl
		panic(errors.New("error expected"))
	}
	client.httpClient = cl

	assertRequest(t, "GET", "Media/%s", MediaID)
}
