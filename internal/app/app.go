package app

import (
	"html/template"

	"github.com/kevindurb/togo/internal/database"
	"github.com/kevindurb/togo/web"
)

type App struct {
	queries       *database.Queries
	listTodosTmpl *template.Template
}

func New() App {
	db := database.MustOpen()

	return App{
		queries: database.New(db),
		listTodosTmpl: template.Must(template.ParseFS(
			web.Files,
			"templates/layouts/base.gohtml",
			"templates/pages/list_todos.gohtml",
		)),
	}
}
