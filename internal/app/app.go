package app

import (
	"html/template"

	"github.com/gorilla/schema"
	"github.com/kevindurb/togo/internal/database"
	"github.com/kevindurb/togo/web"
)

type App struct {
	queries       *database.Queries
	decoder       *schema.Decoder
	listTodosTmpl *template.Template
	showLoginTmpl *template.Template
	newUserTmpl   *template.Template
}

func New() App {
	db := database.MustOpen()

	return App{
		queries: database.New(db),
		decoder: schema.NewDecoder(),
		listTodosTmpl: template.Must(template.ParseFS(
			web.Files,
			"templates/layouts/base.gohtml",
			"templates/pages/list_todos.gohtml",
		)),
		showLoginTmpl: template.Must(template.ParseFS(
			web.Files,
			"templates/layouts/base.gohtml",
			"templates/pages/show_login.gohtml",
		)),
		newUserTmpl: template.Must(template.ParseFS(
			web.Files,
			"templates/layouts/base.gohtml",
			"templates/pages/new_user.gohtml",
		)),
	}
}
