package model

type District struct {
	ID         uint   `gorm:"primary_key" json:"-"`
	Code       int    `json:"code"`
	Name       string `gorm:"size:50" json:"name"`
	ParentCode int    `json:"parent_code"`
	Level      int    `json:"level"`
}

func GetDistrictProvinces() []District {
	var provinces []District
	db.Where("level = ?", 1).Find(&provinces)
	return provinces
}

func GetDistrictCities(provinceCode int) []District {
	var cities []District
	db.Where("parent_code = ?", provinceCode).Find(&cities)
	return cities
}

func GetDistrictAreas(cityCode int) []District {
	var areas []District
	db.Where("parent_code = ?", cityCode).Find(&areas)
	return areas
}

func GetDistrictStreets(cityCode int) []District {
	var streets []District
	db.Where("parent_code = ?", cityCode).Find(&streets)
	return streets
}

func GetDistrictCodes(codes []string) []District {
	var districts []District
	db.Where("parent_code in (?)", codes).First(&districts)
	return districts
}
