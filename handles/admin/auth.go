package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/lenuse/mall/models"
	"github.com/lenuse/mall/transport"
	"github.com/lenuse/mall/utils"
)

// SignIn 后台登陆
func SignIn(ctx *gin.Context) {
	var argument transport.AdminSignIn
	if err := ctx.ShouldBindJSON(&argument); err != nil {
		utils.NewRespJSON(ctx, utils.StateCodeInvalidArgument, nil, err.Error())
		return
	}
	admin := models.GetAdminByUsername(argument.Username)
	if !utils.VerifyBcryptHash(admin.Password, argument.Password) {
		utils.NewRespJSON(ctx, utils.StateCodeUnauthorized, nil, "")
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

func Create(ctx *gin.Context) {
	var argument transport.AdminCreate
	if err := ctx.ShouldBindJSON(&argument); err != nil {
		utils.NewRespJSON(ctx, utils.StateCodeInvalidArgument, nil, err.Error())
		return
	}
	if argument.RepeatInput != argument.Password {
		return utils.NewRespJSON(ctx, utils.StateCodeInvalidArgument, nil, err.Error())
	}
	models.UmsAdmin{
		Email: argument.Email,
	}

}
