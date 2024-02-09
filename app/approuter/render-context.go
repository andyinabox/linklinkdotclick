package approuter

import (
	"github.com/andyinabox/linkydink/app"
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

type HomePageBody struct {
	User          app.User
	Links         []app.Link
	EmptyLink     app.Link
	IsDefaultUser bool
}

type HomePageRenderContext struct {
	Head HeadRenderContext
	Body HomePageBody
}

type InfoPageBody struct {
	Message  string
	LinkUrl  string
	LinkText string
}

type InfoPageRenderContext struct {
	Head HeadRenderContext
	Body InfoPageBody
}

func NewHeadRenderContext(ctx *gin.Context) HeadRenderContext {
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
