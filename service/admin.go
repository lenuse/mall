package service

import (
	"time"

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
	return admin.Save()
}

func NewRole(argument transport.AdminRoleCreate) error {
	role := repository.RoleRepository{
		UmsRole: entity.UmsRole{
			Name:        argument.Name,
			Description: argument.Description,
			CreatedAt:   time.Now(),
			Status:      argument.Status.Int(),
			Sort:        argument.Sort,
		},
	}
	return role.Save()
}
