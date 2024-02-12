package ginhelper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamUint(ctx *gin.Context, key string) (uint, error) {
	str := ctx.Param(key)
	id, err := strconv.Atoi(str)
	return uint(id), err
}
