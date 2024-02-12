package responses

import (
	"errors"
	"io/ioutil"
	"net/http"
)

var ErrServerError = errors.New("server error")

func GetBody(res *http.Response) ([]byte, error) {
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func GetBodyFromUrl(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, ErrServerError
	}
	return GetBody(res)
}
