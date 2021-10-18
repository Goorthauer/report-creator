package controllers

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"report-creator/app/internal/models"
	"report-creator/app/platform/database"
	"strconv"
)

func StartServer() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	_ = os.Mkdir("files", 0755)
	MaxWorkerPool, err := strconv.Atoi(os.Getenv("MAX_COUNT_WORKER"))
	if err != nil {
		fmt.Println(err)
	}
	models.WorkerPull = make(chan struct{}, MaxWorkerPool)
	if os.Getenv("AUTO_DELETE_MODE") == "ON" {
		go BeginTicker()
	}

	GetHandle(router)
	SwaggerRoute(router)
	database.InitDB()
	bindAddr := os.Getenv("BIND_ADDR")
	log.Printf("SERVER ON!")
	err = http.ListenAndServe(bindAddr, router)
	if err != nil {
		fmt.Println(err)
	}
}
