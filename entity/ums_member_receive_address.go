package entity

type UmsMemberReceiveAddress struct {
	Id            int64  `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	MemberId      int64  `json:"member_id" xorm:"default NULL BIGINT(20) 'member_id'"`
	Name          string `json:"name" xorm:"default 'NULL' comment('收货人名称') VARCHAR(100) 'name'"`
	PhoneNumber   string `json:"phone_number" xorm:"default 'NULL' VARCHAR(64) 'phone_number'"`
	DefaultStatus int    `json:"default_status" xorm:"default NULL comment('是否为默认') INT(1) 'default_status'"`
	PostCode      string `json:"post_code" xorm:"default 'NULL' comment('邮政编码') VARCHAR(100) 'post_code'"`
	Province      string `json:"province" xorm:"default 'NULL' comment('省份/直辖市') VARCHAR(100) 'province'"`
	City          string `json:"city" xorm:"default 'NULL' comment('城市') VARCHAR(100) 'city'"`
	Region        string `json:"region" xorm:"default 'NULL' comment('区') VARCHAR(100) 'region'"`
	DetailAddress string `json:"detail_address" xorm:"default 'NULL' comment('详细地址(街道)') VARCHAR(128) 'detail_address'"`
}
