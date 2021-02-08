package main

import (
	"github.com/AliyevH/exportDB/configs"
	"github.com/AliyevH/exportDB/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/exportExcel", routers.ExportToExcel)

	r.Run(configs.DbConfig.Host + ":" + configs.DbConfig.Port)

}
