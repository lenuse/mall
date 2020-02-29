package model

import (
	"basic_mall/transports"
	"github.com/jinzhu/gorm"
	"time"
)

type Shop struct {
	gorm.Model
	ShopName  string
	Address  string
	Province  int
	City   int
	District  int
	Street   int
	RealName   string
	IDCard   string
	State   int
	Longitude   string
	Latitude   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt time.Time
}

func (shopInfo Shop) GetTableName() string {
	return "shop"
}


const (
	ShopStateCreate int = iota
)


func ShopCreate(data interface{}) (Shop, error) {
	info := data.(transports.CreateShopTransport)
	shopInfo := Shop{
		ShopName:  info.ShopName,
		Address :  info.Address,
		Province : info.Province,
		City :     info.City,
		District : info.District,
		Street :   info.Street,
		State:     ShopStateCreate,
	}
	err := db.Create(&shopInfo).Error
	return shopInfo, err
}

func ShopUpdate(shopInfo Shop, data map[string]interface{}) (Shop, error) {
	err := db.Model(&shopInfo).Updates(data).Error
	return shopInfo, err
}

func ShopDelete(id uint)  {
	var shopInfo Shop
	db.Table(shopInfo.GetTableName()).Where("id = ?", id).Update(map[string]interface{}{"deleted_at": time.Now()})
}