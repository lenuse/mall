package entity

import (
	"time"
)

type UmsIntegrationChangeHistory struct {
	Id          int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	MemberId    int64     `json:"member_id" xorm:"default NULL BIGINT(20) 'member_id'"`
	CreateTime  time.Time `json:"create_time" xorm:"default 'NULL' DATETIME 'create_time'"`
	ChangeType  int       `json:"change_type" xorm:"default NULL comment('改变类型：0->增加；1->减少') INT(1) 'change_type'"`
	ChangeCount int       `json:"change_count" xorm:"default NULL comment('积分改变数量') INT(11) 'change_count'"`
	OperateMan  string    `json:"operate_man" xorm:"default 'NULL' comment('操作人员') VARCHAR(100) 'operate_man'"`
	OperateNote string    `json:"operate_note" xorm:"default 'NULL' comment('操作备注') VARCHAR(200) 'operate_note'"`
	SourceType  int       `json:"source_type" xorm:"default NULL comment('积分来源：0->购物；1->管理员修改') INT(1) 'source_type'"`
}
