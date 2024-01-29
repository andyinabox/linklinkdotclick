package app

import (
	"encoding/json"
	"net/http"

	"github.com/andyinabox/linkydink-sketch/pkg/simpleserver"
)

func (a *App) ApiLinksGet(ctx *simpleserver.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// load data
		data, err := ctx.Resources.ReadFile("res/static/data.json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		testData := testData{}
		err = json.Unmarshal(data, &testData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		var linksJson []byte
		linksJson, err = json.Marshal(testData.Links)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// send response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(linksJson)
	}
}