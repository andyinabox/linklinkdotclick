package htmlresponsehelper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pageLandingBody struct{}

func (h *Helper) PageLanding(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "base.html.tmpl", h.newRenderContext(ctx, pageLanding, &pageLandingBody{}))
}
