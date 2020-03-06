package entity

type UmsMemberRuleSetting struct {
	Id                int64  `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	ContinueSignDay   int    `json:"continue_sign_day" xorm:"default NULL comment('连续签到天数') INT(11) 'continue_sign_day'"`
	ContinueSignPoint int    `json:"continue_sign_point" xorm:"default NULL comment('连续签到赠送数量') INT(11) 'continue_sign_point'"`
	ConsumePerPoint   string `json:"consume_per_point" xorm:"default NULL comment('每消费多少元获取1个点') DECIMAL(10,2) 'consume_per_point'"`
	LowOrderAmount    string `json:"low_order_amount" xorm:"default NULL comment('最低获取点数的订单金额') DECIMAL(10,2) 'low_order_amount'"`
	MaxPointPerOrder  int    `json:"max_point_per_order" xorm:"default NULL comment('每笔订单最高获取点数') INT(11) 'max_point_per_order'"`
	Type              int    `json:"type" xorm:"default NULL comment('类型：0->积分规则；1->成长值规则') INT(1) 'type'"`
}
