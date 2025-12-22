package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Connect() sqlx.DB {
	db, err := sqlx.Connect("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	return *db
}
