package admin

import (
	"basic_mall/customerror"
	"basic_mall/services"
	"basic_mall/transports"
	"basic_mall/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

func create(ctx *gin.Context)  {
	var transport transports.CreateShopTransport
	if err := ctx.ShouldBind(transport); err != nil {
		utils.AbortJSON(ctx, customerror.New(4000001, errors.New("传输错误")))
	}
	shopInfo := services.ShopRegister(transport)
	utils.ReturnJSON(ctx, shopInfo)
}

func update(ctx *gin.Context)  {
	shopId := ctx.Param("id")
	var transport transports.UpdateShopTransport

}
