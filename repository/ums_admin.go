package repository

import "github.com/lenuse/mall/entity"

type AdminRepository struct {
	entity.UmsAdmin
}

func (r *AdminRepository) Save() error {
	newId, err := engine.Collection(r.UmsAdmin.TableName()).Insert(r.UmsAdmin)
	if err != nil {
		return err
	}
	r.ID = newId
	return nil
}
