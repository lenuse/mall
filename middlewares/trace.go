package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lenuse/mall/utils"
)

func SetTraceId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, exists := ctx.Get(utils.TraceIdKey)
		if !exists {
			id := uuid.New().String()
			ctx.Set(utils.TraceIdKey, id)
		}
		ctx.Next()
	}
}