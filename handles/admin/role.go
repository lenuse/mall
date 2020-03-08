package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/lenuse/mall/service"
	"github.com/lenuse/mall/transport"
	"github.com/lenuse/mall/utils"
)

func CreateRole(ctx *gin.Context) {
	var err error
	var argument transport.AdminRoleCreate
	err = ctx.ShouldBindJSON(&argument)
	if err != nil {
		utils.NewRespJSON(utils.InvalidArgument).SetMessage(err.Error()).WriteJson(ctx)
		return
	}
	err = service.NewRole(argument)
	if err != nil {
		utils.NewRespJSON(utils.InsertError).SetMessage(err.Error()).WriteJson(ctx)
		return
	}
	utils.NewRespJSON(utils.Success).WriteJson(ctx)
	return
}

func RoleList(ctx *gin.Context) {

}
