package app

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JSONResponseBody struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

func (a *App) ResponseJSON(ctx *gin.Context, code int, success bool, err error, data interface{}) {

	response := &JSONResponseBody{
		Success: success,
		Payload: data,
	}

	if err != nil {
		response.Error = err.Error()
	}

	ctx.JSON(code, response)
}

func (a *App) ErrorResponse(ctx *gin.Context, code int, err error) {
	a.ResponseJSON(ctx, code, false, err, nil)
}

func (a *App) SuccessResponse(ctx *gin.Context) {
	a.ResponseJSON(ctx, http.StatusOK, true, nil, nil)
}

func (a *App) SuccessResponseJSON(ctx *gin.Context, data interface{}) {
	a.ResponseJSON(ctx, http.StatusOK, true, nil, data)
}

func (a *App) CreatedResponseJSON(ctx *gin.Context, data interface{}) {
	a.ResponseJSON(ctx, http.StatusCreated, true, nil, data)
}

func (a *App) NotFoundResponse(ctx *gin.Context) {
	a.ResponseJSON(ctx, http.StatusNotFound, false, errors.New("not found"), nil)
}

func (a *App) GetID(ctx *gin.Context) (uint, error) {
	str := ctx.Param("id")
	id, err := strconv.Atoi(str)
	return uint(id), err
}
