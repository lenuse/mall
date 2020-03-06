package entity

import (
	"time"
)

type OmsOrderOperateHistory struct {
	Id          int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	OrderId     int64     `json:"order_id" xorm:"default NULL comment('订单id') BIGINT(20) 'order_id'"`
	OperateMan  string    `json:"operate_man" xorm:"default 'NULL' comment('操作人：用户；系统；后台管理员') VARCHAR(100) 'operate_man'"`
	CreateTime  time.Time `json:"create_time" xorm:"default 'NULL' comment('操作时间') DATETIME 'create_time'"`
	OrderStatus int       `json:"order_status" xorm:"default NULL comment('订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单') INT(1) 'order_status'"`
	Note        string    `json:"note" xorm:"default 'NULL' comment('备注') VARCHAR(500) 'note'"`
}
