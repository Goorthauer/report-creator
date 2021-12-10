package models

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"

	"report-creator/app/internal/helpers"
)

var MaxRowExcel = 1048576

type Save struct {
	Chunk         ChunkInterface
	Callback      CallbackInterface
	Columns       ColumnsMapInterface
	Query         string
	ShortFileName string
}

//SaveInterface ...
type SaveInterface interface {
	Create()
}

//NewSaveType создаем новый тип сохранения
func NewSaveType(params map[string]interface{}, shortFileName string) SaveInterface {
	var columns ColumnsMapInterface
	var query string
	chunk := NewChunkType()
	if val, ok := params["chunk"]; ok {
		paramsChunk, err := strconv.Atoi(fmt.Sprintf("%v", val))
		if err != nil {
			fmt.Println(err)
		}
		chunk.AssignValue(paramsChunk)
	}
	callback := NewCallback(params)
	if val, ok := params["columns"]; ok {
		columns = GetColumnsMap(helpers.InterfaceSlice(val))
	}
	if val, ok := params["query"]; ok {
		query = fmt.Sprintf("%v", val)
	}
	return &Save{
		Chunk:         chunk,
		Callback:      callback,
		Columns:       columns,
		Query:         query,
		ShortFileName: shortFileName,
	}
}

//Create Создание XLSX с присвоением наименования по параметру
func (model *Save) Create() {
	var exists bool
	start := time.Now()
	progress := NewProgress(model.Callback)
	model.Callback.SetBody(helpers.GenerateReportLink(model.ShortFileName))
	model.Callback.SendStart()
	progress.StandartStep()
	countSheet := 1
	fmt.Printf("Получаем количество записей\n")
	countRow, err := model.Chunk.GetCountRow(model.Query)
	if err != nil {
		model.Callback.SendErr(err)
		return
	}
	fmt.Printf("Количество записей:%v\n", countRow)
	countPages, err := model.Chunk.GetCountPages(countRow)
	progress.StandartStep()
	if err != nil {
		model.Callback.SendErr(err)
		return
	}
	progress.setTotal(countPages)
	fmt.Printf("Количество страниц:%v\n", countPages)
	valueChunk := model.Chunk.GetValue()
	pageChan := make(chan PageChunkInterface, countPages)
	mapColumns := model.Columns.GetColumn("Name")
	SetChannel(pageChan, countPages, model, valueChunk, mapColumns, progress)
	if countRow/MaxRowExcel > 1 {
		countSheet = countRow / MaxRowExcel
	}
	if countSheet > 1 {
		exists = model.createCsv(pageChan)
	} else {
		exists = model.createXlsx(pageChan)
	}
	if exists {
		elapsed := time.Since(start)
		logData := fmt.Sprintf("Файл записан. %v", elapsed)
		fmt.Printf(logData + "\n")
		progress.Finish()
		model.Callback.SendFinish()
	}
}

func (model *Save) createXlsx(pageChan chan PageChunkInterface) bool {
	fileName := "files/" + model.ShortFileName + ".xlsx"
	file := excelize.NewFile()
	model.Columns.SetStyle(file, "Sheet1")
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		model.Callback.SendErr(err)
		return false
	}
	style := model.Columns.GetStyle(file)
	model.Columns.SetHeader(streamWriter)
	readChannelXlsx(pageChan, streamWriter, style)
	if err := streamWriter.Flush(); err != nil {
		model.Callback.SendErr(err)
		return false
	}
	if err := file.SaveAs(fileName); err != nil {
		model.Callback.SendErr(err)
		return false
	}
	return true
}

func (model *Save) createCsv(pageChan chan PageChunkInterface) bool {
	fileName := "files/" + model.ShortFileName + ".csv"
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer func() {
		_ = f.Close()
	}()
	writer := csv.NewWriter(f)
	writer.Comma = '\t'
	defer writer.Flush()
	columnsName := model.Columns.GetColumn("Label")
	SetDataCsv(writer, columnsName)
	for val := range pageChan {
		val.SetCsv(writer)
	}
	return true
}
