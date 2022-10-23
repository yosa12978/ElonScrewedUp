package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Connect(connectionString string) error {
	database, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return err
	}
	err = database.Ping()
	db = database
	return err
}

func GetDB() *sql.DB {
	return db
}
