package routers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/AliyevH/exportDB/export"
	"github.com/AliyevH/exportDB/initdb"
	"github.com/AliyevH/exportDB/schemas"
)

// ExportToExcel function to export
func ExportToExcel(c *gin.Context) {
	var wg sync.WaitGroup

	var data schemas.JSONData

	err := c.ShouldBindJSON(&data)

	if err != nil {
		log.Println("err:", err)
	}

	result := []map[string]interface{}{}

	initdb.DB.Table(data.TableName).Where(data.Condition, data.Values...).Scan(&result)

	wg.Add(1)
	export.ToExcel(data.TableName, data.Columns, &result, &wg)

	c.JSON(http.StatusOK, result)
	wg.Wait()
}
