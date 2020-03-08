package entity

import (
	"time"
)

type UmsRole struct {
	Id          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	AdminCount  int       `json:"admin_count" db:"admin_count"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	Status      int       `json:"status" db:"status"`
	Sort        int       `json:"sort" db:"sort"`
}

func (m *UmsRole) TableName() string {
	return "ums_role"
}
