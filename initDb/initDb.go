package initdb

import (
	"github.com/AliyevH/exportDB/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is instance of *gorm.DB
var DB *gorm.DB

// Init Config
func init() {

	db, err := gorm.Open(mysql.Open(configs.DbConfig.DSN), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	DB = db
}
