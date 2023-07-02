package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ForbiddenResp(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 403,
		"msg":  err.Error(),
	})
}

func SuccessResp(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": data,
	})
}

func EncryptResp(ctx *gin.Context, data CryptForm) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": data,
	})
}

func EncryptAnyResp(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": EncryptToForm(data),
	})
}
