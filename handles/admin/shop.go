package admin

import (
	"errors"
	"mall/customerror"
	"mall/services"
	"mall/transports"
	"mall/utils"

	"github.com/gin-gonic/gin"
)

func create(ctx *gin.Context) {
	var transport transports.CreateShopTransport
	if err := ctx.ShouldBind(transport); err != nil {
		utils.AbortJSON(ctx, customerror.New(4000001, errors.New("传输错误")))
	}
	shopInfo := services.ShopRegister(transport)
	utils.ReturnJSON(ctx, shopInfo)
}

func update(ctx *gin.Context) {
	shopId := ctx.Param("id")
	var transport transports.UpdateShopTransport

}
