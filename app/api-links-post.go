package app

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiLinksPostBody struct {
	Url string
}

func (a *App) ApiLinksPost(ctx *gin.Context) {

	var body ApiLinksPostBody
	err := ctx.BindJSON(&body)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	if body.Url == "" {
		a.ErrorResponse(ctx, http.StatusBadRequest, errors.New("missing url"))
		return
	}

	// feed, feedUrl, err := a.feedreader.Parse(body.Url)
	// if err != nil {
	// 	a.ErrorResponse(ctx, http.StatusInternalServerError, err)
	// 	return
	// }
	// if feed == nil {
	// 	a.ErrorResponse(ctx, http.StatusUnprocessableEntity, errors.New("no feed detectd"))
	// 	return
	// }

	// link := &Link{
	// 	SiteName:    strings.TrimSpace(feed.Title),
	// 	SiteUrl:     strings.TrimSpace(feed.Link),
	// 	FeedUrl:     feedUrl,
	// 	OriginalUrl: body.Url,
	// 	UnreadCount: int16(len(feed.Items)),
	// 	LastClicked: time.Date(1993, time.April, 30, 12, 0, 0, 0, time.UTC),
	// 	LastFetched: time.Now(),
	// }

	// tx := a.db.Create(link)
	// err = tx.Error
	// if err != nil {
	// 	a.ErrorResponse(ctx, http.StatusInternalServerError, err)
	// 	return
	// }

	link, err := a.ls.CreateLink(body.Url)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	a.CreatedResponseJSON(ctx, link)

}
