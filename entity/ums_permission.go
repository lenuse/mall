package entity

import (
	"time"
)

type UmsPermission struct {
	Id         int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	Pid        int64     `json:"pid" xorm:"default NULL comment('父级权限id') BIGINT(20) 'pid'"`
	Name       string    `json:"name" xorm:"default 'NULL' comment('名称') VARCHAR(100) 'name'"`
	Value      string    `json:"value" xorm:"default 'NULL' comment('权限值') VARCHAR(200) 'value'"`
	Icon       string    `json:"icon" xorm:"default 'NULL' comment('图标') VARCHAR(500) 'icon'"`
	Type       int       `json:"type" xorm:"default NULL comment('权限类型：0->目录；1->菜单；2->按钮（接口绑定权限）') INT(1) 'type'"`
	Uri        string    `json:"uri" xorm:"default 'NULL' comment('前端资源路径') VARCHAR(200) 'uri'"`
	Status     int       `json:"status" xorm:"default NULL comment('启用状态；0->禁用；1->启用') INT(1) 'status'"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('创建时间') DATETIME 'create_time'"`
	Sort       int       `json:"sort" xorm:"default NULL comment('排序') INT(11) 'sort'"`
}
