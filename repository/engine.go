package repository

import (
	"github.com/lenuse/mall/config"
	_ "github.com/lib/pq"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

var engine sqlbuilder.Database

//Status 启用装填
type Status int8

// Int8 转化为int8
func (s Status) Int8() int8 {
	return int8(s)
}

func getConnURL() postgresql.ConnectionURL {
	options := map[string]string{
		"sslmode":     config.New().Datebases.Beta.Ssl,
		"search_path": config.New().Datebases.Beta.Schema,
	}
	return postgresql.ConnectionURL{
		Host:     config.New().Datebases.Beta.Host,
		Database: config.New().Datebases.Beta.DbName,
		User:     config.New().Datebases.Beta.User,
		Password: config.New().Datebases.Beta.Password,
		Options:  options,
	}
}

const (
	// NonStatus 无状态预留
	NonStatus Status = iota
	// EnableStatus 启用
	EnableStatus
	// DisableStatus 停用
	DisableStatus
)

// Init 初始化数据库连接
func Init() (err error) {
	settings := getConnURL()
	engine, err = postgresql.Open(settings)
	return
}

// Close 关闭数据库
func Close() error {
	return engine.Close()
}
