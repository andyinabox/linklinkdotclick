package approuter

import (
	"github.com/gin-gonic/gin"
)

func (r *Router) AboutGet(ctx *gin.Context) {
	r.hrh.AboutPage(ctx)
}
