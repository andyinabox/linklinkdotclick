package htmlresponsehelper

import "github.com/gin-gonic/gin"

type FootRenderContext struct {
	Version string
}

func (h *Helper) NewFootRenderContext(ctx *gin.Context) FootRenderContext {
	return FootRenderContext{
		Version: h.conf.AppVersion,
	}
}
