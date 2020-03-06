package entity

import (
	"time"
)

type UmsRole struct {
	Id          int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	Name        string    `json:"name" xorm:"default 'NULL' comment('名称') VARCHAR(100) 'name'"`
	Description string    `json:"description" xorm:"default 'NULL' comment('描述') VARCHAR(500) 'description'"`
	AdminCount  int       `json:"admin_count" xorm:"default NULL comment('后台用户数量') INT(11) 'admin_count'"`
	CreateTime  time.Time `json:"create_time" xorm:"default 'NULL' comment('创建时间') DATETIME 'create_time'"`
	Status      int       `json:"status" xorm:"default 1 comment('启用状态：0->禁用；1->启用') INT(1) 'status'"`
	Sort        int       `json:"sort" xorm:"default 0 INT(11) 'sort'"`
}
