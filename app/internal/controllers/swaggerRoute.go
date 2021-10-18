package controllers

import (
	_ "report-creator/docs"

	"github.com/gorilla/mux"
	swagger "github.com/swaggo/http-swagger"
)

// SwaggerRoute Роут для свагера
func SwaggerRoute(router *mux.Router) {
	router.PathPrefix("/swagger/").Handler(swagger.WrapHandler)
}
