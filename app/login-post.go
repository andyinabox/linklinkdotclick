package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var emailBodyTemplate = `
Here you go!

https://%s/login/%s

🖇
`

func (a *App) LoginPost(ctx *gin.Context) {
	email := ctx.PostForm("email")

	user, err := a.us.FetchOrCreateUserByEmail(email)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	hash, err := a.us.GetLoginHashForUser(user)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// err := a.ms.Send(&mailservice.Email{
	// 	From:    mail.Address{"Linky", "linky@" + a.conf.Domain},
	// 	To:      mail.Address{"You", user.Email},
	// 	Subject: c.conf.Domain + " magic login link ✨",
	// 	Body:    fmt.Sprintf(emailBodyTemplate, a.conf.Domain, hash),
	// })

	magicLink := fmt.Sprintf("https://%s/%s", ctx.Request.Host, hash)
	ctx.String(http.StatusOK, magicLink)
}
