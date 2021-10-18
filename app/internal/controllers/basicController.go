package controllers

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"report-creator/app/internal/helpers"
	"time"
)

var router = mux.NewRouter()

//beforeController перед любым действием проверка на JWT
func beforeController(w http.ResponseWriter, r *http.Request) (error, int) {
	now := time.Now().Unix()
	// Get claims from JWT.
	claims, err := helpers.ExtractTokenMetadata(r)
	if err != nil {
		return errors.New("Токен отсутствует"), 500
	}

	// Set expiration time from JWT data of current book.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		http.Error(w, http.StatusText(401), 401)
		return errors.New("Время работы токена истекло!"), 401

	}

	return nil, 200
}
