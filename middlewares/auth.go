package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tonyzzzzzz/sspanel-go/models"
	"github.com/tonyzzzzzz/sspanel-go/utils"
)

// AuthMiddleWare identifies the user and proceed to authorized zones
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", 1)
		claims, err := utils.ParseToken(token)

		if token == "" || err != nil || claims == nil {
			c.Request.Header.Set("Authorization", "")
			c.JSON(http.StatusForbidden, gin.H{"msg": "Token Invalid"})
			c.Abort()
			return
		}

		var user models.User
		db := utils.GetMySQLInstance().Database
		db.First(&user, claims.UserID)

		c.Set("user", &user)
		c.Next()

	}
}
