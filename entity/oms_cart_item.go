package entity

import (
	"time"
)

type OmsCartItem struct {
	Id                int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	ProductId         int64     `json:"product_id" xorm:"default NULL BIGINT(20) 'product_id'"`
	ProductSkuId      int64     `json:"product_sku_id" xorm:"default NULL BIGINT(20) 'product_sku_id'"`
	MemberId          int64     `json:"member_id" xorm:"default NULL BIGINT(20) 'member_id'"`
	Quantity          int       `json:"quantity" xorm:"default NULL comment('购买数量') INT(11) 'quantity'"`
	Price             string    `json:"price" xorm:"default NULL comment('添加到购物车的价格') DECIMAL(10,2) 'price'"`
	Sp1               string    `json:"sp1" xorm:"default 'NULL' comment('销售属性1') VARCHAR(200) 'sp1'"`
	Sp2               string    `json:"sp2" xorm:"default 'NULL' comment('销售属性2') VARCHAR(200) 'sp2'"`
	Sp3               string    `json:"sp3" xorm:"default 'NULL' comment('销售属性3') VARCHAR(200) 'sp3'"`
	ProductPic        string    `json:"product_pic" xorm:"default 'NULL' comment('商品主图') VARCHAR(1000) 'product_pic'"`
	ProductName       string    `json:"product_name" xorm:"default 'NULL' comment('商品名称') VARCHAR(500) 'product_name'"`
	ProductSubTitle   string    `json:"product_sub_title" xorm:"default 'NULL' comment('商品副标题（卖点）') VARCHAR(500) 'product_sub_title'"`
	ProductSkuCode    string    `json:"product_sku_code" xorm:"default 'NULL' comment('商品sku条码') VARCHAR(200) 'product_sku_code'"`
	MemberNickname    string    `json:"member_nickname" xorm:"default 'NULL' comment('会员昵称') VARCHAR(500) 'member_nickname'"`
	CreateDate        time.Time `json:"create_date" xorm:"default 'NULL' comment('创建时间') DATETIME 'create_date'"`
	ModifyDate        time.Time `json:"modify_date" xorm:"default 'NULL' comment('修改时间') DATETIME 'modify_date'"`
	DeleteStatus      int       `json:"delete_status" xorm:"default 0 comment('是否删除') INT(1) 'delete_status'"`
	ProductCategoryId int64     `json:"product_category_id" xorm:"default NULL comment('商品分类') BIGINT(20) 'product_category_id'"`
	ProductBrand      string    `json:"product_brand" xorm:"default 'NULL' VARCHAR(200) 'product_brand'"`
	ProductSn         string    `json:"product_sn" xorm:"default 'NULL' VARCHAR(200) 'product_sn'"`
	ProductAttr       string    `json:"product_attr" xorm:"default 'NULL' comment('商品销售属性:[{\"key\":\"颜色\",\"value\":\"颜色\"},{\"key\":\"容量\",\"value\":\"4G\"}]') VARCHAR(500) 'product_attr'"`
}
