package app

import (
	"html/template"

	"github.com/jmoiron/sqlx"
	"github.com/kevindurb/togo/internal/database"
	"github.com/kevindurb/togo/web"
)

type App struct {
	listTodosTmpl *template.Template
	db            *sqlx.DB
}

func New() App {
	db := database.Connect()

	return App{
		db: &db,
		listTodosTmpl: template.Must(template.ParseFS(
			web.Files,
			"templates/layouts/base.gohtml",
			"templates/pages/list_todos.gohtml",
		)),
	}
}
