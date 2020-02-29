package middlewares

import (
	"basic_mall/conf"
	cerr "basic_mall/customerror"
	"basic_mall/model"
	"basic_mall/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		var tokenStr string
		tokenStr = context.DefaultQuery("token", "")
		if tokenStr == "" {
			bearerTokenStr := context.GetHeader("Authorization")
			if bearerTokenStr == "" {
				utils.AbortJSON(context, cerr.Unauthorized)
				return
			}
			tokenStr = strings.Replace(bearerTokenStr, "Bearer ", "", -1)
		}
		token, err := jwt.ParseWithClaims(tokenStr, &model.UserClaims{}, func(token *jwt.Token) (i interface{}, e error) {
			return conf.GetJwtSecret(), nil
		})
		if err != nil {
			utils.AbortJSON(context, cerr.Custom(cerr.Unauthorized.GetCode(), err))
		}
		if claims, ok := token.Claims.(*model.UserClaims); ok && token.Valid {
			context.Set("user_id", claims.ID)
		} else {
			utils.AbortJSON(context, cerr.Custom(cerr.Unauthorized.GetCode(), err))
			return
		}
		context.Next()
	}
}
