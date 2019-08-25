package shop

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonyzzzzzz/sspanel-go/models"
	"github.com/tonyzzzzzz/sspanel-go/utils"
)

// GetInfo fetches the shop list and returns to client
func GetInfo(c *gin.Context) {
	var shops []models.Shop
	db := utils.GetMySQLInstance().Database
	db.Where("status = 1").Find(&shops)

	var respShops []map[string]interface{}
	for _, shop := range shops {
		respShops = append(respShops, map[string]interface{}{
			"Id":      shop.Id,
			"Name":    shop.Name,
			"Price":   shop.Price,
			"Content": shop.GetContent(),
		})
	}
	c.JSON(http.StatusOK, respShops)
}
