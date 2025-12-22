package database

import (
	"database/sql"
	_ "embed"
	"log"
)

//go:embed schema.sql
var ddl string

func ExecSchema() {
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(ddl); err != nil {
		log.Fatal(err)
	}
}
