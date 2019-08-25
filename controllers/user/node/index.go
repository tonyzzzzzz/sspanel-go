package node

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetInfo returns the info of all nodes
func GetInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "OK"})
}
