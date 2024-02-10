package approuter

import (
	"fmt"
	"io"
	"net/http"

	"github.com/andyinabox/linkydink/pkg/opmlparser"
	"github.com/gin-gonic/gin"
)

func (r *Router) OpmlPost(ctx *gin.Context) {
	id, _, err := r.hh.GetUserIdFromSession(ctx)
	if err != nil {
		r.InfoMessageError(ctx, http.StatusUnauthorized, err)
		return
	}

	// single file
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		r.InfoMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		r.InfoMessageError(ctx, http.StatusInternalServerError, err)
		return
	}
	defer file.Close()
	fileContents, err := io.ReadAll(file)
	if err != nil {
		r.InfoMessageError(ctx, http.StatusInternalServerError, err)
		return
	}

	feeds, err := opmlparser.ParseXml(fileContents)
	if err != nil {
		r.InfoMessageError(ctx, http.StatusInternalServerError, err)
		return
	}

	ls := r.sc.LinkService()
	for _, feed := range feeds {
		_, err = ls.CreateLinkFromFeed(id, feed.Title, feed.XmlUrl, feed.HtmlUrl)
		if err != nil {
			r.InfoMessageError(ctx, http.StatusInternalServerError, err)
			return
		}
	}

	r.InfoMessageSuccess(ctx, fmt.Sprintf("✅ Success! %d links added", len(feeds)))
}
