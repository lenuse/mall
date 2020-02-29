package model

import "github.com/jinzhu/gorm"

type GoodsBrand struct {
	gorm.Model
	BrandName  string `json:"brand_name"`
	BrandUrl   string `json:"brand_url"`
	BrandDesc  string `json:"brand_desc"`
	BrandLogo  string `json:"brand_logo"`
	BrandAlias string `json:"brand_alias"`
	IsAble     bool   `json:"-"`
	OrderNum   uint8  `json:"order_num"`
}
