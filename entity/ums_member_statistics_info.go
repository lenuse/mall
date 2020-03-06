package entity

import (
	"time"
)

type UmsMemberStatisticsInfo struct {
	Id                  int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	MemberId            int64     `json:"member_id" xorm:"default NULL BIGINT(20) 'member_id'"`
	ConsumeAmount       string    `json:"consume_amount" xorm:"default NULL comment('累计消费金额') DECIMAL(10,2) 'consume_amount'"`
	OrderCount          int       `json:"order_count" xorm:"default NULL comment('订单数量') INT(11) 'order_count'"`
	CouponCount         int       `json:"coupon_count" xorm:"default NULL comment('优惠券数量') INT(11) 'coupon_count'"`
	CommentCount        int       `json:"comment_count" xorm:"default NULL comment('评价数') INT(11) 'comment_count'"`
	ReturnOrderCount    int       `json:"return_order_count" xorm:"default NULL comment('退货数量') INT(11) 'return_order_count'"`
	LoginCount          int       `json:"login_count" xorm:"default NULL comment('登录次数') INT(11) 'login_count'"`
	AttendCount         int       `json:"attend_count" xorm:"default NULL comment('关注数量') INT(11) 'attend_count'"`
	FansCount           int       `json:"fans_count" xorm:"default NULL comment('粉丝数量') INT(11) 'fans_count'"`
	CollectProductCount int       `json:"collect_product_count" xorm:"default NULL INT(11) 'collect_product_count'"`
	CollectSubjectCount int       `json:"collect_subject_count" xorm:"default NULL INT(11) 'collect_subject_count'"`
	CollectTopicCount   int       `json:"collect_topic_count" xorm:"default NULL INT(11) 'collect_topic_count'"`
	CollectCommentCount int       `json:"collect_comment_count" xorm:"default NULL INT(11) 'collect_comment_count'"`
	InviteFriendCount   int       `json:"invite_friend_count" xorm:"default NULL INT(11) 'invite_friend_count'"`
	RecentOrderTime     time.Time `json:"recent_order_time" xorm:"default 'NULL' comment('最后一次下订单时间') DATETIME 'recent_order_time'"`
}
