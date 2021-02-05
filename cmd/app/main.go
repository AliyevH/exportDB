package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/AliyevH/exportDB/exporter"
	"github.com/AliyevH/exportDB/scripts"
	sqlcommands "github.com/AliyevH/exportDB/sqlCommands"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	tables := sqlcommands.GetTables()
	var wg sync.WaitGroup
	start := time.Now()

	scripts.MakeDirectoryIfNotExists("./database")

	for _, table := range tables {
		wg.Add(1)
		// if i == 10 {
		// 	break
		// }
		data := sqlcommands.SelectAllFromTable(table)
		columnNames := sqlcommands.GetColumnNames(table)

		go exporter.ExportToCsv("./database", data, table, columnNames, &wg)
		wg.Wait()
	}

	finished := time.Since(start)
	fmt.Println("Finished in", finished)

}
