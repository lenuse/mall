package transport

import "github.com/lenuse/mall/repository"

type AdminRole struct {
	Name        string            `form:"name" json:"name" binding:"required"`
	Description string            `form:"description" json:"description"`
	Status      repository.Status `form:"description" json:"description" bind:"required,min=1,max=2"`
	Sort        int               `form:"sort" json:"sort"`
}

type AdminRoleStatusUpdate struct {
	Id     []int64           `form:"id" json:"sort" bind:"gte=1,dive,required"`
	Status repository.Status `form:"status" json:"sort" bind:"required,min=1,max=2"`
}
