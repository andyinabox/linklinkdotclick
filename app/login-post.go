package app

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

func (a *App) LoginPost(ctx *gin.Context) {
	email := ctx.PostForm("email")
	if email == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userService := a.sc.UserService()

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

	err = a.sc.MailService().Send(&mailservice.Email{
		From:    mail.Address{"Linky", "linky@" + a.conf.Domain},
		To:      mail.Address{"You", user.Email},
		Subject: a.conf.Domain + " magic login link ✨",
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
