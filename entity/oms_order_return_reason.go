package entity

import (
	"time"
)

type OmsOrderReturnReason struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	Name       string    `json:"name" xorm:"default 'NULL' comment('退货类型') VARCHAR(100) 'name'"`
	Sort       int       `json:"sort" xorm:"default NULL INT(11) 'sort'"`
	Status     int       `json:"status" xorm:"default NULL comment('状态：0->不启用；1->启用') INT(1) 'status'"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('添加时间') DATETIME 'create_time'"`
}
