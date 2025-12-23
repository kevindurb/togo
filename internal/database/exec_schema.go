package database

import (
	_ "embed"
	"log"
)

//go:embed schema.sql
var ddl string

func ExecSchema() {
	db := Open()

	if _, err := db.Exec(ddl); err != nil {
		log.Fatal(err)
	}
}
