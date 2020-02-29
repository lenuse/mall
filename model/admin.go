package model

import (
	"github.com/jinzhu/gorm"
)

type Admin struct {
	gorm.Model
	Name     string `grom:"size:25"`
	Username string `grom:"size:50"`
	Password string `grom:"size:50"`
}

func GetAdminByUsername(name string) *Admin {
	emptyModel := new(Admin)
	db.Where("name = ?", name).First(emptyModel)
	return emptyModel
}

func GetAdminById(id uint64) *Admin {
	emptyModel := new(Admin)
	db.First(emptyModel, id)
	return emptyModel
}
