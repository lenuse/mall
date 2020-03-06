package entity

type UmsRolePermissionRelation struct {
	Id           int64 `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	RoleId       int64 `json:"role_id" xorm:"default NULL BIGINT(20) 'role_id'"`
	PermissionId int64 `json:"permission_id" xorm:"default NULL BIGINT(20) 'permission_id'"`
}
