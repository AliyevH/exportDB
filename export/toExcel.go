package export

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	sqlcommands "gitlab.com/farqlanma_nishani/dbexport/sqlCommands"
)

// ToExcel Export to Excel
func ToExcel(tableName string, columns string, result *[]map[string]interface{}, wg *sync.WaitGroup) {

	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var columnsSlice []string

	if columns == "*" {
		columnsSlice = sqlcommands.GetColumnNames(tableName)
	} else {
		columnsSlice = strings.Split(columns, ",")

	}

	f := excelize.NewFile()

	sheet := "Sheet1"

	for i, v := range columnsSlice {
		header := fmt.Sprintf("%s1", string(alphabet[i]))
		headerValue := v
		f.SetCellValue(sheet, header, headerValue)
	}

	for index, v := range *result {
		for alphIndex, column := range columnsSlice {
			header := fmt.Sprintf("%s%d", string(alphabet[alphIndex]), index+2)
			headerValue := v[column]

			f.SetCellValue(sheet, header, headerValue)
		}
	}
	unixTime := time.Now().Unix()

	fileName := fmt.Sprintf("%s-%d.xlsx", tableName, unixTime)
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}

	wg.Done()

}
