package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"report-creator/app/internal/helpers"
	"report-creator/app/internal/models"
)

// create func generate xlsx file.
// @Description Создает новый xlsx отчет.
// @Summary Создание нового xlsx отчета
// @Tags Report
// @Accept json
// @Produce json
// @Param body body models.PostMessageBody true "body"
// @Success 200 {object} models.Response
// @Security ApiKeyAuth
// @Router /report [post]
func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	needJwt := os.Getenv("JWT_AUTORIZE")
	if needJwt == "ON" {
		status, err := beforeController(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			_ = json.NewEncoder(w).Encode(err)
			return
		}
	}

	start := time.Now()
	var params map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	category := fmt.Sprintf("%v", params["category"])
	shortFileName := category + "_" + fmt.Sprint(start.UnixNano())
	go models.CreateWork(params, shortFileName)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	route := helpers.GenerateReportLink(shortFileName)
	response := models.Response{
		Route: route,
	}
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}
