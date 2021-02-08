package sqlcommands

import (
	"fmt"

	initdb "gitlab.com/farqlanma_nishani/dbexport/initDB"
)

// GetColumnNames function returns Column Names
func GetColumnNames(table string) []string {

	columns := []map[string]interface{}{}
	columnNames := []string{}

	sqlCommand := fmt.Sprintf("desc %s", table)

	initdb.DB.Raw(sqlCommand).Scan(&columns)

	for _, v := range columns {
		columnNames = append(columnNames, v["Field"].(string))
	}

	return columnNames
}
