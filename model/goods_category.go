package model

import "github.com/jinzhu/gorm"

type GoodsCategory struct {
	gorm.Model
	ParentID   uint   `json:"parent_id"`
	CatName    string `json:"cat_name"`
	IsLeaf     bool   `json:"is_leaf"`
	Disable    bool   `json:"-"`
	GoodsCount uint   `json:"goods_count"`
	ChildCount uint   `json:"child_count"`
}
