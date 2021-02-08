package initdb

import (
	"log"
	"fmt"
	"gitlab.com/farqlanma_nishani/dbexport/configs"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

// DB is instance of *gorm.DB
var DB *gorm.DB
var err error

func init() {
	if configs.DbConfig.DbDriver == "mysql"{
		DB, err = gorm.Open(mysql.Open(configs.DbConfig.DSN), &gorm.Config{})

	}else if configs.DbConfig.DbDriver == "postgresql"{
		fmt.Println("Postgresql initialize")
		fmt.Println("config",configs.DbConfig)
		DB, err = gorm.Open(postgres.Open(configs.DbConfig.DSN), &gorm.Config{})

	}

	if err != nil {
		log.Println("err: ", err)
	}
}
