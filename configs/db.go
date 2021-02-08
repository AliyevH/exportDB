package configs

import (
	"fmt"
	"os"
)

type config struct {
	DSN      string
	Debug    bool
	Host     string
	Port     string
	DbDriver string
}

// DbConfig object is used as configuration for all go project
var DbConfig config

// DevConfig is used to configure development environment configurations
func (*config) devConfig() {
	DbConfig.DSN = "root:admin@tcp(127.0.0.1:3306)/fn?charset=utf8mb4&parseTime=True&loc=Local"
	DbConfig.Debug = true
	DbConfig.DbDriver = "mysql"
	DbConfig.Host = "localhost"
	DbConfig.Port = "8000"
}

// ProdConfig is used to configure production environment configurations
func (*config) prodConfig() {
	DbDriver := os.Getenv("DB_DRIVER")
	DbUser := os.Getenv("DB_USER")
	DbPass := os.Getenv("DB_PASS")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbName := os.Getenv("DB_NAME")
	DbConfig.Host = "localhost"
	DbConfig.Port = "8000"

	DbConfig.DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DbHost,
		DbUser,
		DbPass,
		DbName,
		DbPort,
	)
	DbConfig.DbDriver = DbDriver
	DbConfig.Debug = false
}

func init() {
	if s := os.Getenv("settings"); s == "prod" {
		DbConfig.prodConfig()
	} else {
		DbConfig.devConfig()
	}
}
