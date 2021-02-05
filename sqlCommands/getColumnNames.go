package sqlcommands

import (
	"fmt"

	initdb "github.com/AliyevH/exportDB/initDb"
)

// GetColumnNames function returns Column Names
func GetColumnNames(table string) []string {

	db := initdb.DB

	columns := []map[string]interface{}{}
	columnNames := []string{}

	sqlCommand := fmt.Sprintf("desc %s", table)

	db.Raw(sqlCommand).Scan(&columns)

	for _, v := range columns {
		columnNames = append(columnNames, v["Field"].(string))
	}

	return columnNames
}
