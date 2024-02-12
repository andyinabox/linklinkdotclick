package htmlresponsehelper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AboutPageBody struct {
}

type AboutPageRenderContext struct {
	Head HeadRenderContext
	Body AboutPageBody
	Foot FootRenderContext
}

func (h *Helper) AboutPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "about.html.tmpl", &AboutPageRenderContext{
		h.NewHeadRenderContext(ctx),
		AboutPageBody{},
		h.NewFootRenderContext(ctx),
	})
}
