package entity

import (
	"time"
)

type UmsAdmin struct {
	ID         int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	Username   string    `json:"username" xorm:"default 'NULL' VARCHAR(64) 'username'"`
	Password   string    `json:"password" xorm:"default 'NULL' VARCHAR(64) 'password'"`
	Icon       string    `json:"icon" xorm:"default 'NULL' comment('头像') VARCHAR(500) 'icon'"`
	Email      string    `json:"email" xorm:"default 'NULL' comment('邮箱') VARCHAR(100) 'email'"`
	NickName   string    `json:"nick_name" xorm:"default 'NULL' comment('昵称') VARCHAR(200) 'nick_name'"`
	Note       string    `json:"note" xorm:"default 'NULL' comment('备注信息') VARCHAR(500) 'note'"`
	CreateTime time.Time `json:"create_time" xorm:"default 'NULL' comment('创建时间') DATETIME 'create_time'"`
	LoginTime  time.Time `json:"login_time" xorm:"default 'NULL' comment('最后登录时间') DATETIME 'login_time'"`
	Status     int       `json:"status" xorm:"default 1 comment('帐号启用状态：0->禁用；1->启用') INT(1) 'status'"`
}

func (m *UmsAdmin) Save() error {
	_, err := engine.InsertOne(m)
	return err
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
