package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/tonyzzzzzz/sspanel-go/router"
	"github.com/tonyzzzzzz/sspanel-go/utils"
)

func main() {

	r := router.GetRouter()

	r.Run(utils.GetConfig().GetString("server.address"))

}
