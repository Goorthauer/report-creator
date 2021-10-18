package models

var WorkerPull chan struct{}

//CreateWork Создаем работу с ограничением по числу воркеров
func CreateWork(params map[string]interface{}, shortFileName string) {
	WorkerPull <- struct{}{}
	defer func() { <-WorkerPull }()
	model := NewSaveType(params, shortFileName)
	model.Create()
}
