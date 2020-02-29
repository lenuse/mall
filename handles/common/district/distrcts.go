package district

import (
	"basic_mall/customerror"
	"basic_mall/model"
	"basic_mall/until"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//GetProvinces 获取省级列表
func GetProvinces(c *gin.Context) {
	provinces := model.GetDistrictProvinces()
	until.ReturnJson(c, provinces)
	return
}

//GetCities 获取城市列表
func GetCities(c *gin.Context) {
	provinceCode := c.Param("province_code")
	code, err := strconv.Atoi(provinceCode)
	if err != nil {
		until.AbortJson(c, customerror.Custom(customerror.ParameterError.GetCode(), err))
	}
	cities := model.GetDistrictCities(code)
	until.ReturnJson(c, cities)
	return
}

//GetAreas 获取区县
func GetAreas(c *gin.Context) {
	cityCode := c.Param("city_code")
	code, err := strconv.Atoi(cityCode)
	if err != nil {
		until.AbortJson(c, customerror.Custom(customerror.ParameterError.GetCode(), err))
	}
	areas := model.GetDistrictAreas(code)
	until.ReturnJson(c, areas)
	return
}

//GetStreets 获取乡镇
func GetStreets(c *gin.Context) {
	areaCode := c.Param("area_code")
	code, err := strconv.Atoi(areaCode)
	if err != nil {
		until.AbortJson(c, customerror.Custom(customerror.ParameterError.GetCode(), err))
	}
	streets := model.GetDistrictStreets(code)
	until.ReturnJson(c, streets)
	return
}

//GetDistricts 获取指定id的城市
func GetDistricts(c *gin.Context) {
	codes := c.DefaultQuery("codes", "11")
	codeSlice := strings.Split(",", codes)
	fmt.Println(codeSlice)
	districts := model.GetDistrictCodes(codeSlice)
	until.ReturnJson(c, districts)
	return
}
