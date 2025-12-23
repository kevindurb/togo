package main

import (
	"log"
	"net/http"

	"github.com/kevindurb/togo/internal/app"
	"github.com/kevindurb/togo/internal/database"
)

func main() {
	database.MigrateAll()

	mux := http.NewServeMux()

	app := app.New()

	mux.Handle("GET /static/", http.StripPrefix("/static/", app.StaticFileServer()))
	mux.HandleFunc("GET /todos", app.ProtectRoute(app.ListTodos))
	mux.HandleFunc("POST /todos", app.ProtectRoute(app.CreateTodo))
	mux.HandleFunc("POST /users", app.CreateUser)
	mux.HandleFunc("GET /users/new", app.NewUser)
	mux.HandleFunc("POST /login", app.Login)
	mux.HandleFunc("GET /login", app.ShowLogin)
	mux.Handle("GET /", http.RedirectHandler("/todos", http.StatusFound))

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
