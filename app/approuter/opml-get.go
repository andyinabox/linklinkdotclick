package approuter

import (
	"net/http"

	"github.com/andyinabox/linkydink/pkg/opmlparser"
	"github.com/gin-gonic/gin"
)

func (r *Router) OpmlGet(ctx *gin.Context) {
	logger := r.sc.LogService()
	user, err := r.ah.User(ctx)

	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.PageInfoError(ctx, http.StatusUnauthorized, err)
		return
	}

	links, err := r.sc.LinkService().FetchLinks(user.ID)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
		return
	}

	feeds := make([]opmlparser.Feed, len(links))
	for i, link := range links {
		feeds[i] = opmlparser.Feed{
			Title:   link.SiteName,
			XmlUrl:  link.FeedUrl,
			HtmlUrl: link.SiteUrl,
		}
	}

	b, err := opmlparser.MarshallXml(feeds, user.SiteTitle)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.PageInfoError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Writer.Header().Set("Content-Type", "application/xml")
	ctx.Writer.Header().Set("Content-Disposition", "attachment; filename=\"my-linklinkclick-links.opml\"")
	ctx.Writer.Write(b)
	ctx.Writer.WriteHeader(http.StatusOK)

}
