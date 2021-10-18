package models

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strings"
)

//FileStyle стили генерация стилей для определенного XLSX файла
type FileStyle struct {
	file *excelize.File
}

type FileStyleInterface interface {
	GetStyle(interface{}) int
}

func NewFileStyle(file *excelize.File) FileStyleInterface {
	return &FileStyle{file: file}
}

//GetStyle Получить стиль по имени
func (f FileStyle) GetStyle(styleParams interface{}) int {
	formatterStyle := setStyle(styleParams)
	formatterStyle.Alignment = &excelize.Alignment{WrapText: true}
	style, err := f.file.NewStyle(&formatterStyle)
	if err != nil {
		log.Print(err)
	}
	return style
}

//setStyle Получить стиль исходя из параметров
func setStyle(styleParams interface{}) excelize.Style {
	params := fmt.Sprint(styleParams)
	switch format := strings.ToLower(params); format {
	case "datetime":
		exp := "YYYY-MM-DD HH:MM:SS"
		return excelize.Style{CustomNumFmt: &exp}
	case "date":
		exp := "YYYY-MM-DD"
		return excelize.Style{CustomNumFmt: &exp}
	case "time":
		return excelize.Style{NumFmt: 21}
	case "currency":
		return excelize.Style{NumFmt: 278}
	case "percent":
		return excelize.Style{NumFmt: 9}
	default:
		return excelize.Style{}
	}
}
