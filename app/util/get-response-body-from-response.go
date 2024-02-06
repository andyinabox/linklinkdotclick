package util

import (
	"io/ioutil"
	"net/http"
)

func GetResponseBodyFromResponse(res *http.Response) ([]byte, error) {
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
