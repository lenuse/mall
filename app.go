package mall

import (
	"github.com/gin-gonic/gin"
	"github.com/lenuse/mall/handles/admin"
	"github.com/lenuse/mall/middlewares"
)

func New() *gin.Engine {
	app := gin.Default()
	app.Use(middlewares.SetTraceId())

	app.POST("/v1/admin/login", admin.SignIn)
	app.POST("/v1/admin/create", admin.CreateAdminUser)
	v1 := app.Group("/v1/admin").Use(middlewares.JwtVerify())
	{
		v1.GET("/static")
	}
	return app
}
