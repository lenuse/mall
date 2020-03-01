package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StateCode 定义返回结构
type StateCode int

// 通用100 前台200 后台300
const (
	StateCodeUnauthorized    StateCode = 10040101
	StateCodeSuccess         StateCode = 10020001
	StateCodeInvalidArgument StateCode = 10040001
	StateCodeJwtError        StateCode = 10050001
)

//返回code解释
var statusText = map[StateCode]string{
	StateCodeJwtError:        "签名生成错误",
	StateCodeSuccess:         "成功",
	StateCodeUnauthorized:    "登录验证失败",
	StateCodeInvalidArgument: "参数错误",
}

// StatusText 获取code的解释
func StatusText(code StateCode) string {
	if text, ok := statusText[code]; ok {
		return text
	}
	return "未知错误"
}

// NewRespJSON 实例化返回
func NewRespJSON(ctx *gin.Context, code StateCode, data interface{}, msg string) {
	if msg == "" {
		msg = StatusText(code)
	}
	traceID, _ := ctx.Get(TraceIdKey)
	content := gin.H{
		"state":    code,
		"message":  msg,
		"data":     data,
		"trace_id": traceID,
	}
	ctx.JSON(http.StatusOK, content)
	return
}
