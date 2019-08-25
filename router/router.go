package router

import (
	"net/http"

	"github.com/dvwright/xss-mw"
	"github.com/gin-gonic/gin"
	"github.com/tonyzzzzzz/sspanel-go/controllers/auth"
	"github.com/tonyzzzzzz/sspanel-go/controllers/user"
	"github.com/tonyzzzzzz/sspanel-go/controllers/user/node"
	"github.com/tonyzzzzzz/sspanel-go/controllers/user/shop"
	"github.com/tonyzzzzzz/sspanel-go/middlewares"
)

// GetRouter generates the router of the app
func GetRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	var xssMdlwr xss.XssMw
	r.Use(xssMdlwr.RemoveXss())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "Blinkload API V2"})
	})

	v2 := r.Group("/")

	Auth := v2.Group("/auth")
	Auth.Use(middlewares.GuestMiddleWare())
	{
		Auth.POST("/login", auth.Login)
		Auth.POST("/register")
		Auth.GET("/email_verification")
		Auth.POST("/forget")

	}

	User := v2.Group("/user")
	User.Use(middlewares.AuthMiddleWare())
	{
		// General Info Query
		User.GET("/info", user.GetInfo)
		User.GET("/shop/info", shop.GetInfo)

		User.GET("/node/info", node.GetInfo)

		User.POST("/shop/buy")

		// Billing Interface
		User.GET("/invoice/info")

	}

	Admin := v2.Group("/admin")
	Admin.Use()
	{
		Admin.GET("/info")
		Admin.GET("/incidents")

	}
	return r
}
