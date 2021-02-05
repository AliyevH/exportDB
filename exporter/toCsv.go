package exporter

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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
	keys := []int{}

	// data is a map that will hold all data from table
	data := map[int][]string{}

	// row data is used for every iteration to hold data from that row
	rowData := []string{}

	for _, row := range table {
		// Resetting rowData in every iteration to hold new row
		rowData = nil

		var id int

		for _, columnName := range columnNames {

			if columnName == "id" {
				id = int(row["id"].(int32))
				keys = append(keys, id)
			} else {
				switch row[columnName].(type) {
				case int:
					rowData = append(rowData, strconv.Itoa(row[columnName].(int)))
				case int32:
					tmp := int(row[columnName].(int32))
					rowData = append(rowData, strconv.Itoa(tmp))
				case string:
					rowData = append(rowData, row[columnName].(string))
				case nil:
					rowData = append(rowData, "")
				}
			}

		}
		data[id] = rowData
	}

	sort.Ints(keys)

	csvData := [][]string{}

	for _, id := range keys {
		d := []string{}
		d = append(d, strconv.Itoa(id))
		d = append(d, data[id]...)
		csvData = append(csvData, d)
	}

	if err := w.WriteAll(csvData); err != nil {
		log.Fatalln("error writing record to file", err)
	}

}
