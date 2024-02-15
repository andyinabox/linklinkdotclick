package approuter

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/andyinabox/linkydink/pkg/postman"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	loginBtn  string = "btn-login"
	logoutBtn string = "btn-logout"
)

type EmailTemplateData struct {
	Subject   string
	MagicLink string
	ImageUrl  string
}

func (r *Router) SessionPost(ctx *gin.Context) {

	if ctx.PostForm(loginBtn) == "1" {
		r.login(ctx)
		return
	} else if ctx.PostForm(logoutBtn) == "1" {
		r.logout(ctx)
		return
	}

	err := errors.New("invalid post options")
	r.sc.LogService().Error().Println(err.Error())
	r.hrh.InfoPageError(ctx, http.StatusBadRequest, err)
}

func (r *Router) login(ctx *gin.Context) {
	logger := r.sc.LogService()
	var err error
	email := ctx.PostForm("email")
	if email == "" {
		err = errors.New("no email provided")
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusBadRequest, err)
		return
	}

	userService := r.sc.UserService()

	user, err := userService.FetchOrCreateUserByEmail(email)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err)
		return
	}

	hash, err := userService.GetLoginHashForUser(user)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err)
		return
	}

	magicLink := fmt.Sprintf("https://%s/session/?h=%s", ctx.Request.Host, hash)

	bodyBuffer := &bytes.Buffer{}
	bodyData := &EmailTemplateData{
		Subject:   "🖇 Your linklinkclick login link",
		MagicLink: magicLink,
		ImageUrl:  "https://" + ctx.Request.Host + "/static/android-chrome-192x192.png",
	}
	err = r.conf.Templates.ExecuteTemplate(bodyBuffer, "email.html.tmpl", bodyData)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err)
	}

	postmanEmail := &postman.Email{
		From: mail.Address{
			Name:    "Linky",
			Address: "noreply@" + ctx.Request.Host,
		},
		To: mail.Address{
			Name:    "You",
			Address: user.Email,
		},
		Subject: bodyData.Subject,
		Mime:    postman.MimeHtml,
		Body:    bodyBuffer.String(),
	}

	err = postman.Send(postmanEmail, r.conf.SmtpAddr)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err)
		return
	}

	r.hrh.InfoPageSuccess(ctx, "✨ A magic link is on its way to your inbox!")
}

func (r *Router) logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	ctx.Redirect(http.StatusSeeOther, "/")
}
