package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/lenuse/mall/entity"
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
	admin := entity.GetAdminByUsername(argument.Username)
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
	if !entity.VerifyAdminUsernameAndEmailUnique(argument.Username, argument.Email) {
		utils.NewRespJSON(ctx, utils.StateCodeAdminNotUnique, nil, "")
		return
	}

	admin := entity.UmsAdmin{
		Email:    argument.Email,
		Password: argument.Password,
		Username: argument.Username,
		Note:     argument.Note,
		Status:   entity.EnableStatus.Int8(),
	}
	err := admin.Save()
	if err != nil {
		utils.NewRespJSON(ctx, utils.StateCodeInsertError, nil, "")
		return
	}
	utils.NewRespJSON(ctx, utils.StateCodeSuccess, nil, "")
	return
}
