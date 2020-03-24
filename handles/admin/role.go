package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/lenuse/mall/repository"
	"github.com/lenuse/mall/service"
	"github.com/lenuse/mall/transport"
	"github.com/lenuse/mall/utils"
)

// CreateRole 创建规则名
func CreateRole(ctx *gin.Context) {
	var err error
	var argument transport.AdminRole
	err = ctx.ShouldBindJSON(&argument)
	if err != nil {
		utils.NewResponse(utils.InvalidArgument).SetMessage(err.Error()).WriteJson(ctx)
		return
	}
	err = service.CreateRole(argument)
	if err != nil {
		utils.NewResponse(utils.InsertError).SetMessage(err.Error()).WriteJson(ctx)
		return
	}
	utils.NewResponse(utils.Success).WriteJson(ctx)
	return
}

// RoleList 规格列表
func RoleList(ctx *gin.Context) {
	roleName := ctx.DefaultQuery("name", "")
	roleList, err := repository.GetRoleRepositoryList(roleName)
	if err != nil {
		utils.NewResponse(utils.QueryFail).SetMessage(err.Error()).WriteJson(ctx)
		return
	}
	utils.NewResponse(utils.Success).SetData(roleList).WriteJson(ctx)
}

func UpdateRoleStatus(ctx *gin.Context) {
	var err error
	var argument transport.AdminRoleStatusUpdate
	err = ctx.ShouldBindJSON(&argument)
	if err != nil {
		utils.NewResponse(utils.InvalidArgument).SetMessage(err.Error()).WriteJson(ctx)
		return
	}

}
