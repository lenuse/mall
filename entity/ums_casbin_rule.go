package entity

type CasbinRule struct {
	ID    uint   `db:"id"`
	PType string `db:"p_type"`
	V0    string `db:"v0"`
	V1    string `db:"v1"`
	V2    string `db:"v2"`
	V3    string `db:"v3"`
	V4    string `db:"v4"`
	V5    string `db:"v5"`
}

func (m *CasbinRule) TableName() string {
	return "ums_casbin_rule"
}