package htmlresponsehelper

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-gonic/gin"
)

type pageHomeBody struct {
	User      app.User
	Links     []app.Link
	EmptyLink app.Link
	EditMode  bool
}

func (h *Helper) PageHome(ctx *gin.Context, user *app.User, isDefaultUser bool, links []app.Link, editMode bool) {
	ctx.HTML(http.StatusOK, "base.html.tmpl", h.newRenderContext(ctx, pageHome, &pageHomeBody{
		User:     *user,
		Links:    links,
		EditMode: editMode,
	}))
}
