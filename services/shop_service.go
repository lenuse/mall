package services

import (
	"mall/model"
	"mall/transports"
)

func ShopRegister(transport transports.CreateShopTransport) model.Shop {
	return model.ShopCreate(transport)
}

func ShopUpdate(id int, transport transports.CreateShopTransport) {

}
