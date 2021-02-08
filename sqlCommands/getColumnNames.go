package sqlcommands

import (
	"fmt"

	"github.com/AliyevH/exportDB/configs"
	"github.com/AliyevH/exportDB/initdb"
)

// GetColumnNames function returns Column Names
func GetColumnNames(table string) []string {
	var columns []string
	if configs.DbConfig.DbDriver == "mysql" {
		columns = getColumnsMySQL(table)
	} else if configs.DbConfig.DbDriver == "postgresql" {
		columns = getColumnsPG(table)
	}
	return columns
}

// GetColumnNamesPG function returns Column Names
func getColumnsPG(table string) []string {

	columns := []map[string]interface{}{}
	columnNames := []string{}

	sqlCommand := fmt.Sprintf("SELECT column_name FROM information_schema.columns WHERE table_name='%s'", "pipeline_steps")

	initdb.DB.Raw(sqlCommand).Scan(&columns)

	for _, v := range columns {
		columnNames = append(columnNames, v["column_name"].(string))
	}

	return columnNames
}

func getColumnsMySQL(table string) []string {
	columns := []map[string]interface{}{}
	columnNames := []string{}

	sqlCommand := fmt.Sprintf("desc %s", table)

	initdb.DB.Raw(sqlCommand).Scan(&columns)

	for _, v := range columns {
		columnNames = append(columnNames, v["Field"].(string))
	}

	return columnNames
}
