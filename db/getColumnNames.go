package db

import (
	"fmt"
)

// GetColumnNames function returns Column Names
func GetColumnNames(table string) []string {

	columns := []map[string]interface{}{}
	columnNames := []string{}

	sqlCommand := fmt.Sprintf("desc %s", table)

	DB.Raw(sqlCommand).Scan(&columns)

	for _, v := range columns {
		columnNames = append(columnNames, v["Field"].(string))
	}

	return columnNames
}
