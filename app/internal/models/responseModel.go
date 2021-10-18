package models

//PostMessageBody  пример структуры запроса POST
type PostMessageBody struct {
	Query    string            `json:"query" example:"select * from cheque.cheques order by id desc limit 111"`
	Category string            `json:"category" example:"cheque"`
	Chunk    int               `json:"chunk" example:"5000"`
	Callback map[string]string `json:"callback" example:"start:127.0.0.1/start,status:127.0.0.1/status,finish:127.0.0.1/finish,failed:127.0.0.1/failed"`
	Columns  []struct {
		Name   string `json:"name" example:"cheque_date"`
		Label  string `json:"label" example:"Время чека"`
		Format string `json:"format" example:"datetime"`
	} `json:"columns" `
}

//GetResponseBody пример структуры ответа на генерацию токена
type GetResponseBody struct {
	Token  string `json:"access_token"`
	Status int    `json:"status"`
}

//Response быстрый ответ на функцию
type Response struct {
	Route string `json:"route"`
}
