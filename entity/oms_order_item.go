package entity

type OmsOrderItem struct {
	Id                int64  `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	OrderId           int64  `json:"order_id" xorm:"default NULL comment('订单id') BIGINT(20) 'order_id'"`
	OrderSn           string `json:"order_sn" xorm:"default 'NULL' comment('订单编号') VARCHAR(64) 'order_sn'"`
	ProductId         int64  `json:"product_id" xorm:"default NULL BIGINT(20) 'product_id'"`
	ProductPic        string `json:"product_pic" xorm:"default 'NULL' VARCHAR(500) 'product_pic'"`
	ProductName       string `json:"product_name" xorm:"default 'NULL' VARCHAR(200) 'product_name'"`
	ProductBrand      string `json:"product_brand" xorm:"default 'NULL' VARCHAR(200) 'product_brand'"`
	ProductSn         string `json:"product_sn" xorm:"default 'NULL' VARCHAR(64) 'product_sn'"`
	ProductPrice      string `json:"product_price" xorm:"default NULL comment('销售价格') DECIMAL(10,2) 'product_price'"`
	ProductQuantity   int    `json:"product_quantity" xorm:"default NULL comment('购买数量') INT(11) 'product_quantity'"`
	ProductSkuId      int64  `json:"product_sku_id" xorm:"default NULL comment('商品sku编号') BIGINT(20) 'product_sku_id'"`
	ProductSkuCode    string `json:"product_sku_code" xorm:"default 'NULL' comment('商品sku条码') VARCHAR(50) 'product_sku_code'"`
	ProductCategoryId int64  `json:"product_category_id" xorm:"default NULL comment('商品分类id') BIGINT(20) 'product_category_id'"`
	Sp1               string `json:"sp1" xorm:"default 'NULL' comment('商品的销售属性') VARCHAR(100) 'sp1'"`
	Sp2               string `json:"sp2" xorm:"default 'NULL' VARCHAR(100) 'sp2'"`
	Sp3               string `json:"sp3" xorm:"default 'NULL' VARCHAR(100) 'sp3'"`
	PromotionName     string `json:"promotion_name" xorm:"default 'NULL' comment('商品促销名称') VARCHAR(200) 'promotion_name'"`
	PromotionAmount   string `json:"promotion_amount" xorm:"default NULL comment('商品促销分解金额') DECIMAL(10,2) 'promotion_amount'"`
	CouponAmount      string `json:"coupon_amount" xorm:"default NULL comment('优惠券优惠分解金额') DECIMAL(10,2) 'coupon_amount'"`
	IntegrationAmount string `json:"integration_amount" xorm:"default NULL comment('积分优惠分解金额') DECIMAL(10,2) 'integration_amount'"`
	RealAmount        string `json:"real_amount" xorm:"default NULL comment('该商品经过优惠后的分解金额') DECIMAL(10,2) 'real_amount'"`
	GiftIntegration   int    `json:"gift_integration" xorm:"default 0 INT(11) 'gift_integration'"`
	GiftGrowth        int    `json:"gift_growth" xorm:"default 0 INT(11) 'gift_growth'"`
	ProductAttr       string `json:"product_attr" xorm:"default 'NULL' comment('商品销售属性:[{\"key\":\"颜色\",\"value\":\"颜色\"},{\"key\":\"容量\",\"value\":\"4G\"}]') VARCHAR(500) 'product_attr'"`
}
