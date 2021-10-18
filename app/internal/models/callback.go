package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"report-creator/app/internal/helpers"
)

type CallbackInterface interface {
	SendErr(err error)
	SendFinish()
	SendStart()
	SendStatus(string)
	SetBody(string)
}

type Callback struct {
	Start  string
	Finish string
	Status string
	Failed string
	Body   string
}

//SendStart отправка сообщения при старте
func (cb *Callback) SendStart() {
	log.Printf("Send start to %v\n", cb.Start)
	data := url.Values{"path": {cb.Body}}
	helpers.SendPostMessage(cb.Start, data)
}

//SendFinish отправка сообщения при финише
func (cb *Callback) SendFinish() {
	log.Printf("Send finish to %v\n", cb.Finish)
	data := url.Values{"path": {cb.Body}}
	helpers.SendPostMessage(cb.Finish, data)
}

//SendStatus отправка сообщения при финише
func (cb *Callback) SendStatus(progressStatus string) {
	log.Printf("Send status to %v\n", cb.Status)
	log.Printf("progress: %s percent", progressStatus)
	data := url.Values{"path": {cb.Body}, "progress": {progressStatus}}
	helpers.SendPostMessage(cb.Status, data)
}

//SendErr отправка сообщения при ошибке
func (cb *Callback) SendErr(err error) {
	log.Printf("Send err to %v. err: %v\n", cb.Failed, err)
	data := url.Values{"error": {err.Error()}}
	helpers.SendPostMessage(cb.Failed, data)
}

//SetBody добавить body строку
func (cb *Callback) SetBody(fileName string) {
	cb.Body = fileName
}

//NewCallback создает новый CallbackInterface
func NewCallback(params map[string]interface{}) CallbackInterface {
	callback := new(Callback)
	if val, ok := params["callback"]; ok {
		jsonMarshalBody, err := json.Marshal(val)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(jsonMarshalBody, &callback)
		if err != nil {
			fmt.Println(err)
		}
	}
	return callback
}
