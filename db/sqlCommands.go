package db

import (
	"fmt"
)

// SelectAllFromTable select * from
func SelectAllFromTable(tableName string) []map[string]interface{} {
	c := []map[string]interface{}{}
	sqlCommand := fmt.Sprintf("select * from %s", tableName)
	DB.Raw(sqlCommand).Scan(&c)
	return c
}
