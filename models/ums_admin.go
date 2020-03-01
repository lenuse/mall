package models

import "time"

// UmsAdmin 后台用户表映射
type UmsAdmin struct {
	ID          int64
	Username    string
	Password    string
	Icon        string
	Email       string
	NickName    string
	Note        string
	CreatedTime time.Time
	LoginTime   time.Time
	Status      int8
}

// GetAdminByUsername 通过用户名获取用户
func GetAdminByUsername(username string) (admin UmsAdmin) {
	engine.Where("username", username).Get(&admin)
	return
}