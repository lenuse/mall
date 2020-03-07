package service

import (
	"github.com/lenuse/mall/entity"
	"github.com/lenuse/mall/repository"
	"github.com/lenuse/mall/transport"
	"github.com/lenuse/mall/utils"
)

func NewAdminUser(argument transport.AdminCreate) error {
	password, _ := utils.GetBcryptHash(argument.Password)
	admin := repository.AdminRepository{
		UmsAdmin: entity.UmsAdmin{
			Email:    argument.Email,
			Password: password,
			Username: argument.Username,
			Note:     argument.Note,
			Status:   repository.EnableStatus.Int(),
		},
	}
	return admin.Save()
}
