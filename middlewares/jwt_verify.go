package middlewares

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/lenuse/mall/utils"

	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func JwtVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := request.ParseFromRequest(ctx.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil //TODO: 修改secret获取
		})
		if err != nil {
			utils.NewRespJSON(ctx, utils.Unauthorized, nil, err.Error())
			return
		}
		claims, ok := token.Claims.(jwt.StandardClaims)
		if !token.Valid || !ok {
			utils.NewRespJSON(ctx, utils.Unauthorized, nil, "")
			return
		}
		ctx.Set(utils.UerIdKey, claims.Id)
		ctx.Next()
	}
}
