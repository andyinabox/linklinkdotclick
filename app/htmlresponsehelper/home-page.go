package htmlresponsehelper

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-gonic/gin"
)

type homePageBody struct {
	User          app.User
	Links         []app.Link
	EmptyLink     app.Link
	IsDefaultUser bool
	EditMode      bool
}

func (h *Helper) HomePage(ctx *gin.Context, user *app.User, isDefaultUser bool, links []app.Link, editMode bool) {
	ctx.HTML(http.StatusOK, "index.html.tmpl", h.newRenderContext(ctx, &homePageBody{
		User:          *user,
		Links:         links,
		IsDefaultUser: isDefaultUser,
		EditMode:      editMode,
	}, &headOptions{
		RedirectTimeoutSeconds: 60,
		RedirectUrl:            "/",
		RedirectNoscriptOnly:   true,
	}))
}
