package plivo

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
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
	Size        int    `json:"size,omitempty" url:"size,omitempty"`
	UploadTime  string `json:"upload_time,omitempty" url:"upload_time,omitempty"`
	URL         string `json:"url,omitempty" url:"url,omitempty"`
}

type MediaUploadResponse struct {
	ContentType  string `json:"content_type,omitempty" url:"content_type,omitempty"`
	FileName     string `json:"file_name,omitempty" url:"file_name,omitempty"`
	MediaID      string `json:"media_id,omitempty" url:"media_id,omitempty"`
	Size         int    `json:"size,omitempty" url:"size,omitempty"`
	UploadTime   string `json:"upload_time,omitempty" url:"upload_time,omitempty"`
	URL          string `json:"url,omitempty" url:"url,omitempty"`
	Status       string `json:"status,omitempty" url:"status,omitempty"`
	StatusCode   int    `json:"status_code,omitempty" url:"status_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty" url:"error_message,omitempty"`
	ErrorCode    int    `json:"error_code,omitempty" url:"error_code,omitempty"`
}

type MediaMeta struct {
	Previous   *string
	Next       *string
	TotalCount int `json:"total_count" url:"api_id"`
	Offset     int `json:"offset,omitempty" url:"offset,omitempty"`
	Limit      int `json:"limit,omitempty" url:"limit,omitempty"`
}
type MediaResponseBody struct {
	Media []MediaUploadResponse `json:"objects" url:"objects"`
	ApiID string                `json:"api_id" url:"api_id"`
}

type BaseListMediaResponse struct {
	ApiID string    `json:"api_id" url:"api_id"`
	Meta  MediaMeta `json:"meta" url:"meta"`
	Media []Media   `json:"objects" url:"objects"`
}

type MediaUpload struct {
	UploadFiles []Files
}
type Files struct {
	FilePath    string
	ContentType string
}

type MediaListParams struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}
func (service *MediaService) Upload(params MediaUpload) (response *MediaResponseBody, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for i := 0; i < len(params.UploadFiles); i++ {
		file, errFile1 := os.Open(params.UploadFiles[i].FilePath)
		defer file.Close()
		filename := filepath.Base(params.UploadFiles[i].FilePath)
		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition",
			fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
				escapeQuotes("file"), escapeQuotes(filename)))
		h.Set("Content-Type", params.UploadFiles[i].ContentType)
		part1, errFile1 := writer.CreatePart(h)
		_, errFile1 = io.Copy(part1, file)
		if errFile1 != nil {
			return nil, errFile1
		}
		filerror := writer.Close()
		if filerror != nil {
			return nil, filerror
		}
	}
	requestUrl := service.client.BaseUrl
	requestUrl.Path = fmt.Sprintf(baseRequestString, fmt.Sprintf(service.client.AuthId+"/Media"))
	request, err := http.NewRequest("POST", requestUrl.String(), payload)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.SetBasicAuth(service.client.AuthId, service.client.AuthToken)
	response = &MediaResponseBody{}
	err = service.client.ExecuteRequest(request, response)
	return
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}

func (service *MediaService) Get(media_id string) (response *Media, err error) {
	req, err := service.client.NewRequest("GET", nil, "Media/%s", media_id)
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

func (service *MediaService) List(param MediaListParams) (response *BaseListMediaResponse, err error) {
	req, err := service.client.NewRequest("GET", param, "Media")
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
