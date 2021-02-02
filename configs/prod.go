package configs

import (
	"fmt"
	"os"
	"reflect"
)

type config struct {
	DSN      string
	Debug    bool
	Port     string `default:"3306"`
	Host     string `default:"127.0.0.1"`
	Username string `default:"root"`
	Password string `default:"admin"`
	DB       string
	Auth     string
}

func (c *config) initDefaults() {
	if port := os.Getenv("PORT"); port != "" {
		Config.Port = port
	} else {
		port, _ := reflect.TypeOf(Config).FieldByName("Port")

		Config.Port = port.Tag.Get("default")
	}

	if host := os.Getenv("HOST"); host != "" {
		Config.Host = host
	} else {
		host, _ := reflect.TypeOf(Config).FieldByName("Host")
		Config.Host = host.Tag.Get("default")
	}

}

// Config for project
var Config config

// DevConfig Development configuration
func DevConfig() {
	Config.DSN = "root:admin@tcp(127.0.0.1:3306)/fn?charset=utf8mb4&parseTime=True&loc=Local"
	// Config.DSN = "root:admin@tcp(127.0.0.1:3306)/fn"
	Config.Debug = true
	Config.DB = "mysql"

}

// ProdConfig Production configuration
func ProdConfig() {
	Config.DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=enable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	Config.Debug = false
}

// Init Config
func init() {
	if s := os.Getenv("settings"); s == "prod" {
		ProdConfig()
	} else {
		DevConfig()
		Config.initDefaults()
	}
}
