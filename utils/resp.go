package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func NewRespJson(ctx *gin.Context, code StateCode, data interface{}, msg string)   {
	if msg == "" {
		msg = StatusText(code)
	}
	traceId,_ := ctx.Get(TraceIdKey)
	content := gin.H{
		"state":code,
		"message":msg,
		"data":data,
		"trace_id":traceId,
	}
	ctx.JSON(http.StatusOK,content)
	return
}