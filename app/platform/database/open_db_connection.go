package database

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var Connect *sqlx.DB

//GetConnect Получить коннект к БД и использовать
func GetConnect(dataBaseName string) (*sqlx.DB, error) {
	switch dataBaseName {
	case "postgresql":
		return PostgreSQLConnection()
	default:
		return PostgreSQLConnection()
	}
}

func InitDB() {
	var err error
	dataBaseName := os.Getenv("DB_NAME_CONNECTION")
	Connect, err = GetConnect(dataBaseName)
	if err != nil {
		log.Panic(err)
	}

	if err = Connect.Ping(); err != nil {
		defer Connect.Close()
		log.Panic(err)
	}
}
