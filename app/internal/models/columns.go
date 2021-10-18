package models

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"reflect"
	"report-creator/app/internal/helpers"
)

type StylesColumns struct {
	Style int
}
type ColumnMap struct {
	Name   string `json:"name" `
	Label  string `json:"label"`
	Format string `json:"format"`
}

type ColumnInterface interface {
	getField(string) (string, bool)
}

//getField Получить значения из columnMap и его валидность
func (columnMap *ColumnMap) getField(field string) (string, bool) {
	r := reflect.ValueOf(columnMap)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String(), f.IsValid()
}

type columnsMap struct {
	Items []ColumnMap
}

//ColumnsMapInterface интерфейс колонок. вся работа с колонками должна проходить только через них
type ColumnsMapInterface interface {
	Append(ColumnMap)
	GetColumn(string) []interface{}
	SetStyle(file *excelize.File, sheetName string)
	SetHeader(*excelize.StreamWriter)
	GetStyle(file *excelize.File) []StylesColumns
}

//GetColumn Получить значения колонок по наименованию
func (cm *columnsMap) GetColumn(columnName string) []interface{} {
	var mapColumnFormat []interface{}
	for _, val := range cm.Items {
		name, ok := val.getField(columnName)
		if ok {
			mapColumnFormat = append(mapColumnFormat, name)
		}
	}
	return mapColumnFormat
}

//Append добавить строку ColumnMap в ColumnsMapInterface
func (cm *columnsMap) Append(item ColumnMap) {
	cm.Items = append(cm.Items, item)
}

//SetStyle установить стили для колонок
func (cm *columnsMap) SetStyle(file *excelize.File, sheetName string) {
	columnsFormat := cm.GetColumn("Format")
	columnsMap := helpers.InterfaceToMap(columnsFormat)
	setAutoFilter(columnsMap, file, sheetName)
	//style := NewFileStyle(file)
	//for i, val := range columnsMap {
	//	format := style.GetStyle(val)
	//	coordinate, err := excelize.ColumnNumberToName(i + 1)
	//	if err != nil {
	//		log.Print(err)
	//	}
	//	err = file.SetColStyle(sheetName, coordinate, format)
	//	if err != nil {
	//		log.Print(err)
	//	}
	//	from, _ := excelize.CoordinatesToCellName(i+1, 2)
	//	to, _ := excelize.CoordinatesToCellName(i+1, countRow)
	//	err = file.SetCellStyle(sheetName, from, to, format)
	//}
}

func (cm *columnsMap) GetStyle(file *excelize.File) []StylesColumns {
	var stylesColumn []StylesColumns
	columnsFormat := cm.GetColumn("Format")
	columnsMap := helpers.InterfaceToMap(columnsFormat)
	style := NewFileStyle(file)
	for _, val := range columnsMap {
		stylesColumn = append(stylesColumn, StylesColumns{
			Style: style.GetStyle(val),
		})
	}
	return stylesColumn
}

//setAutoFilter Добавить автофильтр к каждой колонке
func setAutoFilter(columnsMap []string, file *excelize.File, sheetName string) {
	firstColumn := fmt.Sprintf("%v%v", "A", 1)
	coordinate, err := excelize.ColumnNumberToName(len(columnsMap))
	if err != nil {
		coordinate = "A"
	}
	lastColumn := fmt.Sprintf("%v%v", coordinate, 1)
	err = file.AutoFilter(sheetName, firstColumn, lastColumn, "")
	if err != nil {
		log.Print(err)
	}
	//err = file.SetColWidth(sheetName, "A", coordinate, 33)
	//if err != nil {
	//	log.Print(err)
	//}
}

//SetHeader добавить шапку
func (cm *columnsMap) SetHeader(streamWriter *excelize.StreamWriter) {
	columnsName := cm.GetColumn("Label")
	err := streamWriter.SetColWidth(1, len(columnsName), 24)
	if err != nil {
		log.Print(err)
	}
	for i, data := range columnsName {
		colName, _ := excelize.ColumnNumberToName(i + 1)
		cell := &excelize.Cell{
			StyleID: 0,
			Formula: "",
			Value:   data,
		}
		SetDataXlsx(streamWriter, colName+"1", []interface{}{cell})
	}
}

func newColumnsMap() ColumnsMapInterface {
	model := new(columnsMap)
	return model
}

func GetColumnsMap(columns []interface{}) ColumnsMapInterface {
	columnsMap := newColumnsMap()
	for _, val := range columns {

		columnMap := new(ColumnMap)
		jsonMarshalBody, err := json.Marshal(val)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(jsonMarshalBody, &columnMap)
		if err != nil {
			fmt.Println(err)
		}
		if columnMap.Name != "" {
			columnsMap.Append(*columnMap)
		}
	}
	return columnsMap
}

//SetDataXlsx Записать строку из массива slice
func SetDataXlsx(streamWriter *excelize.StreamWriter, cellValue string, val []interface{}) {
	if err := streamWriter.SetRow(cellValue, val); err != nil {
		fmt.Println(err)
	}
}
func SetDataCsv(writer *csv.Writer, values []interface{}) {
	mapValues := helpers.InterfaceToMap(values)
	if err := writer.Write(mapValues); err != nil {
		log.Fatalln("error writing record to file", err)
	}
}
