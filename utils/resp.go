package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//定义返回结构
type StateCode int

const (
	// 通用100 前台200 后台300
	StateCodeUnauthorized StateCode = 100401
	StateCodeSuccess      StateCode = 100200
)

//返回code解释
var statusText = map[StateCode]string{
	StateCodeSuccess:      "成功",
	StateCodeUnauthorized: "登录验证失败",
}

func StatusText(code StateCode) string {
	if text, ok := statusText[code]; ok {
		return text
	}
	return "未知错误"
}

//NewResqJson 构造返回json
func NewResqJson(ctx *gin.Context, code StateCode, data interface{}, message string) {
	traceId, _ := ctx.Get("traceId")
	if message == "" {
		message = StatusText(code)
	}
	respMap := gin.H{
		"state":   code,
		"message": message,
		"data":    data,
		"traceId": traceId,
	}
	ctx.JSON(http.StatusOK, respMap)
}
