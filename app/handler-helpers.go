package app

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
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

func (a *App) GetUserIdFromSession(ctx *gin.Context) (id uint, err error) {
	session := sessions.Default(ctx)
	value := session.Get("user")
	if value == nil {
		err = ErrUnauthorized
		return
	}
	var ok bool
	id, ok = value.(uint)
	if !ok {
		err = ErrServerError
		return
	}
	return
}

// if user is not found, use the default link service instead
func (a *App) GetUserLinkServiceFromSession(ctx *gin.Context) (LinkService, error) {
	user, err := a.GetUserFromSession(ctx)
	if err != nil {
		return a.ls, nil
	}
	return a.us.GetUserLinkService(user)
}

func (a *App) GetUserFromSession(ctx *gin.Context) (*User, error) {
	id, err := a.GetUserIdFromSession(ctx)
	if err != nil {
		return nil, err
	}
	return a.us.FetchUser(id)
}
