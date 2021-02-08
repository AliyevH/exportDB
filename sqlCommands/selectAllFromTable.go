package sqlcommands

import (
	"fmt"

	"gitlab.com/farqlanma_nishani/dbexport/initdb"
)

// SelectAllFromTable select * from
func SelectAllFromTable(tableName string) []map[string]interface{} {

	c := []map[string]interface{}{}
	sqlCommand := fmt.Sprintf("select * from %s", tableName)
	initdb.DB.Raw(sqlCommand).Scan(&c)
	return c
}
