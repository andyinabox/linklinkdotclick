package htmlresponsehelper

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-gonic/gin"
)

type HomePageBody struct {
	User          app.User
	Links         []app.Link
	EmptyLink     app.Link
	IsDefaultUser bool
}

type HomePageRenderContext struct {
	Head HeadRenderContext
	Body HomePageBody
	Foot FootRenderContext
}

func (h *Helper) HomePage(ctx *gin.Context, user *app.User, isDefaultUser bool, links []app.Link) {
	ctx.HTML(http.StatusOK, "index.html.tmpl", &HomePageRenderContext{
		h.NewHeadRenderContext(ctx),
		HomePageBody{
			User:          *user,
			Links:         links,
			IsDefaultUser: isDefaultUser,
		},
		h.NewFootRenderContext(ctx),
	})
}
