package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lenuse/mall/config"
	"github.com/lenuse/mall/handles/admin"
	"github.com/lenuse/mall/middlewares"
	"github.com/lenuse/mall/repository"
)

func main() {
	repository.Init()
	defer repository.Close()

	app := gin.Default()
	Router(app)

	app.Run(config.New().Server.Port)
}

func Router(engine *gin.Engine) {
	engine.Use(middlewares.SetTraceId())

	engine.POST("/v1/admin/login", admin.SignIn)
	engine.POST("/v1/admin/create", admin.CreateAdminUser)
	v1 := engine.Group("/v1/admin").Use(middlewares.JwtVerify())
	{
		v1.GET("/static")
	}
	return
}
