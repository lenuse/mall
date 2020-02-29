package router

import (
	"basic_mall/handles/admin/auth"
	"basic_mall/handles/common/district"
	"basic_mall/middlewares"

	"github.com/gin-gonic/gin"
)

func New(app *gin.Engine) *gin.Engine {
	adminGroup := app.Group("/admin")
	{
		adminGroup.POST("/v1/login", auth.Login)
		adminV1Group := adminGroup.Group("/v1").Use(middlewares.JwtUser())
		{
			adminV1Group.GET("/provinces", district.GetProvinces)
			adminV1Group.GET("/districts", district.GetDistricts)
			adminV1Group.GET("/cities/:province_code", district.GetCities)
			adminV1Group.GET("/areas/:city_code", district.GetAreas)
			adminV1Group.GET("/streets/:area_code", district.GetStreets)
		}

	}

	// userGroup := app.Group("/user")
	// {
	// 	userGroup.POST("/v1/login")
	// 	userV1Group := userGroup.Group("/v1")
	// 	{
	// 		userV1Group.GET("/provinces", district.GetProvinces)
	// 		userV1Group.GET("/districts", district.GetDistricts)
	// 		userV1Group.GET("/cities/:province_code", district.GetCities)
	// 		userV1Group.GET("/areas/:city_code", district.GetAreas)
	// 		userV1Group.GET("/streets/:area_code", district.GetStreets)
	// 	}
	// }

	return app
}
