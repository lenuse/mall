package models

import (
	"fmt"
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/test2?sslmode=disable")
	if err != nil {
		_ = fmt.Errorf("数据库连接失败，%s", err.Error())
		os.Exit(500)
	}
}
