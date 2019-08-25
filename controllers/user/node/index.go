package node

import (
	"net/http"

	"github.com/tonyzzzzzz/sspanel-go/models"
	"github.com/tonyzzzzzz/sspanel-go/utils"

	"github.com/gin-gonic/gin"
)

// GetInfo returns the info of all nodes
func GetInfo(c *gin.Context) {
	db := utils.GetMySQLInstance().Database
	var rawNodes []models.SsNode
	user, _ := c.Get("user")
	db.Where("type = 1").Where("node_class <= ?", user.(*models.User).Class).Where("sort != 9").Order("name").Find(&rawNodes)

	var nodes []models.Node
	for _, rawNode := range rawNodes {
		nodes = append(nodes, rawNode.GetNode())
	}
	c.JSON(http.StatusOK, nodes)
}
