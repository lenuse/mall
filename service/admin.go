package service

import (
	"fmt"

	"github.com/lenuse/mall/entity"
	"github.com/lenuse/mall/repository"
	"github.com/lenuse/mall/transport"
	"github.com/lenuse/mall/utils"
)

func NewAdminUser(argument transport.AdminCreate) error {
	password, err := utils.GetBcryptHash(argument.Password)
	if err != nil {
		return err
	}
	admin := repository.AdminRepository{
		UmsAdmin: entity.UmsAdmin{
			Email:      argument.Email,
			SecretCode: password,
			Username:   argument.Username,
			Note:       argument.Note,
			Status:     repository.EnableStatus.Int(),
		},
	}
	fmt.Println(admin)
	return admin.Save()
}
