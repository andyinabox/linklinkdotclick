package htmlresponsehelper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pageAboutBody struct{}

func (h *Helper) PageAbout(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "base.html.tmpl", h.newRenderContext(ctx, pageAbout, &pageAboutBody{}))
}
