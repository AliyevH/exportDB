package sqlcommands

import (
	"fmt"

	"github.com/AliyevH/exportDB/initdb"
)

// SelectAllFromTable select * from
func SelectAllFromTable(tableName string) []map[string]interface{} {

	c := []map[string]interface{}{}
	sqlCommand := fmt.Sprintf("select * from %s", tableName)
	initdb.DB.Raw(sqlCommand).Scan(&c)
	return c
}
