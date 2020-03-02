package models

import (
	"fmt"
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var engine *xorm.Engine

//Status 启用装填
type Status int8

// Int8 转化为int8
func (s Status) Int8() int8 {
	return int8(s)
}

const (
	// NonStatus 无状态预留
	NonStatus Status = iota
	// EnableStatus 启用
	EnableStatus
	// DisableStatus 停用
	DisableStatus
)

func init() {
	var err error
	engine, err = xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/test2?sslmode=disable")
	if err != nil {
		_ = fmt.Errorf("数据库连接失败，%s", err.Error())
		os.Exit(500)
	}
}
