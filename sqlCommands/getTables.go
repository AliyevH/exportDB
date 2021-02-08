package sqlcommands

import "gitlab.com/farqlanma_nishani/dbexport/initdb"

// GetTables function returns tables list
func GetTables() []string {
	var tableName []string
	initdb.DB.Raw("show tables").Scan(&tableName)
	return tableName
}
