package repository

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/lenuse/mall/entity"
	"time"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

type RoleRepository struct {
	entity.UmsRole
	Permissions PermissionRepositoryList
}

type RoleRepositoryList []RoleRepository

func (r *RoleRepository) Save() error {
	r.CreatedAt = time.Now()
	newId, err := engine.Collection(r.TableName()).Insert(r.UmsRole)
	if err != nil {
		return err
	}
	r.Id = newId.(int64)
	return nil
}

func (r *RoleRepository) Init() error {
	permissions, err := getPermissionsByRoleId(r.Id)
	if err != nil {
		return err
	}
	r.Permissions = permissions
	return nil
}

func GetRolesByPermissionId(permissionId int64) (RoleRepositoryList, error) {
	return getRolesByPermissionId(permissionId)
}

func getRolesByPermissionId(permissionId int64) (RoleRepositoryList, error) {
	var relation entity.UmsRolePermissionRelation
	var role entity.UmsRole
	var roles RoleRepositoryList
	err := engine.Select(db.Raw("p.*")).From(relation.TableName()+" as rpr").
		LeftJoin(role.TableName()+" as r").On("r.id = rpr.role_id").
		Where("rpr.role_id = ?", permissionId).
		All(&roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func createPermissionRelation(roleId int64, permissionIds []int64) error {
	ctx := context.Background()
	if len(permissionIds) == 0 {
		return errors.New("permissionIds is null")
	}
	var relation entity.UmsRolePermissionRelation
	return engine.Tx(ctx, func(sess sqlbuilder.Tx) (err error) {
		_, err = sess.DeleteFrom(relation.TableName()).Where("role_id", roleId).Exec()
		if err != nil {
			return err
		}
		var bt bytes.Buffer
		sql := fmt.Sprintf("INSERT INTO %s (role_id, permission_id) VALUES", relation.TableName())
		bt.WriteString(sql)
		size := len(permissionIds)
		for i, permissionId := range permissionIds {
			if i+1 == size {
				bt.WriteString(fmt.Sprintf("(%d, %d)", roleId, permissionId))
			} else {
				bt.WriteString(fmt.Sprintf("(%d, %d),", roleId, permissionId))
			}
		}
		_, err = sess.Exec(bt.String())
		return
	})
}

func getRoleRepositoryList(name string) (RoleRepositoryList, error) {
	var role entity.UmsRole
	var roleList RoleRepositoryList
	res := engine.Collection(role.TableName()).Find()
	if name != "" {
		res.Where("name = ?", name)
	}
	err := res.OrderBy("-sort").All(&roleList)
	if err != nil {
		return nil, err
	}
	return roleList, nil
}

func GetRoleRepositoryList(name string) (RoleRepositoryList, error) {
	return getRoleRepositoryList(name)
}
