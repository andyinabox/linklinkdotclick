package util

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
)

func GetResponseBodyFromUrl(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, app.ErrServerError
	}
	return GetResponseBodyFromResponse(res)
}
