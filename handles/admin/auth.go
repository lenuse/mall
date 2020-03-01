package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/lenuse/mall/models"
	"github.com/lenuse/mall/transport"
	"github.com/lenuse/mall/utils"
)

func Login(ctx *gin.Context) {
	var argument transport.AdminLogin
	if err := ctx.ShouldBindJSON(&argument); err != nil {
		utils.NewRespJSON(ctx, utils.StateCodeInvalidArgument, nil, err.Error())
		return
	}
	admin := models.GetAdminByUsername(argument.Username)
	if admin.Password != argument.Password {
		return
	}
	token, err := utils.GetJwtToken(string(admin.ID), admin.NickName, "admin")
	if err != nil {
		utils.NewRespJSON(ctx, utils.StateCodeJwtError, nil, err.Error())
		return
	}
	utils.NewRespJSON(ctx, utils.StateCodeSuccess, map[string]string{"token": token}, "")
	return
}
