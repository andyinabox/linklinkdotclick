package approuter

import (
	"errors"
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

func (r *Router) IndexGet(ctx *gin.Context) {
	logger := r.sc.LogService()

	user, isDefaultUser, err := r.hh.GetUserFromSession(ctx)

	// it's ok if no user is found, but we want to abort for server errors
	if err != nil && errors.Is(err, app.ErrServerError) {
		logger.Error().Println(err.Error())
		r.InfoMessageError(ctx, http.StatusInternalServerError, err)
		return
	}

	links, err := r.sc.LinkService().FetchLinks(user.ID)
	if err != nil {
		logger.Error().Println(err.Error())
		r.InfoMessageError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.HTML(http.StatusOK, "index.html.tmpl", &HomePageRenderContext{
		r.NewHeadRenderContext(ctx),
		HomePageBody{
			User:          *user,
			Links:         links,
			IsDefaultUser: isDefaultUser,
		},
		r.NewFootRenderContext(ctx),
	})
}
