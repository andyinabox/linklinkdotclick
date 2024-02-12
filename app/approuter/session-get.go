package approuter

import (
	"errors"
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (r *Router) SessionGetHash(ctx *gin.Context) {
	logger := r.sc.LogService()

	hash := ctx.Param("hash")
	user, err := r.sc.UserService().GetUserFromLoginHash(hash)

	if err == nil && user == nil {
		err = errors.New("no user found for hash")
	}

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
