package approuter

import (
	"net/http"

	"github.com/andyinabox/linkydink/pkg/opmlparser"
	"github.com/gin-gonic/gin"
)

func (r *Router) OpmlGet(ctx *gin.Context) {
	user, _, err := r.hh.GetUserFromSession(ctx)

	if err != nil {
		r.InfoMessageError(ctx, http.StatusUnauthorized, err)
		return
	}

	links, err := r.sc.LinkService().FetchLinks(user.ID)
	if err != nil {
		r.InfoMessageError(ctx, http.StatusInternalServerError, err)
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
		r.InfoMessageError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Writer.Header().Set("Content-Type", "application/xml")
	ctx.Writer.Write(b)
	ctx.Writer.WriteHeader(http.StatusOK)

}
