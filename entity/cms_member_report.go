package entity

import (
	"time"
)

type CmsMemberReport struct {
	Id               int64     `json:"id" xorm:"default NULL BIGINT(20) 'id'"`
	ReportType       int       `json:"report_type" xorm:"default NULL comment('举报类型：0->商品评价；1->话题内容；2->用户评论') INT(1) 'report_type'"`
	ReportMemberName string    `json:"report_member_name" xorm:"default 'NULL' comment('举报人') VARCHAR(100) 'report_member_name'"`
	CreateTime       time.Time `json:"create_time" xorm:"default 'NULL' DATETIME 'create_time'"`
	ReportObject     string    `json:"report_object" xorm:"default 'NULL' VARCHAR(100) 'report_object'"`
	ReportStatus     int       `json:"report_status" xorm:"default NULL comment('举报状态：0->未处理；1->已处理') INT(1) 'report_status'"`
	HandleStatus     int       `json:"handle_status" xorm:"default NULL comment('处理结果：0->无效；1->有效；2->恶意') INT(1) 'handle_status'"`
	Note             string    `json:"note" xorm:"default 'NULL' VARCHAR(200) 'note'"`
}
