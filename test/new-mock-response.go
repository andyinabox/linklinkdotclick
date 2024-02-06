package test

import (
	"net/http"
	"net/url"
	"os"
	"testing"
)

func NewMockResponse(pageUrl string, filePath string, t *testing.T) *http.Response {
	url, err := url.Parse(pageUrl)
	if err != nil {
		t.Fatal(err.Error())
	}

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err.Error())
	}

	return &http.Response{
		Body: file,
		Request: &http.Request{
			URL: url,
		},
	}
}
