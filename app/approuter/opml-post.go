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
		r.hrh.PageInfoError(ctx, http.StatusUnauthorized, err)
		return
	}
	// single file
	fileHeader, err := ctx.FormFile("opml_file")
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.PageInfoError(ctx, http.StatusBadRequest, err)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
		return
	}
	defer file.Close()
	fileContents, err := io.ReadAll(file)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
		return
	}

	feeds, err := opmlparser.ParseXml(fileContents)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
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
		r.hrh.PageInfoSuccess(ctx, fmt.Sprintf("☹️ successfully parsed OPML file, but %d links successfully added", len(feeds)))
	} else if successCount > failedCount {
		r.hrh.PageInfoSuccess(ctx, fmt.Sprintf("✅ %d links added, %d failed", successCount, failedCount))
	} else {
		r.hrh.PageInfoSuccess(ctx, fmt.Sprintf("🫤 %d links added, %d failed", successCount, failedCount))
	}

}
