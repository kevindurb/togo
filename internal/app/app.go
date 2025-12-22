package app

import (
	"database/sql"
	"html/template"
	"log"

	"github.com/kevindurb/togo/internal/database"
	"github.com/kevindurb/togo/web"
)

type App struct {
	queries       *database.Queries
	listTodosTmpl *template.Template
}

func New() App {
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatal(err)
	}

	return App{
		queries: database.New(db),
		listTodosTmpl: template.Must(template.ParseFS(
			web.Files,
			"templates/layouts/base.gohtml",
			"templates/pages/list_todos.gohtml",
		)),
	}
}
