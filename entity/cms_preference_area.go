package entity

type CmsPreferenceArea struct {
	Id         int64  `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	Name       string `json:"name" xorm:"default 'NULL' VARCHAR(255) 'name'"`
	SubTitle   string `json:"sub_title" xorm:"default 'NULL' VARCHAR(255) 'sub_title'"`
	Pic        []byte `json:"pic" xorm:"default NULL comment('展示图片') VARBINARY(500) 'pic'"`
	Sort       int    `json:"sort" xorm:"default NULL INT(11) 'sort'"`
	ShowStatus int    `json:"show_status" xorm:"default NULL INT(1) 'show_status'"`
}
