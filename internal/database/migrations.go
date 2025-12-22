package database

import (
	"embed"
	"log"
)

//go:embed migrations/*
var migrationFiles embed.FS

func MigrateAll() error {
	db := Connect()
	files, err := migrationFiles.ReadDir("migrations")
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		log.Printf("Applying migration: %s...", fileName)
		content, err := migrationFiles.ReadFile("migrations/" + fileName)
		if err != nil {
			return err
		}

		_, err = db.Exec(string(content))
		if err != nil {
			return err
		}

		log.Println("Done!")
	}

	return nil
}
