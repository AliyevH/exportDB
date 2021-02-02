
<h3 align="center">Database Exporter Tool</h3>

<!-- <div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/kylelobo/The-Documentation-Compendium.svg)](https://github.com/MadeByMads/mad-migration/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/kylelobo/The-Documentation-Compendium.svg)](https://github.com/MadeByMads/mad-migration/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div> -->


## 🧐 About <a name = "about"></a>

The Database Exporter Tool written in go was designed for those looking to export their data from database to csv. Currently, supported databases are the MySQL and Mariadb.


## 🏁 Getting Started <a name = "getting_started"></a>

### Installing

```bash
go get github.com/AliyevH/exportDB
```

## 🎈 Usage <a name="usage"></a>

```go
package main

import (
	"fmt"
	"sync"

	"github.com/AliyevH/exportDB/configs"
	"github.com/AliyevH/exportDB/db"
	"github.com/AliyevH/exportDB/exporter"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
    DB, err := gorm.Open(mysql.Open("root:admin@tcp(127.0.0.1:3306)/fn?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
    
	if err != nil {
        panic(err.Error())
    }

	// db from exporter 
	// We have created function to get table names from db.
	tables := db.GetTables()

	var wg sync.WaitGroup

	for _, table := range tables {
		wg.Add(1)

        // db from exporter SelectAllFromTable (select * from 'table')
        data := db.SelectAllFromTable(table)
        
        // db from exporter GetColumnNames (desc 'table')
		columnNames := db.GetColumnNames(table)

        // For every table we are creating csv file. Folder must exist.
		go exporter.ExportToExcel("./database", data, table, columnNames, &wg)
		wg.Wait()
	}

}
```

