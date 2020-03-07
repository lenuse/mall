package entity

import (
	"database/sql"
	"time"
)

type UmsAdmin struct {
	ID         int64        `json:"id" db:"id"`
	Username   string       `json:"username" db:"username"`
	SecretCode string       `json:"secret_code" db:"secret_code"`
	Icon       string       `json:"icon" db:"icon"`
	Email      string       `json:"email" db:"email"`
	NickName   string       `json:"nick_name" db:"nick_name"`
	Note       string       `json:"note" db:"note"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	LoginAt    time.Time    `json:"login_at" db:"login_at"`
	Status     int          `json:"status" db:"status"`
	UpdatedAt  sql.NullTime `json:"updated_at" db:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

func (m *UmsAdmin) TableName() string {
	return "ums_admin"
}
