package entity

type CmsHelpCategory struct {
	Id         int64  `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	Name       string `json:"name" xorm:"default 'NULL' VARCHAR(100) 'name'"`
	Icon       string `json:"icon" xorm:"default 'NULL' comment('分类图标') VARCHAR(500) 'icon'"`
	HelpCount  int    `json:"help_count" xorm:"default NULL comment('专题数量') INT(11) 'help_count'"`
	ShowStatus int    `json:"show_status" xorm:"default NULL INT(2) 'show_status'"`
	Sort       int    `json:"sort" xorm:"default NULL INT(11) 'sort'"`
}
