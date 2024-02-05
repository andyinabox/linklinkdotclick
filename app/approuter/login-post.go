package approuter

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/andyinabox/linkydink/pkg/mailservice"
	"github.com/gin-gonic/gin"
)

var emailBodyTemplate = `
Here you go!

%s

🖇
`

func (r *Router) LoginPost(ctx *gin.Context) {
	email := ctx.PostForm("email")
	if email == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userService := r.sc.UserService()

	user, err := userService.FetchOrCreateUserByEmail(email)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	hash, err := userService.GetLoginHashForUser(user)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	magicLink := fmt.Sprintf("https://%s/login/%s", ctx.Request.Host, hash)

	err = r.sc.MailService().Send(&mailservice.Email{
		From: mail.Address{
			Name:    "Linky",
			Address: "noreply@" + ctx.Request.Host,
		},
		To: mail.Address{
			Name:    "You",
			Address: user.Email,
		},
		Subject: ctx.Request.Host + " magic login link ✨",
		Body:    fmt.Sprintf(emailBodyTemplate, magicLink),
	})
	if err != nil {
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}
