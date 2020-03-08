package utils

import (
	"github.com/lenuse/mall/transport"
	"upper.io/db.v3"
)

// StateCode 定义返回结构
type StateCode int

// 通用100 前台200 后台300
const (
	Unauthorized    StateCode = 10040101
	Success         StateCode = 10020000
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

type Paginate struct {
	PerPage      uint   `json:"per_page"`
	Page         uint   `json:"Page"`
	TotalPages   uint   `json:"total_pages"`
	TotalEntries uint64 `json:"total_entries"`
}

func NewPaginate(result db.Result, page transport.Page) Paginate {
	totalPages, _ := result.TotalPages()
	totalEntries, _ := result.TotalEntries()
	return Paginate{
		PerPage:      page.GetPageSize(),
		Page:         page.GetPageNumber(),
		TotalPages:   totalPages,
		TotalEntries: totalEntries,
	}
}

type RespJson struct {
	State   StateCode   `json:"state"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    struct {
		TraceId  string   `json:"trace_id"`
		Paginate Paginate `json:"paginate"`
	} `json:"meta"`
}

func (r *RespJson) SetMessage(msg string) {
	r.Message = msg
	return
}

// NewRespJSON 实例化返回
func NewRespJSON(code StateCode) RespJson {
	msg := StatusText(code)
	return RespJson{
		State:   code,
		Message: msg,
	}
}
