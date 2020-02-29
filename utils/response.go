package utils

import (
	cerr "basic_mall/customerror"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AbortJSON 发生错误是否是返回
func AbortJSON(ctx *gin.Context, err cerr.BaseError) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": err.GetCode(),
		"msg":    err.GetMessage(),
		"data":   "",
	})
	ctx.Abort()
}

// ReturnJSON 正常返回json数据
func ReturnJSON(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "success",
		"data":   data,
	})
	ctx.Abort()
}
