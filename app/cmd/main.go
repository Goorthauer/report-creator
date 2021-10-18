package main

import (
	"report-creator/app/internal/controllers"
)

// @title API
// @version 1.0.0
// @description Генерация отчетов XLSX. Для генерации нужен token. Сначала генерируем, а только потом скачиваем. Файл скачивается один раз - далее удаляется
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	controllers.StartServer()
}
