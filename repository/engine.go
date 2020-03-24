package repository

import (
	"github.com/lenuse/mall/config"
	_ "github.com/lib/pq"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

var engine sqlbuilder.Database

//Status 启用装填
type Status int

// Int8 转化为int8
func (s Status) Int() int {
	return int(s)
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

//状态
const (
	// NonStatus 无状态预留
	NonStatus Status = iota
	// EnableStatus 启用
	EnableStatus
	// DisableStatus 停用
	DisableStatus
)

// sort排序
const (
	SortASC  string = "sort"
	SortDesc string = "-sort"
)

// Open 初始化数据库连接
func Open() error {
	settings := getConnURL()
	engine, err := postgresql.Open(settings)
	if err != nil {
		return err
	}
	engine.SetLogging(true)
	return nil
}

// Close 关闭数据库
func Close() error {
	return engine.Close()
}

type Repository interface {
	Save() error
	Delete() error
	Update() error
}
