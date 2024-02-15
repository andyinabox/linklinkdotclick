package approuter

import (
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type sessionGetQuery struct {
	Hash string `form:"h" binding:"required"`
}

func (r *Router) SessionGet(ctx *gin.Context) {
	logger := r.sc.LogService()

	var query sessionGetQuery
	err := ctx.BindQuery(&query)
	if err != nil {
		r.hrh.InfoPageError(ctx, http.StatusBadRequest, err)
	}

	user, err := r.sc.UserService().GetUserFromLoginHash(query.Hash)
	if err != nil {
		logger.Error().Println(err.Error())

		r.hrh.InfoPage(ctx, http.StatusUnauthorized, &app.HtmlInfoMessageOptions{
			Message:  "🔒 That link isn't working. Did you already use it?",
			Error:    err,
			LinkUrl:  "/",
			LinkText: "Try sending a new one",
		})
		return
	}

	session := sessions.Default(ctx)
	session.Set("user", user.ID)
	session.Save()

	ctx.Redirect(http.StatusSeeOther, "/")
}
