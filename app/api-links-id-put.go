package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) ApiLinksIdPut(ctx *gin.Context) {
	_, refresh := ctx.GetQuery("refresh")

	id, err := a.GetID(ctx)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	var link Link

	// // look for link in database, if not send 404
	// tx := a.db.First(&link, id)
	// err = tx.Error
	// if err != nil {
	// 	a.NotFoundResponse(ctx)
	// 	return
	// }

	err = ctx.BindJSON(&link)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// // if the ids don't match, send bad request error
	// if link.ID != uint(id) {
	// 	a.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("id in request body (%d) does not match url (%d)", link.ID, id))
	// 	return
	// }

	// // refresh feed data if indicated
	// // refreshing will automaticaly save
	// if refresh {
	// 	err = a.RefreshLink(&link)
	// 	if err != nil {
	// 		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
	// 		return
	// 	}
	// 	// otherwise save without refresh
	// } else {
	// 	tx = a.db.Save(&link)
	// 	err = tx.Error
	// 	if err != nil {
	// 		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
	// 		return
	// 	}
	// }

	updatedLink, err := a.ls.UpdateLink(id, link, refresh)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	// send response
	a.SuccessResponseJSON(ctx, updatedLink)
}
