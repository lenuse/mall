package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lenuse/mall/transport"
	"upper.io/db.v3"
)

// StateCode 定义返回结构
type StateCode int

func (code StateCode) Int() int {
	return int(code)
}

// 通用100 前台200 后台300
const (
	Unauthorized    StateCode = 10040101
	Success         StateCode = 10020000
	InvalidArgument StateCode = 10040001
)

//前台100 500 错误
const (
	ApiError StateCode = 10050000 + iota
)

//后台300 500 错误
const (
	ManageError StateCode = 30050000 + iota
	AdminNotUnique
)

// 通用100 500 错误
const (
	GeneralError StateCode = 10050000 + iota
	JwtError
	InsertError
	QueryFail
)

//返回code解释
var statusText = map[StateCode]string{
	JwtError:        "签名生成错误",
	Success:         "成功",
	Unauthorized:    "登录验证失败",
	InvalidArgument: "参数错误",
	AdminNotUnique:  "用户账号重复",
	InsertError:     "存储失败",
	QueryFail:       "查询失败",
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

type MetaInfo struct {
	TraceId  string   `json:"trace_id"`
	Paginate Paginate `json:"paginate"`
}

type Response struct {
	State   StateCode   `json:"state"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    MetaInfo    `json:"meta"`
}

func (r *Response) SetMessage(msg string) *Response {
	r.Message = msg
	return r
}

func (r *Response) SetData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) SetPaginate(p Paginate) *Response {
	r.Meta.Paginate = p
	return r
}

func (r *Response) WriteJson(ctx *gin.Context) {
	trace, _ := ctx.Get(TraceIdKey)
	r.Meta.TraceId = trace.(string)
	ctx.JSON(http.StatusOK, *r)
}

// NewResponse 实例化返回
func NewResponse(code StateCode) *Response {
	msg := StatusText(code)
	return &Response{
		State:   code,
		Message: msg,
	}
}
