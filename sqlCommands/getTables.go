package sqlcommands

import initdb "gitlab.com/farqlanma_nishani/dbexport/initDB"

// GetTables function returns tables list
func GetTables() []string {
	var tableName []string
	initdb.DB.Raw("show tables").Scan(&tableName)
	return tableName
}
