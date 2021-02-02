package db

import (
	"github.com/AliyevH/exportDB/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// Init Config
func init() {

	DB, err = gorm.Open(mysql.Open(configs.Config.DSN), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
}
