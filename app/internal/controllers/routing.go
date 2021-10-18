package controllers

import "github.com/gorilla/mux"

//GetHandle получить роуты
func GetHandle(router *mux.Router) {
	//router.HandleFunc("/token", generateToken)
	router.HandleFunc("/report", create).Methods("POST")
	router.HandleFunc("/report/{name}", download)

}
