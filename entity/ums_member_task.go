package entity

type UmsMemberTask struct {
	Id           int64  `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	Name         string `json:"name" xorm:"default 'NULL' VARCHAR(100) 'name'"`
	Growth       int    `json:"growth" xorm:"default NULL comment('赠送成长值') INT(11) 'growth'"`
	Intergration int    `json:"intergration" xorm:"default NULL comment('赠送积分') INT(11) 'intergration'"`
	Type         int    `json:"type" xorm:"default NULL comment('任务类型：0->新手任务；1->日常任务') INT(1) 'type'"`
}
