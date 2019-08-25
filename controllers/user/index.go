package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetInfo handles user info request
func GetInfo(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, user)
}
