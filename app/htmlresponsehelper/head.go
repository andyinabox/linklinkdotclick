package htmlresponsehelper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	UserStyles             string
}

func (h *Helper) newHeadRenderContext(ctx *gin.Context) headRenderContext {

	userId := ctx.GetUint("userId")
	user, err := h.sc.UserService().FetchUser(userId)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
	}

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
		UserStyles:        user.StyleSheet,
	}
}
