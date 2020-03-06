package entity

import (
	"time"
)

type CmsHelp struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	CategoryId int64     `json:"category_id" xorm:"default NULL BIGINT(20) 'category_id'"`
	Icon       string    `json:"icon" xorm:"default 'NULL' VARCHAR(500) 'icon'"`
	Title      string    `json:"title" xorm:"default 'NULL' VARCHAR(100) 'title'"`
	ShowStatus int       `json:"show_status" xorm:"default NULL INT(1) 'show_status'"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' DATETIME 'create_time'"`
	ReadCount  int       `json:"read_count" xorm:"default NULL INT(1) 'read_count'"`
	Content    string    `json:"content" xorm:"TEXT 'content'"`
}
