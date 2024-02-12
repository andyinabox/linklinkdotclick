package htmlresponsehelper

import "github.com/gin-gonic/gin"

type footRenderContext struct {
	Version string
}

func (h *Helper) newFootRenderContext(ctx *gin.Context) footRenderContext {
	return footRenderContext{
		Version: h.conf.AppVersion,
	}
}
