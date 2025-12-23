package app

import (
	"log"
	"net/http"

	"github.com/kevindurb/togo/internal/database"
	"github.com/kevindurb/togo/internal/utils"
)

type CreateTodoBody struct {
	Description string `schema:"description,required"`
}

type ListTodosPageData struct {
	Todos []database.ListTodosRow
}

func (a *App) ListTodos(w http.ResponseWriter, r *http.Request) {
	userID := userIDFromContext(r.Context())
	todos, _ := a.queries.ListTodos(r.Context(), userID)
	a.listTodosTmpl.ExecuteTemplate(w, "base", ListTodosPageData{
		Todos: todos,
	})
}

func (a *App) CreateTodo(w http.ResponseWriter, r *http.Request) {
	userID := userIDFromContext(r.Context())
	var body CreateTodoBody
	if err := utils.DecodePostForm(&body, r, a.decoder); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := a.queries.CreateTodo(r.Context(), database.CreateTodoParams{
		Description: body.Description,
		UserID:      userID,
	})

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/todos", http.StatusFound)
}
