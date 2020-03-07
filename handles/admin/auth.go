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
		utils.NewRespJSON(ctx, utils.InvalidArgument, nil, err.Error())
		return
	}
	admin, _ := repository.GetAdminByUsername(argument.Username)
	if !utils.VerifyBcryptHash(admin.SecretCode, argument.Password) {
		utils.NewRespJSON(ctx, utils.Unauthorized, nil, "")
		return
	}
	token, err := utils.GetJwtToken(string(admin.ID), admin.NickName, "admin")
	if err != nil {
		utils.NewRespJSON(ctx, utils.JwtError, nil, err.Error())
		return
	}
	utils.NewRespJSON(ctx, utils.Success, map[string]string{"token": token}, "")
	return
}

func CreateAdminUser(ctx *gin.Context) {
	var argument transport.AdminCreate
	if err := ctx.ShouldBindJSON(&argument); err != nil {
		utils.NewRespJSON(ctx, utils.InvalidArgument, nil, err.Error())
		return
	}
	isUnique := repository.VerifyAdminUsernameAndEmailUnique(argument.Username, argument.Email)
	if !isUnique {
		utils.NewRespJSON(ctx, utils.AdminNotUnique, nil, "")
		return
	}
	err := service.NewAdminUser(argument)
	if err != nil {
		utils.NewRespJSON(ctx, utils.InsertError, nil, err.Error())
		return
	}
	utils.NewRespJSON(ctx, utils.Success, nil, "")
	return
}
