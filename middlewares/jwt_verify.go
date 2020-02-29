package middlewares

import (
	"mall/conf"
	cerr "mall/customerror"
	"mall/model"
	"mall/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		secret := ctx.GetHeader("Authorization")
		if secret == "" && len(secret) <= 7 {
			utils.NewResqJson(ctx, utils.StateCodeUnauthorized, nil)
		}
		token := secret[7:]
		jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {

		})

	}
}

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
