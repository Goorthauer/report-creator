package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xuri/excelize/v2"
)

// report func gets file export.
// @Description Получить файл xlsx.
// @Summary получить xlsx
// @Tags Report
// @Accept json
// @Produce json
// @Param name path string true "names"
// @Success 200 {string} status "ok"
// @Router /report/{name} [get]
func download(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	params := mux.Vars(r)
	fileName := "files/" + params["name"] + ".xlsx"
	_, err := excelize.OpenFile(fileName)
	if err != nil {
		fileName = "files/" + params["name"] + ".csv"
	}
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment;filename="+params["name"]+".xlsx")
	http.ServeFile(w, r, fileName)

}
