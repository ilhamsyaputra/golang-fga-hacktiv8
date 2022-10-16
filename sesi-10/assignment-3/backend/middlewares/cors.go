package middlewares

import (
	"github.com/gin-gonic/gin"
)

func EnableCORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Next()
	}
}
