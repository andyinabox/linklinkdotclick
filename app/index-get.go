package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexRenderContext struct {
	Links []Link
}

func (a *App) IndexGet(ctx *gin.Context) {

	var links []Link
	tx := a.db.Order("last_clicked").Find(&links)
	err := tx.Error
	if err != nil {
		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.HTML(http.StatusOK, "index.html.tmpl", &IndexRenderContext{
		Links: links,
	})
}
