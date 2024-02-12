package htmlresponsehelper

import "github.com/gin-gonic/gin"

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

func (h *Helper) NewHeadRenderContext(ctx *gin.Context) HeadRenderContext {
	return HeadRenderContext{
		Title:             h.conf.SiteTitle,
		Description:       h.conf.Description,
		Url:               "https://" + ctx.Request.Host,
		FavIconUrl:        h.conf.FavIconUrl,
		AppleTouchIconUrl: h.conf.AppleTouchIconUrl,
		ManifestUrl:       h.conf.ManifestUrl,
		OgTitle:           h.conf.SiteTitle,
		OgImageUrl:        h.conf.OgImageUrl,
		OgImageAlt:        h.conf.OgImageAlt,
	}
}
