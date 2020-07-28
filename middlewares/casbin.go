package middlewares

import (
	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lenuse/mall/utils"
)

func Casbin(e casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.URL.RequestURI()
		act := c.Request.Method
		claims,ok := c.Get(utils.JWTClaims)
		if ok {
			utils.NewResponse(utils.Unauthorized)
			c.Abort()
		}
		sub := claims.(jwt.StandardClaims).Issuer

		e.Enforce(sub, uri, act)
	}
}
