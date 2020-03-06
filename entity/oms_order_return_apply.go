package entity

import (
	"time"
)

type OmsOrderReturnApply struct {
	Id               int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	OrderId          int64     `json:"order_id" xorm:"default NULL comment('订单id') BIGINT(20) 'order_id'"`
	CompanyAddressId int64     `json:"company_address_id" xorm:"default NULL comment('收货地址表id') BIGINT(20) 'company_address_id'"`
	ProductId        int64     `json:"product_id" xorm:"default NULL comment('退货商品id') BIGINT(20) 'product_id'"`
	OrderSn          string    `json:"order_sn" xorm:"default 'NULL' comment('订单编号') VARCHAR(64) 'order_sn'"`
	CreateTime       time.Time `json:"create_time" xorm:"default 'NULL' comment('申请时间') DATETIME 'create_time'"`
	MemberUsername   string    `json:"member_username" xorm:"default 'NULL' comment('会员用户名') VARCHAR(64) 'member_username'"`
	ReturnAmount     string    `json:"return_amount" xorm:"default NULL comment('退款金额') DECIMAL(10,2) 'return_amount'"`
	ReturnName       string    `json:"return_name" xorm:"default 'NULL' comment('退货人姓名') VARCHAR(100) 'return_name'"`
	ReturnPhone      string    `json:"return_phone" xorm:"default 'NULL' comment('退货人电话') VARCHAR(100) 'return_phone'"`
	Status           int       `json:"status" xorm:"default NULL comment('申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝') INT(1) 'status'"`
	HandleTime       time.Time `json:"handle_time" xorm:"default 'NULL' comment('处理时间') DATETIME 'handle_time'"`
	ProductPic       string    `json:"product_pic" xorm:"default 'NULL' comment('商品图片') VARCHAR(500) 'product_pic'"`
	ProductName      string    `json:"product_name" xorm:"default 'NULL' comment('商品名称') VARCHAR(200) 'product_name'"`
	ProductBrand     string    `json:"product_brand" xorm:"default 'NULL' comment('商品品牌') VARCHAR(200) 'product_brand'"`
	ProductAttr      string    `json:"product_attr" xorm:"default 'NULL' comment('商品销售属性：颜色：红色；尺码：xl;') VARCHAR(500) 'product_attr'"`
	ProductCount     int       `json:"product_count" xorm:"default NULL comment('退货数量') INT(11) 'product_count'"`
	ProductPrice     string    `json:"product_price" xorm:"default NULL comment('商品单价') DECIMAL(10,2) 'product_price'"`
	ProductRealPrice string    `json:"product_real_price" xorm:"default NULL comment('商品实际支付单价') DECIMAL(10,2) 'product_real_price'"`
	Reason           string    `json:"reason" xorm:"default 'NULL' comment('原因') VARCHAR(200) 'reason'"`
	Description      string    `json:"description" xorm:"default 'NULL' comment('描述') VARCHAR(500) 'description'"`
	ProofPics        string    `json:"proof_pics" xorm:"default 'NULL' comment('凭证图片，以逗号隔开') VARCHAR(1000) 'proof_pics'"`
	HandleNote       string    `json:"handle_note" xorm:"default 'NULL' comment('处理备注') VARCHAR(500) 'handle_note'"`
	HandleMan        string    `json:"handle_man" xorm:"default 'NULL' comment('处理人员') VARCHAR(100) 'handle_man'"`
	ReceiveMan       string    `json:"receive_man" xorm:"default 'NULL' comment('收货人') VARCHAR(100) 'receive_man'"`
	ReceiveTime      time.Time `json:"receive_time" xorm:"default 'NULL' comment('收货时间') DATETIME 'receive_time'"`
	ReceiveNote      string    `json:"receive_note" xorm:"default 'NULL' comment('收货备注') VARCHAR(500) 'receive_note'"`
}
