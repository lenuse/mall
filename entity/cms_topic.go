package entity

import (
	"time"
)

type CmsTopic struct {
	Id             int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	CategoryId     int64     `json:"category_id" xorm:"default NULL BIGINT(20) 'category_id'"`
	Name           string    `json:"name" xorm:"default 'NULL' VARCHAR(255) 'name'"`
	CreateTime     time.Time `json:"create_time" xorm:"default 'NULL' DATETIME 'create_time'"`
	StartTime      time.Time `json:"start_time" xorm:"default 'NULL' DATETIME 'start_time'"`
	EndTime        time.Time `json:"end_time" xorm:"default 'NULL' DATETIME 'end_time'"`
	AttendCount    int       `json:"attend_count" xorm:"default NULL comment('参与人数') INT(11) 'attend_count'"`
	AttentionCount int       `json:"attention_count" xorm:"default NULL comment('关注人数') INT(11) 'attention_count'"`
	ReadCount      int       `json:"read_count" xorm:"default NULL INT(11) 'read_count'"`
	AwardName      string    `json:"award_name" xorm:"default 'NULL' comment('奖品名称') VARCHAR(100) 'award_name'"`
	AttendType     string    `json:"attend_type" xorm:"default 'NULL' comment('参与方式') VARCHAR(100) 'attend_type'"`
	Content        string    `json:"content" xorm:"comment('话题内容') TEXT 'content'"`
}
