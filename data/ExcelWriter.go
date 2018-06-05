package data

import (
	"github.com/tealeg/xlsx"
)

//ExcelWriter create ExcelFile struct
type ExcelWriter struct {
	fileName        string
	file            *xlsx.File
	sheets          map[string]*xlsx.Sheet
	sheetNow        *xlsx.Sheet
	row, row1, row2 *xlsx.Row
	cell            *xlsx.Cell
	err             error
}

//New init ExcelWriter
func (e *ExcelWriter) New() *ExcelWriter {
	return e
}

//SetFileName set excel file name
func (e *ExcelWriter) SetFileName(fileName string) *ExcelWriter {
	e.fileName = fileName
	return e.New()
}

//InitFile init the resource of the excel write
func (e *ExcelWriter) InitFile() *ExcelWriter {
	e.file = xlsx.NewFile()
	return e.New()
}

//InitSheet init or switch sheet
func (e *ExcelWriter) InitSheet(sheetName string) *ExcelWriter {
	if _, ok := e.sheets[sheetName]; ok { //if the key as sheetName is exist
		e.sheetNow = e.sheets[sheetName]
	} else {
		e.sheetNow, e.err = e.file.AddSheet(sheetName)
		e.sheets[sheetName] = e.sheetNow
		e.errorDeal()
	}
	return e.New()
}

//AddRow add a row to excel
func (e *ExcelWriter) AddRow(oneLineData map[string]string) *ExcelWriter {

	e.row = e.sheetNow.AddRow()
	e.row.SetHeightCM(1)
	for _, v := range oneLineData {
		e.cell = e.row.AddCell()
		e.cell.Value = v
	}
	return e.New()
}

//Save create file
func (e *ExcelWriter) Save() bool {
	e.err = e.file.Save(e.fileName)
	e.errorDeal()
	return true
}
func (e *ExcelWriter) errorDeal() {
	if e.err != nil {
		panic(e.err)
	}
}

//Close close the resource of the excel write
func (e *ExcelWriter) Close() bool {
	return true
}
