package plivo

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type MediaService struct {
	client *Client
	Media
}
type Media struct {
	ContentType string `json:"content_type,omitempty" url:"content_type,omitempty"`
	FileName    string `json:"file_name,omitempty" url:"file_name,omitempty"`
	MediaID     string `json:"media_id,omitempty" url:"media_id,omitempty"`
	Size        string `json:"size,omitempty" url:"size,omitempty"`
	ResourceURI string `json:"upload_time,omitempty" url:"upload_time,omitempty"`
	CarrierRate string `json:"url,omitempty" url:"url,omitempty"`
}

type MediaMeta struct {
	Previous   *string
	Next       *string
	TotalCount int `json:"total_count" url:"api_id"`
	Offset     int
	Limit      int
}
type MediaResponseBody struct {
	Media []string `json:"objects" url:"objcts"`
	ApiID string   `json:"api_id" url:"api_id"`
}

type BaseListMediaResponse struct {
	ApiID string    `json:"api_id" url:"api_id"`
	Meta  MediaMeta `json:"meta" url:"meta"`
}

type MediaUpload struct {
	File interface{} `json:"file,omitempty" url:"file,omitempty"`
}

type MediaListParams struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

func IsValidFile(fileName string) error {
	fileParts := strings.Split(fileName, ".")
	if len(fileParts) < 2 {
		err := errors.New("Invalid file specified")
		return err
	}
	// extension := strings.ToLower(fileParts[len(fileParts)-1])

	targetFilePath := filepath.Join(fileName, "./")

	if file, err := os.Stat(targetFilePath); err == nil {
		sizeInMB := file.Size() / (1024 * 1024)
		if sizeInMB > 5 {
			err := errors.New("File size exeeds 5 MB limit")
			return err
		}
	} else if os.IsNotExist(err) {
		err := errors.New("File not found in the specified path")
		return err
	}
	return nil
}

func (service *MediaService) Upload(params MediaUpload) (response *MediaResponseBody, err error) {
	if params.File != nil {
		err = IsValidFile(params.File.(string))
		if err != nil {
			return
		}
		targetFilePath := filepath.Join(params.File.(string), "./")
		file, _ := os.Stat(targetFilePath)
		params.File = file
	}
	request, err := service.client.NewRequest("POST", params, "Media")
	if err != nil {
		return
	}
	response = &MediaResponseBody{}
	err = service.client.ExecuteRequest(request, response)
	return
}

func (service *MediaService) Get(media_id string) (response *Media, err error) {
	req, err := service.client.NewRequest("GET", nil, "Media/%s/", media_id)
	if err != nil {
		return
	}
	resp := &Media{}
	err = service.client.ExecuteRequest(req, resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	return resp, nil
}

func (service *MediaService) List(param *MediaListParams) (response *BaseListMediaResponse, err error) {
	req, err := service.client.NewRequest("GET", param, "Media/")
	if err != nil {
		return
	}
	resp := &BaseListMediaResponse{}
	err = service.client.ExecuteRequest(req, resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	return resp, nil
}
