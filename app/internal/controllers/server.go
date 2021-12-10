package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"report-creator/app/internal/models"
	"report-creator/app/platform/database"
)

func StartServer() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	_ = os.Mkdir("files", 0755)
	maxWorkerPool, err := strconv.Atoi(os.Getenv("MAX_COUNT_WORKER"))
	if err != nil {
		fmt.Println(err)
	}
	models.WorkerPull = make(chan struct{}, maxWorkerPool)
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
