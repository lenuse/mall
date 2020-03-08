package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/lenuse/mall/repository"
	"github.com/lenuse/mall/service"
	"github.com/lenuse/mall/transport"
	"github.com/lenuse/mall/utils"
)

// SignIn 后台登陆
func SignIn(ctx *gin.Context) {
	var argument transport.AdminSignIn
	if err := ctx.ShouldBindJSON(&argument); err != nil {
		utils.NewRespJSON(utils.InvalidArgument).SetMessage(err.Error()).WriteJson(ctx)
		return
	}
	admin, _ := repository.GetAdminByUsername(argument.Username)
	if !utils.VerifyBcryptHash(admin.SecretCode, argument.Password) {
		utils.NewRespJSON(utils.Unauthorized).WriteJson(ctx)
		return
	}
	token, err := utils.GetJwtToken(string(admin.ID), admin.NickName, "admin")
	if err != nil {
		utils.NewRespJSON(utils.JwtError).SetMessage(err.Error()).WriteJson(ctx)
		return
	}
	utils.NewRespJSON(utils.JwtError).SetData(map[string]string{"token": token}).WriteJson(ctx)
	return
}

func CreateAdminUser(ctx *gin.Context) {
	var argument transport.AdminCreate
	if err := ctx.ShouldBindJSON(&argument); err != nil {
		utils.NewRespJSON(utils.InvalidArgument).SetMessage(err.Error()).WriteJson(ctx)
		return
	}
	isUnique := repository.VerifyAdminUsernameAndEmailUnique(argument.Username, argument.Email)
	if !isUnique {
		utils.NewRespJSON(utils.AdminNotUnique).WriteJson(ctx)
		return
	}
	err := service.NewAdminUser(argument)
	if err != nil {
		utils.NewRespJSON(utils.AdminNotUnique).SetMessage(err.Error()).WriteJson(ctx)
		return
	}
	utils.NewRespJSON(utils.Success).WriteJson(ctx)
	return
}
