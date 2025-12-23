package main

import (
	"log"
	"net/http"

	"github.com/kevindurb/togo/internal/app"
	"github.com/kevindurb/togo/internal/database"
)

func main() {
	database.ExecSchema()

	mux := http.NewServeMux()

	app := app.New()

	mux.Handle("GET /static/", http.StripPrefix("/static/", app.StaticFileServer()))
	mux.Handle("GET /", http.RedirectHandler("/todos", http.StatusMovedPermanently))
	mux.HandleFunc("GET /todos", app.ListTodos)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
