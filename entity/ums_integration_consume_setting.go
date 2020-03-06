package entity

type UmsIntegrationConsumeSetting struct {
	Id                 int64 `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	DeductionPerAmount int   `json:"deduction_per_amount" xorm:"default NULL comment('每一元需要抵扣的积分数量') INT(11) 'deduction_per_amount'"`
	MaxPercentPerOrder int   `json:"max_percent_per_order" xorm:"default NULL comment('每笔订单最高抵用百分比') INT(11) 'max_percent_per_order'"`
	UseUnit            int   `json:"use_unit" xorm:"default NULL comment('每次使用积分最小单位100') INT(11) 'use_unit'"`
	CouponStatus       int   `json:"coupon_status" xorm:"default NULL comment('是否可以和优惠券同用；0->不可以；1->可以') INT(1) 'coupon_status'"`
}
