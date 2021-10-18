package helpers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
)

//InterfaceSlice Перевод из interface в slice
func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

//GenerateRoute генерирует универсныльй путь к файлам
func GenerateRoute() string {
	ipAddress := os.Getenv("IP_ADDRESS")
	return ipAddress + "/report/"
}

//GenerateReportLink генерирует окончательный адрес до файла
func GenerateReportLink(fileName string) string {
	return GenerateRoute() + fileName
}

//Implode Перевод массива в строку
func Implode(glue string, ci []interface{}) string {
	data := InterfaceToMap(ci)
	return strings.Join(data, glue)
}

//InterfaceToMap Перевод массива интерфейсов в массив строк
func InterfaceToMap(ci []interface{}) []string {
	data := make([]string, len(ci))
	for i, v := range ci {
		if fmt.Sprint(v) != "" {
			data[i] = fmt.Sprint(v)
		}
	}
	return data
}

//SendPostMessage отправка POST запроса по адресу
func SendPostMessage(url string, data url.Values) {
	go func() {
		_, err := http.PostForm(url, data)
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
}
