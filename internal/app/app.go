package app

import (
	"github.com/kevindurb/togo/web"
	"html/template"
)

type App struct {
	listTodosTmpl *template.Template
}

func New() App {
	return App{
		listTodosTmpl: template.Must(template.ParseFS(
			web.Files,
			"templates/layouts/base.gohtml",
			"templates/pages/list_todos.gohtml",
		)),
	}
}
