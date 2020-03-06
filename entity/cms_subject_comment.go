package entity

import (
	"time"
)

type CmsSubjectComment struct {
	Id             int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	SubjectId      int64     `json:"subject_id" xorm:"default NULL BIGINT(20) 'subject_id'"`
	MemberNickName string    `json:"member_nick_name" xorm:"default 'NULL' VARCHAR(255) 'member_nick_name'"`
	MemberIcon     string    `json:"member_icon" xorm:"default 'NULL' VARCHAR(255) 'member_icon'"`
	Content        string    `json:"content" xorm:"default 'NULL' VARCHAR(1000) 'content'"`
	CreateTime     time.Time `json:"create_time" xorm:"default 'NULL' DATETIME 'create_time'"`
	ShowStatus     int       `json:"show_status" xorm:"default NULL INT(1) 'show_status'"`
}
