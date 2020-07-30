package middlewares

import (
	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/lenuse/mall/utils"
)

func AdminVerify(e casbin.IEnforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//签名验证
		token, err := request.ParseFromRequest(ctx.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil //TODO: 修改secret获取
		})
		if err != nil {
			utils.NewResponse(utils.Unauthorized).WriteJson(ctx)
			ctx.Abort()
		}
		claims, ok := token.Claims.(jwt.StandardClaims)
		if !token.Valid || !ok {
			utils.NewResponse(utils.Unauthorized).WriteJson(ctx)
			ctx.Abort()
		}
		ctx.Set(utils.JWTClaims, claims)
		//权限验证
		uri := ctx.Request.URL.RequestURI()
		act := ctx.Request.Method
		sub := claims.Issuer
		ok, err = e.Enforce(sub, uri, act)
		if err != nil || !ok {
			utils.NewResponse(utils.Forbidden).WriteJson(ctx)
			ctx.Abort()
		}
		ctx.Next()
	}
}
