package exporter

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// ToExcel function
func ToExcel() {
	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "fname")
	f.SetCellValue("Sheet1", "B1", "lname")
	f.SetCellValue("Sheet1", "C1", "age")
	f.SetCellValue("Sheet1", "A2", "hasan")
	f.SetCellValue("Sheet1", "B2", "aliyev")
	f.SetCellValue("Sheet1", "C2", 30)

	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

}
