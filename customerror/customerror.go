package customerror

import (
	"errors"
	"fmt"
)

type baseErrors interface {
	GetCode() int
	GetMessage() error
}

//自定义错误返回
type BaseError struct {
	code    int
	message error
}

func (err BaseError) Error() string {
	return fmt.Sprintf("状态码为：%d，错误原因为：%s", err.GetCode(), err.message.Error())
}

func (err *BaseError) GetCode() int {
	return err.code
}

func (err *BaseError) GetMessage() error {
	return err.message
}

func Custom(code int, msg error) BaseError {
	return BaseError{
		code,
		msg,
	}
}

func New(code int, err error) BaseError {
	return BaseError{
		code,
		err,
	}
}

var (
	Unauthorized   BaseError = New(401, errors.New("unauthorized"))
	ParameterError BaseError = New(403, errors.New("parameter error"))
)

// 标准错误
var (
	InsertError error = errors.New("数据主键已存在")
)


type ErrCode int
const (
	TransportErrCode ErrCode = 40000 + iota
)
