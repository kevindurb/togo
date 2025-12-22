package app

import (
	"net/http"
)

func (a *App) ListTodos(w http.ResponseWriter, r *http.Request) {
	a.listTodosTmpl.ExecuteTemplate(w, "base", nil)
}
