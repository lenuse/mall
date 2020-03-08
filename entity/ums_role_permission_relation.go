package entity

type UmsRolePermissionRelation struct {
	Id           int64 `json:"id" db:"id"`
	RoleId       int64 `json:"role_id" db:"role_id"`
	PermissionId int64 `json:"permission_id" db:"permission_id"`
}

func (m *UmsRolePermissionRelation) TableName() string {
	return "ums_role_permission_relation"
}
