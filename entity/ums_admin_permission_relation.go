package entity

type UmsAdminPermissionRelation struct {
	Id           int64 `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	AdminId      int64 `json:"admin_id" xorm:"default NULL BIGINT(20) 'admin_id'"`
	PermissionId int64 `json:"permission_id" xorm:"default NULL BIGINT(20) 'permission_id'"`
	Type         int   `json:"type" xorm:"default NULL INT(1) 'type'"`
}
