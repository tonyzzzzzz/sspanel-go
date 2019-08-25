package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

// GetConfig returns the global config
func GetConfig() *viper.Viper {
	c := viper.New()
	c.SetConfigType("yaml")
	c.SetConfigName("config")
	c.AddConfigPath(".")
	c.AutomaticEnv()

	c.SetDefault("security.encryption", "sha256") // Now supports sha256 & md5, will support more

	c.SetDefault("database.host", "localhost")
	c.SetDefault("database.port", 3306)
	c.SetDefault("database.name", "sspanel")
	c.SetDefault("database.user", "sspanel")
	c.SetDefault("database.pass", "sspanel")

	c.SetDefault("redis.host", "localhost")
	c.SetDefault("redis.port", 4468)
	c.SetDefault("redis.enableAuth", true)
	c.SetDefault("redis.password", "password")

	c.SetDefault("server.address", ":8080")
	c.SetDefault("server.debug", false)

	c.SetDefault("cacheTTL", 60)
	c.SetDefault("verifyKey", "Hello")
	c.SetDefault("modURL", true)

	replacer := strings.NewReplacer(".", "_")
	c.SetEnvKeyReplacer(replacer)
	c.ReadInConfig()
	return c
}

// GetDataURL returns the database url for gorm connection
func GetDataURL() string {
	config := GetConfig()
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
		config.Get("database.user"), config.Get("database.pass"), config.Get("database.host"),
		strconv.Itoa(config.Get("database.port").(int)), config.Get("database.name"))
}
