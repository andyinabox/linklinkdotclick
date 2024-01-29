package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/andyinabox/linkydink-sketch/pkg/simpleserver"
)

func (a *App) ApiLinksIdGet(ctx *simpleserver.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathVars := ctx.Vars(r)
		queryVars, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			fmt.Printf("Error parsing query vars: %s", err.Error())
		}

		id, err := strconv.Atoi(pathVars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

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

		var link *Link
		for _, l := range testData.Links {
			if l.ID == int64(id) {
				link = &l
			}
		}

		if link == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if queryVars.Has("refresh") {
			fmt.Printf("Fetching rss feed for %s\n", link.FeedUrl)
			result, err := a.feedreader.ParseFeedUrl(link.FeedUrl)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			if result.Feed != nil {
				link.LastFetched = time.Now()
				link.UnreadCount = int16(len(result.Feed.Items))
			} else {
				fmt.Printf("No rss results for %s\n", link.FeedUrl)
			}

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
