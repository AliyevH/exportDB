package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/farqlanma_nishani/dbexport/configs"
	"gitlab.com/farqlanma_nishani/dbexport/routers"
)

func main() {
	r := gin.Default()
	r.POST("/exportCsv", routers.ExportToExcel)

	r.Run(configs.DbConfig.Host + ":" + configs.DbConfig.Port)

}
