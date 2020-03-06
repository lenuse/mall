package entity

type UmsMemberLevel struct {
	Id                    int64  `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	Name                  string `json:"name" xorm:"default 'NULL' VARCHAR(100) 'name'"`
	GrowthPoint           int    `json:"growth_point" xorm:"default NULL INT(11) 'growth_point'"`
	DefaultStatus         int    `json:"default_status" xorm:"default NULL comment('是否为默认等级：0->不是；1->是') INT(1) 'default_status'"`
	FreeFreightPoint      string `json:"free_freight_point" xorm:"default NULL comment('免运费标准') DECIMAL(10,2) 'free_freight_point'"`
	CommentGrowthPoint    int    `json:"comment_growth_point" xorm:"default NULL comment('每次评价获取的成长值') INT(11) 'comment_growth_point'"`
	PriviledgeFreeFreight int    `json:"priviledge_free_freight" xorm:"default NULL comment('是否有免邮特权') INT(1) 'priviledge_free_freight'"`
	PriviledgeSignIn      int    `json:"priviledge_sign_in" xorm:"default NULL comment('是否有签到特权') INT(1) 'priviledge_sign_in'"`
	PriviledgeComment     int    `json:"priviledge_comment" xorm:"default NULL comment('是否有评论获奖励特权') INT(1) 'priviledge_comment'"`
	PriviledgePromotion   int    `json:"priviledge_promotion" xorm:"default NULL comment('是否有专享活动特权') INT(1) 'priviledge_promotion'"`
	PriviledgeMemberPrice int    `json:"priviledge_member_price" xorm:"default NULL comment('是否有会员价格特权') INT(1) 'priviledge_member_price'"`
	PriviledgeBirthday    int    `json:"priviledge_birthday" xorm:"default NULL comment('是否有生日特权') INT(1) 'priviledge_birthday'"`
	Note                  string `json:"note" xorm:"default 'NULL' VARCHAR(200) 'note'"`
}
