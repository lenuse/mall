package entity

import (
	"time"
)

type UmsAdmin struct {
	ID        int64     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password'`
	Icon      string    `json:"icon" db:"icon"`
	Email     string    `json:"email" db:"email"`
	NickName  string    `json:"nick_name" db:"nick_name"`
	Note      string    `json:"note" db:"note"`
	CreatedAt time.Time `json:"create_time" db:"create_time"`
	LoginAt   time.Time `json:"login_time" db:"login_time"`
	Status    int       `json:"status" db:"status"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
}

func (m *UmsAdmin) TableName() string {
	return "ums_admin"
}

// GetAdminByUsername 通过用户名获取用户
func GetAdminByUsername(username string) (admin UmsAdmin) {
	engine.Where("username", username).Get(&admin)
	return
}

// VerifyAdminUsernameAndEmailUnique 验证用户名和电子邮箱是否唯一
func VerifyAdminUsernameAndEmailUnique(username, email string) bool {
	var admin UmsAdmin
	engine.Where("username", username).Or("email", email).Get(&admin)
	return admin.ID == 0
}
