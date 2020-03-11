package repository

import (
	"time"

	"github.com/lenuse/mall/entity"
	"upper.io/db.v3"
)

type PermissionRepository struct {
	entity.UmsPermission
	Roles RoleRepositoryList
}

type PermissionRepositoryList []PermissionRepository

func (r *PermissionRepository) Init() error {
	roles, err := getRolesByPermissionId(r.Id)
	if err != nil {
		return err
	}
	r.Roles = roles
	return nil
}

func (r *PermissionRepository) Save() error {
	r.CreatedAt = time.Now()
	newId, err := engine.Collection(r.TableName()).Insert(r.UmsPermission)
	if err != nil {
		return err
	}
	r.Id = newId.(int64)
	return nil
}

func GetPermissionsByRoleId(roleId int64) (PermissionRepositoryList, error) {
	return getPermissionsByRoleId(roleId)
}

func getPermissionsByRoleId(roleId int64) (PermissionRepositoryList, error) {
	var relation entity.UmsRolePermissionRelation
	var permission entity.UmsPermission
	var permissions PermissionRepositoryList
	err := engine.Select(db.Raw("p.*")).From(relation.TableName()+" as rpr").
		LeftJoin(permission.TableName()+" as p").On("rpr.permission_id = p.id").
		Where("rpr.role_id = ?", roleId).
		All(&permissions)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func getPermissions() (PermissionRepositoryList, error) {
	var permissions PermissionRepositoryList
	var permission entity.UmsPermission
	err := engine.Collection(permission.TableName()).
		Find("status", EnableStatus.Int()).
		OrderBy(SortDesc).All(&permissions)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}
