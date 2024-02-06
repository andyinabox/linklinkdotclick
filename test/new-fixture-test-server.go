package test

import (
	"io/ioutil"
	"mime"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
)

func NewFixtureTestServer(filePath string, t *testing.T) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open(filePath)
		if err != nil {
			t.Fatal(err.Error())
		}

		fileInfo, err := file.Stat()
		if err != nil {
			t.Fatal(err.Error())
		}
		contentType := mime.TypeByExtension(path.Ext(fileInfo.Name()))

		b, err := ioutil.ReadAll(file)
		if err != nil {
			t.Fatal(err.Error())
		}
		w.Header().Add("Content-Type", contentType)
		w.Write(b)
	}))
	t.Cleanup(func() {
		ts.Close()
	})
	return ts
}
