package configs

import (
	"fmt"
	"os"
)

type config struct {
	DSN   string
	DB    string
	Debug bool
	Host  string
	Port  string
}

// DbConfig object is used as configuration for all go project
var DbConfig config

// DevConfig is used to configure development environment configurations
func (*config) devConfig() {
	// ?charset=utf8mb4&parseTime=True&loc=Local"
	DbConfig.DSN = "root:admin@tcp(127.0.0.1:3306)/fn"
	DbConfig.Debug = true
	DbConfig.DB = "mysql"
	DbConfig.Host = "localhost"
	DbConfig.Port = "8000"
}

// ProdConfig is used to configure production environment configurations
func (*config) prodConfig() {
	DbConfig.DB = os.Getenv("DB_DRIVER")
	DbUser := os.Getenv("DB_USER")
	DbPass := os.Getenv("DB_PASS")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_NAME")

	DbConfig.DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbUser,
		DbPass,
		DbHost,
		DbPort,
		DbName,
	)
	DbConfig.Debug = false
}

func init() {
	if s := os.Getenv("settings"); s == "prod" {
		DbConfig.prodConfig()
	} else {
		DbConfig.devConfig()
	}
}
