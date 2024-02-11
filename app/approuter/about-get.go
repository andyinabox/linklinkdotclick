package approuter

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

func (r *Router) AboutGet(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "about.html.tmpl", &AboutPageRenderContext{
		r.NewHeadRenderContext(ctx),
		AboutPageBody{},
		r.NewFootRenderContext(ctx),
	})
}
