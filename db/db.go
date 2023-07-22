package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./test.db")

	return db, err
}
