package exporter

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"
)

// ExportToCsv exports to csv
func ExportToCsv(path string, table []map[string]interface{}, tableName string, columnNames []string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create csv file
	fileName := fmt.Sprintf("%s/%s.csv", path, tableName)
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	// Close file at the end
	defer f.Close()

	// New Writer
	w := csv.NewWriter(f)
	defer w.Flush()

	if err := w.Write(columnNames); err != nil {
		log.Fatalln("error writing record to file", err)
	}

	// Keys are id's. We will use this to access data in sorted by id order
	csvData := PrepareData(table, columnNames)

	if err := w.WriteAll(csvData); err != nil {
		log.Fatalln("error writing record to file", err)
	}

}
