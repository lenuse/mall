package mall

import (
	"github.com/gin-gonic/gin"
	"github.com/lenuse/mall/middlewares"
)

func router(engine *gin.Engine)  {
	engine.Use(middlewares.SetTraceId())

	engine.POST("/v1/admin/login")
	v1 := engine.Group("/v1/admin").Use(middlewares.JwtVerify())
	{
		v1.GET("/static")
	}
}
