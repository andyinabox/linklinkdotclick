package approuter

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (r *Router) SessionDelete(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	r.hrh.InfoPageSuccess(ctx, "👋 You're logged out. Later!", false)
}
