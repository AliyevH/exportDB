package exporter

import (
	"fmt"
	"sort"
	"strconv"
)

// PrepareData convert table values
func PrepareData(table []map[string]interface{}, columnNames []string) [][]string {
	keys := []int{}

	// data is a map that will hold all data from table
	data := map[int][]string{}

	// row data is used for every iteration to hold data from that row

	for _, row := range table {
		// Resetting rowData in every iteration to hold new row
		rowData := []string{}

		var id int

		for _, columnName := range columnNames {

			if columnName == "id" || columnName == "Id" || columnName == "ID" {
				id = int(row[columnName].(int))
				keys = append(keys, id)
			} else {
				switch row[columnName].(type) {
				case int:
					rowData = append(rowData, strconv.Itoa(row[columnName].(int)))
				case int16:
					tmp := fmt.Sprintf("%d", row[columnName])
					rowData = append(rowData, tmp)
				case int32:
					tmp := fmt.Sprintf("%d", row[columnName])
					rowData = append(rowData, tmp)
				case int64:
					tmp := fmt.Sprintf("%d", row[columnName])
					rowData = append(rowData, tmp)
				case float32:
					tmp := fmt.Sprintf("%f", row[columnName])
					rowData = append(rowData, tmp)
				case float64:
					tmp := fmt.Sprintf("%f", row[columnName])
					rowData = append(rowData, tmp)
				case string:
					rowData = append(rowData, row[columnName].(string))
				case bool:
					rowData = append(rowData, strconv.FormatBool(row[columnName].(bool)))
				default:
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
	return csvData
}
