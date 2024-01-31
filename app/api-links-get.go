package app

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) ApiLinksGet(ctx *gin.Context) {

	var links []Link
	tx := a.db.Find(&links)
	err := tx.Error
	if err != nil {
		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	var responseData []byte
	responseData, err = json.Marshal(links)
	if err != nil {
		a.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	a.SuccessResponseJSON(ctx, responseData)
}
