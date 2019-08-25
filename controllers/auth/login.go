package auth

import (
	"fmt"
	"net/http"

	"github.com/tonyzzzzzz/sspanel-go/models"

	"github.com/gin-gonic/gin"
	"github.com/tonyzzzzzz/sspanel-go/utils"
)

// Credentials as JSON binding
type Credentials struct {
	UserName   string `form:"username" json:"username" xml:"username" binding:"required"`
	Password   string `form:"password" json:"password" xml:"password" binding:"required"`
	RememberMe bool   `form:"remember" json:"remember" xml:"remember"`
}

// Login handles user login request
func Login(c *gin.Context) {
	var credentials Credentials
	if err := c.ShouldBind(&credentials); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db := utils.GetMySQLInstance()
	var user models.User
	db.Database.Where("user_name = ? OR email = ?", credentials.UserName, credentials.UserName).First(&user)
	fmt.Println(user)
	if hashed := utils.EncryptString(credentials.Password); hashed != user.Pass {
		c.JSON(http.StatusForbidden, gin.H{"error": "Wrong Username or Password"})
		return
	}

	token, err := utils.GenerateToken(credentials.UserName, credentials.Password, user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
