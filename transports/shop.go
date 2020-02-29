package transports

type CreateShopTransport struct {
	ShopName string `json:"shop_name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Province  int `json:"province" binding:"required"`
	City   int `json:"city" binding:"required"`
	District  int `json:"district" binding:"required"`
	Street   int `json:"street" binding:"required"`
	RealName   string `json:"real_name"`
	IDCard   string `json:"id_card"`
	Longitude   string `json:"longitude"`
	Latitude   string `json:"latitude"`
}

type UpdateShopTransport struct {
	ShopName string `json:"shop_name"`
	Address  string `json:"address"`
	Province  int `json:"province"`
	City   int `json:"city"`
	District  int `json:"district"`
	Street   int `json:"street"`
	RealName   string `json:"real_name"`
	IDCard   string `json:"id_card"`
	Longitude   string `json:"longitude"`
	Latitude   string `json:"latitude"`
}