package entity

type OmsOrderSetting struct {
	Id                  int64 `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	FlashOrderOvertime  int   `json:"flash_order_overtime" xorm:"default NULL comment('秒杀订单超时关闭时间(分)') INT(11) 'flash_order_overtime'"`
	NormalOrderOvertime int   `json:"normal_order_overtime" xorm:"default NULL comment('正常订单超时时间(分)') INT(11) 'normal_order_overtime'"`
	ConfirmOvertime     int   `json:"confirm_overtime" xorm:"default NULL comment('发货后自动确认收货时间（天）') INT(11) 'confirm_overtime'"`
	FinishOvertime      int   `json:"finish_overtime" xorm:"default NULL comment('自动完成交易时间，不能申请售后（天）') INT(11) 'finish_overtime'"`
	CommentOvertime     int   `json:"comment_overtime" xorm:"default NULL comment('订单完成后自动好评时间（天）') INT(11) 'comment_overtime'"`
}
