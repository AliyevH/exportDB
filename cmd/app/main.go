package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/AliyevH/exportDB/db"
	"github.com/AliyevH/exportDB/exporter"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	tables := db.GetTables()
	var wg sync.WaitGroup
	start := time.Now()

	for _, table := range tables {
		wg.Add(1)
		// if i == 10 {
		// 	break
		// }
		data := db.SelectAllFromTable(table)
		columnNames := db.GetColumnNames(table)
		// fmt.Println(columnNames)
		// fmt.Println(data)

		go exporter.ExportToExcel("./database", data, table, columnNames, &wg)
		wg.Wait()
	}

	finished := time.Since(start)
	fmt.Println("Finished in", finished)

}
