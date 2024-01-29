package app

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/andyinabox/linkydink-sketch/pkg/simpleserver"
)

type apiLinksPostBody struct {
	Url string
}

func (a *App) ApiLinksPost(ctx *simpleserver.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		decoder := json.NewDecoder(r.Body)
		var body apiLinksPostBody

		err := decoder.Decode(&body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		result, feedUrl, err := a.feedreader.Parse(body.Url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if result.Feed == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Feed not found"))
			return
		}

		link := &Link{
			ID:          1,
			UserID:      1,
			SiteName:    strings.TrimSpace(result.Title),
			SiteUrl:     strings.TrimSpace(result.Link),
			FeedUrl:     feedUrl,
			OriginalUrl: body.Url,
			LastClicked: time.Now(),
			LastFetched: time.Now(),
		}

		var linkJson []byte
		linkJson, err = json.Marshal(link)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// send response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(linkJson)
	}
}
