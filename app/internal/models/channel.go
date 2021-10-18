package models

import (
	"encoding/csv"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/xuri/excelize/v2"
	"log"
	"report-creator/app/internal/helpers"
	"report-creator/app/platform/database"
	"sync"
	"time"
)

type PageChunk struct {
	page  int
	data  *[]interface{}
	chunk int
}

type PageChunkInterface interface {
	SetXlsx(*excelize.StreamWriter, []StylesColumns)
	SetCsv(writer *csv.Writer)
}

//NewPageChunk ...
func NewPageChunk(page int, data []interface{}, chunk int) PageChunkInterface {
	return &PageChunk{
		page:  page,
		data:  &data,
		chunk: chunk,
	}
}

//SetXlsx дописать в поток значения excel
func (pc *PageChunk) SetXlsx(streamWriter *excelize.StreamWriter, stylesColumns []StylesColumns) {
	var rowNumber int
	rowChunkNumber := pc.chunk * pc.page
	rowNumber = 2
	log.Printf("Запрос №%v Начал запись данных\n", pc.page+1)
	for _, data := range *pc.data {
		cellData := helpers.InterfaceSlice(data)
		cellValue := fmt.Sprintf("%v", rowChunkNumber+rowNumber)
		for i, data := range cellData {
			colName, _ := excelize.ColumnNumberToName(i + 1)
			cell := &excelize.Cell{
				StyleID: stylesColumns[i].Style,
				Formula: "",
				Value:   data,
			}
			SetDataXlsx(streamWriter, colName+cellValue, []interface{}{cell})
		}
		rowNumber++
	}
	timeNow := time.Now()
	fmt.Printf("Запрос №%v успешно отработан.Время выполнения %02d:%02d:%02d \n",
		pc.page+1, timeNow.Hour(), timeNow.Minute(), timeNow.Second())
}
func (pc *PageChunk) SetCsv(writer *csv.Writer) {
	for _, data := range *pc.data {
		cellData := helpers.InterfaceSlice(data)
		SetDataCsv(writer, cellData)
	}
}

//readChannelXlsx забор данных из  канала PageChunkInterface
func readChannelXlsx(pageChan <-chan PageChunkInterface, streamWriter *excelize.StreamWriter, style []StylesColumns) {
	for val := range pageChan {
		val.SetXlsx(streamWriter, style)
	}
}

// SetChannel создать запросы и положить в канал PageChunkInterface
func SetChannel(pageChan chan<- PageChunkInterface, countPages int, model *Save, chunk int, columns []interface{}, progress *Progress) {
	wg := sync.WaitGroup{}
	columnString := "\"" + helpers.Implode("\", \"", columns) + "\""
	for i := 0; i < countPages; i++ {
		active := sq.Select(columnString).From("(" + model.Query + ") t").Limit(uint64(chunk)).Offset(uint64(chunk * i))
		sql, _, err := active.ToSql()
		if err != nil {
			log.Print(err)
		}
		wg.Add(1)
		tempIterator := i
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in f", r)
				}
			}()
			defer wg.Done()
			log.Printf("Запрос №%v: в процессе выполнения\n", tempIterator+1)
			rows, err := database.Connect.Queryx(sql)
			if err != nil {
				log.Print(err)
				model.Callback.SendErr(err)
			}
			log.Printf("Запрос №%v: выполнился\n", tempIterator+1)
			progress.Increment()
			var result []interface{}
			for rows.Next() {
				slice, _ := rows.SliceScan()
				result = append(result, slice)

			}
			pageChan <- NewPageChunk(tempIterator, result, chunk)

		}()
	}
	go func() {
		wg.Wait()
		close(pageChan)
	}()
}
