package approuter

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/andyinabox/linkydink/pkg/mailservice"
	"github.com/gin-gonic/gin"
)

type EmailTemplateData struct {
	Subject   string
	MagicLink string
	ImageUrl  string
}

func (r *Router) LoginPost(ctx *gin.Context) {
	logger := r.sc.LogService()
	var err error
	email := ctx.PostForm("email")
	if email == "" {
		err = errors.New("no email provided")
		logger.Error().Println(err.Error())
		r.InfoMessageError(ctx, http.StatusBadRequest, err)
		return
	}

	userService := r.sc.UserService()

	user, err := userService.FetchOrCreateUserByEmail(email)
	if err != nil {
		logger.Error().Println(err.Error())
		r.InfoMessageError(ctx, http.StatusInternalServerError, err)
		return
	}

	hash, err := userService.GetLoginHashForUser(user)
	if err != nil {
		logger.Error().Println(err.Error())
		r.InfoMessageError(ctx, http.StatusInternalServerError, err)
		return
	}

	magicLink := fmt.Sprintf("https://%s/login/%s", ctx.Request.Host, hash)

	bodyBuffer := &bytes.Buffer{}
	bodyData := &EmailTemplateData{
		Subject:   "🖇 Your linklinkclick login link",
		MagicLink: magicLink,
		ImageUrl:  "https://" + ctx.Request.Host + "/static/android-chrome-192x192.png",
	}
	err = r.conf.Templates.ExecuteTemplate(bodyBuffer, "email.html.tmpl", bodyData)
	if err != nil {
		logger.Error().Println(err.Error())
		r.InfoMessageError(ctx, http.StatusInternalServerError, err)
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
		Mime:    mailservice.MimeHtml,
		Body:    bodyBuffer.String(),
	})
	if err != nil {
		logger.Error().Println(err.Error())
		r.InfoMessageError(ctx, http.StatusInternalServerError, err)
		return
	}

	r.InfoMessageSuccess(ctx, "✨ A magic link is on its way to your inbox!")
}
