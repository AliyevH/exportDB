package db

// GetTables function returns tables list
func GetTables() []string {
	var tableName []string
	DB.Raw("show tables").Scan(&tableName)
	return tableName
}
