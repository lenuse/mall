package entity

import (
	"time"
)

type UmsMember struct {
	Id                    int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	MemberLevelId         int64     `json:"member_level_id" xorm:"default NULL BIGINT(20) 'member_level_id'"`
	Username              string    `json:"username" xorm:"default 'NULL' comment('用户名') VARCHAR(64) 'username'"`
	Password              string    `json:"password" xorm:"default 'NULL' comment('密码') VARCHAR(64) 'password'"`
	Nickname              string    `json:"nickname" xorm:"default 'NULL' comment('昵称') VARCHAR(64) 'nickname'"`
	Phone                 string    `json:"phone" xorm:"default 'NULL' comment('手机号码') VARCHAR(64) 'phone'"`
	Status                int       `json:"status" xorm:"default NULL comment('帐号启用状态:0->禁用；1->启用') INT(1) 'status'"`
	CreateTime            time.Time `json:"create_time" xorm:"default 'NULL' comment('注册时间') DATETIME 'create_time'"`
	Icon                  string    `json:"icon" xorm:"default 'NULL' comment('头像') VARCHAR(500) 'icon'"`
	Gender                int       `json:"gender" xorm:"default NULL comment('性别：0->未知；1->男；2->女') INT(1) 'gender'"`
	Birthday              time.Time `json:"birthday" xorm:"default 'NULL' comment('生日') DATE 'birthday'"`
	City                  string    `json:"city" xorm:"default 'NULL' comment('所做城市') VARCHAR(64) 'city'"`
	Job                   string    `json:"job" xorm:"default 'NULL' comment('职业') VARCHAR(100) 'job'"`
	PersonalizedSignature string    `json:"personalized_signature" xorm:"default 'NULL' comment('个性签名') VARCHAR(200) 'personalized_signature'"`
	SourceType            int       `json:"source_type" xorm:"default NULL comment('用户来源') INT(1) 'source_type'"`
	Integration           int       `json:"integration" xorm:"default NULL comment('积分') INT(11) 'integration'"`
	Growth                int       `json:"growth" xorm:"default NULL comment('成长值') INT(11) 'growth'"`
	LuckeyCount           int       `json:"luckey_count" xorm:"default NULL comment('剩余抽奖次数') INT(11) 'luckey_count'"`
	HistoryIntegration    int       `json:"history_integration" xorm:"default NULL comment('历史积分数量') INT(11) 'history_integration'"`
}
