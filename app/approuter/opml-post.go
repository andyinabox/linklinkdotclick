package approuter

import (
	"fmt"
	"io"
	"net/http"

	"github.com/andyinabox/linkydink/pkg/opmlparser"
	"github.com/gin-gonic/gin"
)

func (r *Router) OpmlPost(ctx *gin.Context) {
	logger := r.sc.LogService()

	id, _, err := r.ah.GetUserIdFromSession(ctx)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusUnauthorized, err, false)
		return
	}
	// single file
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusBadRequest, err, false)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err, false)
		return
	}
	defer file.Close()
	fileContents, err := io.ReadAll(file)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err, false)
		return
	}

	feeds, err := opmlparser.ParseXml(fileContents)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err, false)
		return
	}

	failedCount := 0
	successCount := 0
	ls := r.sc.LinkService()
	for _, feed := range feeds {
		_, err = ls.CreateLinkFromFeed(id, feed.Title, feed.XmlUrl, feed.HtmlUrl)
		if err != nil {
			logger.Info().Printf("opml feed upload failed for %s: %s / %s", feed.Title, feed.XmlUrl, feed.HtmlUrl)
			failedCount++
		} else {
			successCount++
		}
	}

	if successCount == 0 {
		r.hrh.InfoPageSuccess(ctx, fmt.Sprintf("☹️ successfully parsed OPML file, but %d links successfully added", len(feeds)), true)
	} else if successCount > failedCount {
		r.hrh.InfoPageSuccess(ctx, fmt.Sprintf("✅ %d links added, %d failed", successCount, failedCount), true)
	} else {
		r.hrh.InfoPageSuccess(ctx, fmt.Sprintf("🫤 %d links added, %d failed", successCount, failedCount), true)
	}

}
