package middlewares

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/lenuse/mall/utils"
)

func JwtVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := request.ParseFromRequest(ctx.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil //TODO: 修改secret获取
		})
		if err != nil {
			utils.NewRespJSON(utils.Unauthorized).WriteJson(ctx)
			return
		}
		claims, ok := token.Claims.(jwt.StandardClaims)
		if !token.Valid || !ok {
			utils.NewRespJSON(utils.Unauthorized).WriteJson(ctx)
			return
		}
		ctx.Set(utils.UerIdKey, claims.Id)
		ctx.Next()
	}
}
