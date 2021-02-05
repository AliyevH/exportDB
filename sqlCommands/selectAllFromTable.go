package sqlcommands

import (
	"fmt"

	initdb "github.com/AliyevH/exportDB/initDb"
)

// SelectAllFromTable select * from
func SelectAllFromTable(tableName string) []map[string]interface{} {
	db := initdb.DB

	c := []map[string]interface{}{}
	sqlCommand := fmt.Sprintf("select * from %s", tableName)
	db.Raw(sqlCommand).Scan(&c)
	return c
}
