package entity

type CmsSubjectProductRelation struct {
	Id        int64 `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	SubjectId int64 `json:"subject_id" xorm:"default NULL BIGINT(20) 'subject_id'"`
	ProductId int64 `json:"product_id" xorm:"default NULL BIGINT(20) 'product_id'"`
}
