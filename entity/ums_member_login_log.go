package entity

import (
	"time"
)

type UmsMemberLoginLog struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	MemberId   int64     `json:"member_id" xorm:"default NULL BIGINT(20) 'member_id'"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' DATETIME 'create_time'"`
	Ip         string    `json:"ip" xorm:"default 'NULL' VARCHAR(64) 'ip'"`
	City       string    `json:"city" xorm:"default 'NULL' VARCHAR(64) 'city'"`
	LoginType  int       `json:"login_type" xorm:"default NULL comment('登录类型：0->PC；1->android;2->ios;3->小程序') INT(1) 'login_type'"`
	Province   string    `json:"province" xorm:"default 'NULL' VARCHAR(64) 'province'"`
}
