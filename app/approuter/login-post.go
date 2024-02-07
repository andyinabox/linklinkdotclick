package approuter

import (
	"bytes"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/andyinabox/linkydink/pkg/mailservice"
	"github.com/gin-gonic/gin"
)

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

	bodyData := struct {
		Subject   string
		MagicLink string
		ImageUrl  string
	}{
		Subject:   "Your linklink.click magic login link ✨",
		MagicLink: magicLink,
		ImageUrl:  ctx.Request.Proto + "://" + ctx.Request.Host + "/static/android-chrome-192x192.png",
	}

	bodyBuffer := &bytes.Buffer{}
	err = r.conf.Templates.ExecuteTemplate(bodyBuffer, "email.html.tmpl", bodyData)
	if err != nil {
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	err = r.sc.MailService().Send(&mailservice.Email{
		From: mail.Address{
			Name:    "Linky",
			Address: "noreply@" + ctx.Request.Host,
		},
		To: mail.Address{
			Name:    "You",
			Address: user.Email,
		},
		Subject: bodyData.Subject,
		Mime:    "",     // mailservice.MimeHtml,
		Body:    "Test", // bodyBuffer.String(),
	})
	if err != nil {
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}
