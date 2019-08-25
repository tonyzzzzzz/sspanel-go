package middlewares

import (
	"github.com/gin-gonic/gin"
)

func GuestMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
