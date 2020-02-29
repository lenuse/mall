package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Goods struct {
	gorm.Model
	Sku  uint	`json:"sku"`
	Name string `json:"name"`
	Title string `json:"title"`
	Price string `json:"price"`
	Const string `json:"-"`
	StockNum uint `json:"stock_num"`
	CatID uint `json:"cat_id"`
	BrandID uint `json:"brand_id"`
	Marketable bool `json:"marketable"`
	UpTime time.Time `json:"up_time"`
	DownTime time.Time `json:"down_time"`
}
