package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

const DirectoryFiles = "files"

//BeginTicker если включено то тикер автоудаляет файлы по таймеру
func BeginTicker() {
	minutes, err := strconv.Atoi(os.Getenv("AUTO_DELETE_FILES_MINUTES"))
	if err != nil {
		fmt.Println(err)
		minutes = 15
	}
	ticker := time.NewTicker(time.Minute * time.Duration(minutes))
	log.Printf("AUTO DELETE FILE ON")
	defer ticker.Stop()
	for tick := range ticker.C {
		removeFilesFromDirectory(tick)
	}
}

//removeFilesFromDirectory Удалить файлы из папки files, если они живут дольше чем 60 минут
func removeFilesFromDirectory(t time.Time) {
	log.Printf("Запускается автоотчистка файлов")
	files, err := ioutil.ReadDir(DirectoryFiles)
	if err != nil {
		log.Println(err)
	}
	for _, file := range files {
		if t.Sub(file.ModTime()).Minutes() > 30 {
			_, err := removeFiles(DirectoryFiles + "/" + file.Name())
			if err != nil {
				log.Print(err)
			}
		}
	}
	log.Printf("Автоотчистка файлов завершена")
}

//removeFiles механизм удаление файлов
func removeFiles(fileName string) (bool, error) {
	err := os.Remove(fileName)
	if err != nil {
		return false, err
	}
	return true, nil
}
