package approuter

import (
	"errors"
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/pkg/ginhelper"
	"github.com/gin-gonic/gin"
)

const (
	linkAddBtn    string = "btn-link-add"
	linkUpdateBtn string = "link-item-save"
	linkDeleteBtn string = "link-item-delete"
)

func (r *Router) LinksPost(ctx *gin.Context) {

	if ctx.PostForm(linkAddBtn) == "1" {
		r.createLink(ctx)
		return
	} else if ctx.PostForm(linkDeleteBtn) == "1" {
		r.deleteLink(ctx)
		return
	} else if ctx.PostForm(linkUpdateBtn) == "1" {
		r.updateLink(ctx)
		return
	}

	err := errors.New("invalid post options")
	r.sc.LogService().Error().Println(err.Error())
	r.hrh.InfoPageError(ctx, http.StatusBadRequest, err)
}

func (r *Router) createLink(ctx *gin.Context) {
	logger := r.sc.LogService()
	var err error

	originalUrl := ctx.PostForm("url")
	userId := ginhelper.GetUint(ctx, "userId")

	_, err = r.sc.LinkService().CreateLink(userId, originalUrl)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}

func (r *Router) updateLink(ctx *gin.Context) {
	logger := r.sc.LogService()

	userId := ginhelper.GetUint(ctx, "userId")

	var link app.Link
	err := ctx.ShouldBind(&link)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusBadRequest, err)
		return
	}

	_, err = r.sc.LinkService().UpdateLink(userId, link.ID, link, true)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}

func (r *Router) deleteLink(ctx *gin.Context) {
	logger := r.sc.LogService()

	userId := ginhelper.GetUint(ctx, "userId")
	id, err := ginhelper.GetPostFormUint(ctx, "id")
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusBadRequest, err)
		return
	}

	_, err = r.sc.LinkService().DeleteLink(userId, id)
	if err != nil {
		logger.Error().Println(err.Error())
		r.hrh.InfoPageError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/")
}
