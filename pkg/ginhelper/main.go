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

func GetQueryBool(ctx *gin.Context, key string) bool {
	result, _ := ctx.GetQuery(key)
	if result == "true" || result == "1" {
		return true
	}
	return false
}
