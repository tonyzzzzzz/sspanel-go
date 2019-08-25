package node

import (
	"net/http"

	"github.com/fatih/structs"

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

	var nodes []map[string]interface{}
	for _, rawNode := range rawNodes {
		m := structs.Map(rawNode)
		m["Online"] = rawNode.GetOnlineUserCount()
	}
	c.JSON(http.StatusOK, nodes)
}
