package approuter

import (
	"errors"
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-gonic/gin"
)

type SiteDataImages struct {
	OgImage        string
	OgImageAlt     string
	FavIcon        string
	AppleTouchIcon string
}

type SiteData struct {
	Url         string
	Description string
	Manifest    string
	OgTitle     string
	Images      SiteDataImages
}

type IndexRenderContext struct {
	Links         []app.Link
	DummyLink     app.Link
	User          *app.User
	IsDefaultUser bool
	Site          SiteData
}

func (r *Router) IndexGet(ctx *gin.Context) {

	user, isDefaultUser, err := r.hh.GetUserFromSession(ctx)
	// it's ok if no user is found, but we want to abort for server errors
	if err != nil && errors.Is(err, app.ErrServerError) {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	links, err := r.sc.LinkService().FetchLinks(user.ID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.HTML(http.StatusOK, "index.html.tmpl", &IndexRenderContext{
		Links:         links,
		User:          user,
		IsDefaultUser: isDefaultUser,
		Site: SiteData{
			Url:         "https://linklink.click",
			OgTitle:     "linklink.click",
			Description: "Somewhere in-between a link list and RSS reader",
			Manifest:    "/static/site.webmanifest",
			Images: SiteDataImages{
				OgImage:        "/static/android-chrome-512x512.png",
				OgImageAlt:     "Two paperclips entwined",
				FavIcon:        "/static/favicon.ico",
				AppleTouchIcon: "/static/apple-touch-icon.png",
			},
		},
	})
}
