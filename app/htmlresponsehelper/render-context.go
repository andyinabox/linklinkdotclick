package htmlresponsehelper

import "github.com/gin-gonic/gin"

type renderContext struct {
	Head headRenderContext
	Body any
	Foot footRenderContext
}

func (h *Helper) newRenderContext(ctx *gin.Context, body any, headOpts *headOptions) *renderContext {
	return &renderContext{
		h.newHeadRenderContext(ctx, headOpts),
		body,
		h.newFootRenderContext(ctx),
	}
}
