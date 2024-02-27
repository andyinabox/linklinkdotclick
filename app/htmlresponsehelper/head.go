package htmlresponsehelper

import (
	"github.com/gin-gonic/gin"
)

type headRenderContext struct {
	Title             string
	Description       string
	Url               string
	FavIconUrl        string
	AppleTouchIconUrl string
	ManifestUrl       string
	OgTitle           string
	OgImageUrl        string
	OgImageAlt        string
	UserStyles        string
	AppVersion        string
}

func (h *Helper) newHeadRenderContext(ctx *gin.Context) headRenderContext {

	headerContext := headRenderContext{
		Title:             h.conf.SiteTitle,
		Description:       h.conf.Description,
		Url:               "https://" + ctx.Request.Host,
		FavIconUrl:        h.conf.FavIconUrl,
		AppleTouchIconUrl: h.conf.AppleTouchIconUrl,
		ManifestUrl:       h.conf.ManifestUrl,
		OgTitle:           h.conf.SiteTitle,
		OgImageUrl:        h.conf.OgImageUrl,
		OgImageAlt:        h.conf.OgImageAlt,
		AppVersion:        h.conf.AppVersion,
	}

	if h.ah.IsAuthenticated(ctx) {
		user, err := h.ah.User(ctx)
		if err == nil {
			headerContext.UserStyles = user.StyleSheet
		}
	}

	return headerContext
}
