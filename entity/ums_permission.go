package entity

import (
	"time"
)

type UmsPermission struct {
	Id        int64     `json:"id" db:"id"`
	Pid       int64     `json:"pid" db:"pid"`
	Name      string    `json:"name" db:"name"`
	Value     string    `json:"value" db:"value"`
	Icon      string    `json:"icon" db:"icon"`
	Type      int       `json:"type" db:"type"`
	Uri       string    `json:"uri" db:"uri"`
	Status    int       `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Sort      int       `json:"sort" db:"sort"`
}

func (m *UmsPermission) TableName() string {
	return "ums_role_permission_relation"
}
