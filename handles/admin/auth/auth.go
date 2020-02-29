package auth

import (
	"basic_mall/customerror"
	"basic_mall/model"
	"basic_mall/utils"

	"github.com/gin-gonic/gin"
)

// AdminLoginValidation admin登录验证
type AdminLoginValidation struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login admin登录接口
func Login(c *gin.Context) {
	var loginParam AdminLoginValidation
	if err := c.ShouldBind(&loginParam); err != nil {
		utils.AbortJSON(c, customerror.Custom(402, err))
		return
	}
	admin := model.GetAdminByUsername(loginParam.Username)
	reqPassHash, _ := utils.PasswordHash([]byte(loginParam.Password))
	result := utils.PasswordVerify([]byte(admin.Password), []byte(reqPassHash))
	if result {

		token, _ := model.BuildJwtToken(model.AdminClaims{ID: admin.ID, Name: admin.Name})
		utils.ReturnJSON(c, map[string]string{
			"token": token,
		})
		return
	}
	utils.AbortJSON(c, customerror.New(500, "密码错误"))
	return
}
