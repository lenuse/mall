package repository

import (
	"time"

	"github.com/lenuse/mall/entity"
	"upper.io/db.v3"
)

type AdminRepository struct {
	entity.UmsAdmin
}

func (r *AdminRepository) Save() error {
	r.CreatedAt = time.Now()
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
		Find(db.Cond{
			"username":   username,
			"deleted_at": db.IsNull,
		}).
		One(&r)
	return
}

// VerifyAdminUsernameAndEmailUnique 验证用户名和电子邮箱是否唯一
func VerifyAdminUsernameAndEmailUnique(username, email string) bool {
	var r AdminRepository
	cond := db.Or(
		db.Cond{"username": username},
		db.Cond{"email": email},
	)
	collection := engine.Collection(r.TableName())
	collection.Find(cond).One(&r)
	if r.ID != 0 {
		return false
	}
	return true
}
