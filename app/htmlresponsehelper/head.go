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
}

type headOptions struct {
	RedirectTimeoutSeconds int
	RedirectUrl            string
}

func (h *Helper) newHeadRenderContext(ctx *gin.Context, opts *headOptions) headRenderContext {
	return headRenderContext{
		Title:                  h.conf.SiteTitle,
		Description:            h.conf.Description,
		Url:                    "https://" + ctx.Request.Host,
		FavIconUrl:             h.conf.FavIconUrl,
		AppleTouchIconUrl:      h.conf.AppleTouchIconUrl,
		ManifestUrl:            h.conf.ManifestUrl,
		OgTitle:                h.conf.SiteTitle,
		OgImageUrl:             h.conf.OgImageUrl,
		OgImageAlt:             h.conf.OgImageAlt,
		RedirectTimeoutSeconds: opts.RedirectTimeoutSeconds,
		RedirectUrl:            "https://" + ctx.Request.Host + opts.RedirectUrl,
	}
}
