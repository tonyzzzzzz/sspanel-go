package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/tonyzzzzzz/sspanel-go/router"
)

func main() {

	r := router.GetRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")

}
