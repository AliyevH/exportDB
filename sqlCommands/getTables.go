package sqlcommands

import initdb "github.com/AliyevH/exportDB/initDb"

// GetTables function returns tables list
func GetTables() []string {
	db := initdb.DB

	var tableName []string
	db.Raw("show tables").Scan(&tableName)
	return tableName
}
