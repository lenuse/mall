package entity

type OmsCompanyAddress struct {
	Id            int64  `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	AddressName   string `json:"address_name" xorm:"default 'NULL' comment('地址名称') VARCHAR(200) 'address_name'"`
	SendStatus    int    `json:"send_status" xorm:"default NULL comment('默认发货地址：0->否；1->是') INT(1) 'send_status'"`
	ReceiveStatus int    `json:"receive_status" xorm:"default NULL comment('是否默认收货地址：0->否；1->是') INT(1) 'receive_status'"`
	Name          string `json:"name" xorm:"default 'NULL' comment('收发货人姓名') VARCHAR(64) 'name'"`
	Phone         string `json:"phone" xorm:"default 'NULL' comment('收货人电话') VARCHAR(64) 'phone'"`
	Province      string `json:"province" xorm:"default 'NULL' comment('省/直辖市') VARCHAR(64) 'province'"`
	City          string `json:"city" xorm:"default 'NULL' comment('市') VARCHAR(64) 'city'"`
	Region        string `json:"region" xorm:"default 'NULL' comment('区') VARCHAR(64) 'region'"`
	DetailAddress string `json:"detail_address" xorm:"default 'NULL' comment('详细地址') VARCHAR(200) 'detail_address'"`
}
