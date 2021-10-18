package models

import (
	"report-creator/app/platform/database"
)

//chunkType Нам он нужен что бы работать с ним через интерфейс,а не на прямую
type chunkType int

const maxCountPages = 12

type ChunkInterface interface {
	GetCountPages(int) (int, error)
	AssignValue(int)
	GetValue() int
	GetCountRow(string) (int, error)
}

//NewChunkType ...
func NewChunkType() ChunkInterface {
	return new(chunkType)
}

//AssignValue назначить новое значение для chunkType
func (ch *chunkType) AssignValue(value int) {
	*ch = chunkType(value)
}

//GetValue получить значение chunkType
func (ch *chunkType) GetValue() int {
	return int(*ch)
}

//GetCountPages получить количество страниц по чанку и, если страниц > maxCountPages тогда меняем чанк
func (ch *chunkType) GetCountPages(count int) (int, error) {
	chunk := ch.GetValue()
	var countPages int

	if chunk == 0 {
		chunk = count
	}
	if chunk >= count {
		countPages = 1
	} else {
		countPages = count/chunk + 1
	}
	if countPages > maxCountPages {
		countPages = maxCountPages
		chunk = count / countPages
	}
	ch.AssignValue(chunk)
	return countPages, nil
}

//GetCountRow Получить количество строк
func (ch *chunkType) GetCountRow(query string) (int, error) {
	var count int
	row := database.Connect.QueryRow("SELECT COUNT(*) FROM (" + query + ") sub")
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
