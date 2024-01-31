package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (a *App) ErrorResponse(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, gin.H{"success": false, "error": err})
}

func (a *App) SuccessResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

func (a *App) SuccessResponseJSON(ctx *gin.Context, payload interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"success": true, "payload": payload})
}

func (a *App) CreatedResponseJSON(ctx *gin.Context, payload interface{}) {
	ctx.JSON(http.StatusCreated, gin.H{"success": true, "payload": payload})
}

func (a *App) NotFoundResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"success": false, "error": "Not found"})
}

func (a *App) GetID(ctx *gin.Context) (uint, error) {
	str := ctx.Param("id")
	id, err := strconv.Atoi(str)
	return uint(id), err
}
