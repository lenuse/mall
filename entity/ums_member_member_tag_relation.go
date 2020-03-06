package entity

type UmsMemberMemberTagRelation struct {
	Id       int64 `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	MemberId int64 `json:"member_id" xorm:"default NULL BIGINT(20) 'member_id'"`
	TagId    int64 `json:"tag_id" xorm:"default NULL BIGINT(20) 'tag_id'"`
}
