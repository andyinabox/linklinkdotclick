package htmlresponsehelper

import (
	"github.com/gin-gonic/gin"
)

type footRenderContext struct {
	AppVersion string
	HideFooter bool
}

func (h *Helper) newFootRenderContext(ctx *gin.Context) footRenderContext {
	return footRenderContext{
		AppVersion: h.conf.AppVersion,
	}
}
