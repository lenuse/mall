package entity

type CmsPreferenceAreaProductRelation struct {
	Id               int64 `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	PreferenceAreaId int64 `json:"preference_area_id" xorm:"default NULL BIGINT(20) 'preference_area_id'"`
	ProductId        int64 `json:"product_id" xorm:"default NULL BIGINT(20) 'product_id'"`
}
