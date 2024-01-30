package app

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/andyinabox/linkydink/pkg/simpleserver"
)

type ApiLinksPostBody struct {
	Url string
}

func (a *App) ApiLinksPost(ctx *simpleserver.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		decoder := json.NewDecoder(r.Body)
		var body ApiLinksPostBody

		err := decoder.Decode(&body)
		if err != nil {
			ctx.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		if body.Url == "" {
			ctx.WriteError(w, http.StatusBadRequest, errors.New("missing url"))
			return
		}

		feed, feedUrl, err := a.feedreader.Parse(body.Url)
		if err != nil {
			ctx.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		if feed == nil {
			ctx.WriteError(w, http.StatusBadRequest, errors.New("feed not found"))
			return
		}

		link := &Link{
			SiteName:    strings.TrimSpace(feed.Title),
			SiteUrl:     strings.TrimSpace(feed.Link),
			FeedUrl:     feedUrl,
			OriginalUrl: body.Url,
			UnreadCount: int16(len(feed.Items)),
			LastClicked: time.Date(1993, time.April, 30, 12, 0, 0, 0, time.UTC),
			LastFetched: time.Now(),
		}

		tx := a.db.Create(link)
		err = tx.Error
		if err != nil {
			ctx.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		var responseData []byte
		responseData, err = json.Marshal(link)
		if err != nil {
			ctx.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		// send response
		ctx.WriteJSON(w, responseData)
	}
}
