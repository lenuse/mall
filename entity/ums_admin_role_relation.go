package entity

type UmsAdminRoleRelation struct {
	Id      int64 `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	AdminId int64 `json:"admin_id" xorm:"default NULL BIGINT(20) 'admin_id'"`
	RoleId  int64 `json:"role_id" xorm:"default NULL BIGINT(20) 'role_id'"`
}
