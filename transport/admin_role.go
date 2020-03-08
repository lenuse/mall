package transport

import "github.com/lenuse/mall/repository"

type AdminRoleCreate struct {
	Name        string            `form:"name" json:"name" binding:"required"`
	Description string            `form:"description" json:"description"`
	Status      repository.Status `form:"description" json:"description" bind:"required"`
	Sort        int               `form:"sort" json:"sort"`
}

type AdminRoleStatusUpdate struct {
	Ids    string            `form:"sort" json:"sort" bind:"required"`
	Status repository.Status `form:"sort" json:"sort" bind:"required"`
}
