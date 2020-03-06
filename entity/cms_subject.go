package entity

import (
	"time"
)

type CmsSubject struct {
	Id              int64     `json:"id" xorm:"pk autoincr BIGINT(20) 'id'"`
	CategoryId      int64     `json:"category_id" xorm:"not null default 0 BIGINT(20) 'category_id'"`
	Title           string    `json:"title" xorm:"not null default '' VARCHAR(100) 'title'"`
	Pic             string    `json:"pic" xorm:"not null default '' comment('专题主图') VARCHAR(500) 'pic'"`
	ProductCount    int       `json:"product_count" xorm:"default NULL comment('关联产品数量') INT(11) 'product_count'"`
	RecommendStatus int       `json:"recommend_status" xorm:"default NULL INT(1) 'recommend_status'"`
	CreateTime      time.Time `json:"create_time" xorm:"default 'NULL' DATETIME 'create_time'"`
	CollectCount    int       `json:"collect_count" xorm:"default NULL INT(11) 'collect_count'"`
	ReadCount       int       `json:"read_count" xorm:"default NULL INT(11) 'read_count'"`
	CommentCount    int       `json:"comment_count" xorm:"default NULL INT(11) 'comment_count'"`
	AlbumPics       string    `json:"album_pics" xorm:"default 'NULL' comment('画册图片用逗号分割') VARCHAR(1000) 'album_pics'"`
	Description     string    `json:"description" xorm:"default 'NULL' VARCHAR(1000) 'description'"`
	ShowStatus      int       `json:"show_status" xorm:"default NULL comment('显示状态：0->不显示；1->显示') INT(1) 'show_status'"`
	Content         string    `json:"content" xorm:"TEXT 'content'"`
	ForwardCount    int       `json:"forward_count" xorm:"default NULL comment('转发数') INT(11) 'forward_count'"`
	CategoryName    string    `json:"category_name" xorm:"default 'NULL' comment('专题分类名称') VARCHAR(200) 'category_name'"`
}
