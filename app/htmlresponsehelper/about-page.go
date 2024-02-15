package htmlresponsehelper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type aboutPageBody struct{}

func (h *Helper) AboutPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "base.html.tmpl", h.newRenderContext(ctx, about, &aboutPageBody{}))
}
