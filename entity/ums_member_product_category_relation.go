package entity

type UmsMemberProductCategoryRelation struct {
	Id                int64 `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	MemberId          int64 `json:"member_id" xorm:"default NULL BIGINT(20) 'member_id'"`
	ProductCategoryId int64 `json:"product_category_id" xorm:"default NULL BIGINT(20) 'product_category_id'"`
}
