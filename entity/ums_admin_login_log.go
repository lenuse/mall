package entity

import (
	"time"
)

type UmsAdminLoginLog struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	AdminId    int64     `json:"admin_id" xorm:"default NULL BIGINT(20) 'admin_id'"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' DATETIME 'create_time'"`
	Ip         string    `json:"ip" xorm:"default 'NULL' VARCHAR(64) 'ip'"`
	Address    string    `json:"address" xorm:"default 'NULL' VARCHAR(100) 'address'"`
	UserAgent  string    `json:"user_agent" xorm:"default 'NULL' comment('浏览器登录类型') VARCHAR(100) 'user_agent'"`
}
