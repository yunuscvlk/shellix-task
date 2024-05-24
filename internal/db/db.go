package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	DB *sql.DB
	LastError error
}

func Initialize() *DB {
	instance := new(DB)
	
	instance.DB, instance.LastError = sql.Open("sqlite3", "./static/shellix.db")

	if instance.LastError != nil {
		panic(instance.LastError)
	}

	return instance
}