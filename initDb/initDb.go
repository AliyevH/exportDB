package initdb

import (
	"log"

	"gitlab.com/farqlanma_nishani/dbexport/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is instance of *gorm.DB
var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open(mysql.Open(configs.DbConfig.DSN), &gorm.Config{})

	if err != nil {
		log.Println("err: ", err)
	}
}
