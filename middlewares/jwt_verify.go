package middlewares

import (
	"mall/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func JwtVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := request.ParseFromRequest(ctx.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			utils.NewResqJson(ctx, utils.StateCodeUnauthorized, nil, err.Error())
		}
		claims, ok := token.Claims.(jwt.StandardClaims)
		if !token.Valid || !ok {
			utils.NewResqJson(ctx, utils.StateCodeUnauthorized, nil, "")
		}
		ctx.Set("userId", claims.Id)
		ctx.Next()
	}
}
