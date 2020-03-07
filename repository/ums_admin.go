package repository

import (
	"github.com/lenuse/mall/entity"
	"upper.io/db.v3"
)

type AdminRepository struct {
	entity.UmsAdmin
}

func (r *AdminRepository) Save() error {
	newId, err := engine.Collection(r.TableName()).Insert(r.UmsAdmin)
	if err != nil {
		return err
	}
	r.ID = newId.(int64)
	return nil
}

// GetAdminByUsername 通过用户名获取用户
func GetAdminByUsername(username string) (r AdminRepository, err error) {
	err = engine.Collection(r.TableName()).
		Find(db.Cond{"username": username}).
		One(&r)
	return
}

// VerifyAdminUsernameAndEmailUnique 验证用户名和电子邮箱是否唯一
func VerifyAdminUsernameAndEmailUnique(username, email string) bool {
	var r AdminRepository
	cond := db.Or(
		db.Cond{"username":username},
		db.Cond{"email":email},
		)
	err := engine.Collection(r.TableName()).
		Find(cond).One(&r)
	if err != nil || r.ID != 0 {
		return false
	}
	return true
}
