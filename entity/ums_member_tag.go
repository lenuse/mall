package entity

type UmsMemberTag struct {
	Id                int64  `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	Name              string `json:"name" xorm:"default 'NULL' VARCHAR(100) 'name'"`
	FinishOrderCount  int    `json:"finish_order_count" xorm:"default NULL comment('自动打标签完成订单数量') INT(11) 'finish_order_count'"`
	FinishOrderAmount string `json:"finish_order_amount" xorm:"default NULL comment('自动打标签完成订单金额') DECIMAL(10,2) 'finish_order_amount'"`
}
