package entity

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"upper.io/db.v3/postgresql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

//Status 启用装填
type Status int8

// Int8 转化为int8
func (s Status) Int8() int8 {
	return int8(s)
}

var pgOptions = map[string]string{
	"sslmode":     "disable",
	"search_path": "mall",
}

var settings = postgresql.ConnectionURL{
	Host:     "demo.upper.io",
	Database: "booktown",
	User:     "demouser",
	Password: "demop4ss",
	Options:  pgOptions,
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
	engine, err = xorm.NewEngine("postgres", "postgres://postgres:mysecret@118.24.1.111:5432/postgres?sslmode=disable")
	if err != nil {
		_ = fmt.Errorf("数据库连接失败，%s", err.Error())
		os.Exit(500)
	}
}
