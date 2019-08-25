package utils

import (
	"log"
	"sync"

	"github.com/jinzhu/gorm"
)

// SingletonMySQL constructs the Database Connector
type SingletonMySQL struct {
	Database *gorm.DB
}

var instance *SingletonMySQL
var lock sync.Once

// GetMySQLInstance returns a MySQL Instance
// Written By Indexyz @ Cloudhammer/Seeds
func GetMySQLInstance() *SingletonMySQL {
	lock.Do(func() {
		// Init MySQL Instance
		db, err := gorm.Open("mysql", GetDataURL())
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		db.LogMode(GetConfig().GetBool("enableLogger"))
		//db.AutoMigrate(&models.SsNode{})
		//db.AutoMigrate(&models.User{})

		instance = &SingletonMySQL{
			Database: db,
		}
	})
	return instance
}
