package approuter

import (
	"github.com/gin-gonic/gin"
)

type HeadRenderContext struct {
	Title             string
	Description       string
	Url               string
	FavIconUrl        string
	AppleTouchIconUrl string
	ManifestUrl       string
	OgTitle           string
	OgImageUrl        string
	OgImageAlt        string
}

type FootRenderContext struct {
	Version string
}

func (r *Router) NewFootRenderContext(ctx *gin.Context) FootRenderContext {
	return FootRenderContext{
		Version: r.conf.Version,
	}
}

func (r *Router) NewHeadRenderContext(ctx *gin.Context) HeadRenderContext {
	title := "linklink.click"
	return HeadRenderContext{
		Title:             title,
		Description:       "Somewhere in-between a blogroll and an RSS reader",
		Url:               "https://" + ctx.Request.Host,
		FavIconUrl:        "/static/favicon.ico",
		AppleTouchIconUrl: "/static/apple-touch-icon.png",
		ManifestUrl:       "/static/site.webmanifest",
		OgTitle:           title,
		OgImageUrl:        "/static/android-chrome-512x512.png",
		OgImageAlt:        "Two paperclips entwined",
	}
}
