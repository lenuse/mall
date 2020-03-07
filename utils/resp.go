package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StateCode 定义返回结构
type StateCode int

// 通用100 前台200 后台300
const (
	Unauthorized    StateCode = 10040101
	Success         StateCode = 10020001
	InvalidArgument StateCode = 10040001
	JwtError        StateCode = 10050001
	InsertError     StateCode = 10050002
	AdminNotUnique  StateCode = 30050001
)

//返回code解释
var statusText = map[StateCode]string{
	JwtError:        "签名生成错误",
	Success:         "成功",
	Unauthorized:    "登录验证失败",
	InvalidArgument: "参数错误",
	AdminNotUnique:  "用户账号重复",
	InsertError:     "存储失败",
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
