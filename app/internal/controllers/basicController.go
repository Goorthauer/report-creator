package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"report-creator/app/internal/helpers"
)

var router = mux.NewRouter()

//beforeController перед любым действием проверка на JWT
func beforeController(w http.ResponseWriter, r *http.Request) (int, error) {
	now := time.Now().Unix()
	// Get claims from JWT.
	claims, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		return 500, errors.New("Токен отсутствует")
	}

	// Set expiration time from JWT data of current book.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		http.Error(w, http.StatusText(401), 401)
		return 401, errors.New("Время работы токена истекло!")

	}

	return 200, err
}
