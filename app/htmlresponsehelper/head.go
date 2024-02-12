package htmlresponsehelper

import "github.com/gin-gonic/gin"

type headRenderContext struct {
	Title                  string
	Description            string
	Url                    string
	FavIconUrl             string
	AppleTouchIconUrl      string
	ManifestUrl            string
	OgTitle                string
	OgImageUrl             string
	OgImageAlt             string
	RedirectTimeoutSeconds int
	RedirectUrl            string
	RedirectNoscriptOnly   bool
}

func (h *Helper) newHeadRenderContext(ctx *gin.Context) headRenderContext {
	return headRenderContext{
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
