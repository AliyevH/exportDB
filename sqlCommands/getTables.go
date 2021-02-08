package sqlcommands

import "github.com/AliyevH/exportDB/initdb"

// GetTables function returns tables list
func GetTables() []string {
	var tableName []string
	initdb.DB.Raw("show tables").Scan(&tableName)
	return tableName
}
