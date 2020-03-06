package entity

import (
	"time"
)

type OmsOrder struct {
	Id                    int64     `json:"id" xorm:"pk autoincr comment('订单id') BIGINT(20) 'id'"`
	MemberId              int64     `json:"member_id" xorm:"not null BIGINT(20) 'member_id'"`
	CouponId              int64     `json:"coupon_id" xorm:"default NULL BIGINT(20) 'coupon_id'"`
	OrderSn               string    `json:"order_sn" xorm:"default 'NULL' comment('订单编号') VARCHAR(64) 'order_sn'"`
	CreateTime            time.Time `json:"create_time" xorm:"default 'NULL' comment('提交时间') DATETIME 'create_time'"`
	MemberUsername        string    `json:"member_username" xorm:"default 'NULL' comment('用户帐号') VARCHAR(64) 'member_username'"`
	TotalAmount           string    `json:"total_amount" xorm:"default NULL comment('订单总金额') DECIMAL(10,2) 'total_amount'"`
	PayAmount             string    `json:"pay_amount" xorm:"default NULL comment('应付金额（实际支付金额）') DECIMAL(10,2) 'pay_amount'"`
	FreightAmount         string    `json:"freight_amount" xorm:"default NULL comment('运费金额') DECIMAL(10,2) 'freight_amount'"`
	PromotionAmount       string    `json:"promotion_amount" xorm:"default NULL comment('促销优化金额（促销价、满减、阶梯价）') DECIMAL(10,2) 'promotion_amount'"`
	IntegrationAmount     string    `json:"integration_amount" xorm:"default NULL comment('积分抵扣金额') DECIMAL(10,2) 'integration_amount'"`
	CouponAmount          string    `json:"coupon_amount" xorm:"default NULL comment('优惠券抵扣金额') DECIMAL(10,2) 'coupon_amount'"`
	DiscountAmount        string    `json:"discount_amount" xorm:"default NULL comment('管理员后台调整订单使用的折扣金额') DECIMAL(10,2) 'discount_amount'"`
	PayType               int       `json:"pay_type" xorm:"default NULL comment('支付方式：0->未支付；1->支付宝；2->微信') INT(1) 'pay_type'"`
	SourceType            int       `json:"source_type" xorm:"default NULL comment('订单来源：0->PC订单；1->app订单') INT(1) 'source_type'"`
	Status                int       `json:"status" xorm:"default NULL comment('订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单') INT(1) 'status'"`
	OrderType             int       `json:"order_type" xorm:"default NULL comment('订单类型：0->正常订单；1->秒杀订单') INT(1) 'order_type'"`
	DeliveryCompany       string    `json:"delivery_company" xorm:"default 'NULL' comment('物流公司(配送方式)') VARCHAR(64) 'delivery_company'"`
	DeliverySn            string    `json:"delivery_sn" xorm:"default 'NULL' comment('物流单号') VARCHAR(64) 'delivery_sn'"`
	AutoConfirmDay        int       `json:"auto_confirm_day" xorm:"default NULL comment('自动确认时间（天）') INT(11) 'auto_confirm_day'"`
	Integration           int       `json:"integration" xorm:"default NULL comment('可以获得的积分') INT(11) 'integration'"`
	Growth                int       `json:"growth" xorm:"default NULL comment('可以活动的成长值') INT(11) 'growth'"`
	PromotionInfo         string    `json:"promotion_info" xorm:"default 'NULL' comment('活动信息') VARCHAR(100) 'promotion_info'"`
	BillType              int       `json:"bill_type" xorm:"default NULL comment('发票类型：0->不开发票；1->电子发票；2->纸质发票') INT(1) 'bill_type'"`
	BillHeader            string    `json:"bill_header" xorm:"default 'NULL' comment('发票抬头') VARCHAR(200) 'bill_header'"`
	BillContent           string    `json:"bill_content" xorm:"default 'NULL' comment('发票内容') VARCHAR(200) 'bill_content'"`
	BillReceiverPhone     string    `json:"bill_receiver_phone" xorm:"default 'NULL' comment('收票人电话') VARCHAR(32) 'bill_receiver_phone'"`
	BillReceiverEmail     string    `json:"bill_receiver_email" xorm:"default 'NULL' comment('收票人邮箱') VARCHAR(64) 'bill_receiver_email'"`
	ReceiverName          string    `json:"receiver_name" xorm:"not null comment('收货人姓名') VARCHAR(100) 'receiver_name'"`
	ReceiverPhone         string    `json:"receiver_phone" xorm:"not null comment('收货人电话') VARCHAR(32) 'receiver_phone'"`
	ReceiverPostCode      string    `json:"receiver_post_code" xorm:"default 'NULL' comment('收货人邮编') VARCHAR(32) 'receiver_post_code'"`
	ReceiverProvince      string    `json:"receiver_province" xorm:"default 'NULL' comment('省份/直辖市') VARCHAR(32) 'receiver_province'"`
	ReceiverCity          string    `json:"receiver_city" xorm:"default 'NULL' comment('城市') VARCHAR(32) 'receiver_city'"`
	ReceiverRegion        string    `json:"receiver_region" xorm:"default 'NULL' comment('区') VARCHAR(32) 'receiver_region'"`
	ReceiverDetailAddress string    `json:"receiver_detail_address" xorm:"default 'NULL' comment('详细地址') VARCHAR(200) 'receiver_detail_address'"`
	Note                  string    `json:"note" xorm:"default 'NULL' comment('订单备注') VARCHAR(500) 'note'"`
	ConfirmStatus         int       `json:"confirm_status" xorm:"default NULL comment('确认收货状态：0->未确认；1->已确认') INT(1) 'confirm_status'"`
	DeleteStatus          int       `json:"delete_status" xorm:"not null default 0 comment('删除状态：0->未删除；1->已删除') INT(1) 'delete_status'"`
	UseIntegration        int       `json:"use_integration" xorm:"default NULL comment('下单时使用的积分') INT(11) 'use_integration'"`
	PaymentTime           time.Time `json:"payment_time" xorm:"default 'NULL' comment('支付时间') DATETIME 'payment_time'"`
	DeliveryTime          time.Time `json:"delivery_time" xorm:"default 'NULL' comment('发货时间') DATETIME 'delivery_time'"`
	ReceiveTime           time.Time `json:"receive_time" xorm:"default 'NULL' comment('确认收货时间') DATETIME 'receive_time'"`
	CommentTime           time.Time `json:"comment_time" xorm:"default 'NULL' comment('评价时间') DATETIME 'comment_time'"`
	ModifyTime            time.Time `json:"modify_time" xorm:"default 'NULL' comment('修改时间') DATETIME 'modify_time'"`
}
