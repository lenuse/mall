package mall

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lenuse/mall/handles/admin"
	"github.com/lenuse/mall/middlewares"
	"net/http"
	"time"
)

func New() *gin.Engine {
	app := gin.Default()
	app.Use(func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 8*time.Second)
		defer func() {
			if c.Err() == context.DeadlineExceeded {
				c.Writer.WriteHeader(http.StatusRequestTimeout)
				c.Abort()
			}
			cancel()
		}()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	})
	app.Use(middlewares.SetTraceId())

	app.POST("/v1/admin/login", admin.SignIn)
	app.POST("/v1/admin/create", admin.CreateAdminUser)
	v1 := app.Group("/v1/admin").Use(middlewares.JwtVerify())
	{
		v1.GET("/static")
	}
	return app
}
