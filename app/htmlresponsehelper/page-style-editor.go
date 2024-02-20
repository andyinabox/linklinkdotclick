package htmlresponsehelper

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-gonic/gin"
)

type pageStyleEditorBody struct {
	User *app.User
}

func (h *Helper) PageStyleEditor(ctx *gin.Context, user *app.User) {
	ctx.HTML(http.StatusOK, "base.html.tmpl", h.newRenderContext(ctx, styleEditor, &pageStyleEditorBody{user}))
}
