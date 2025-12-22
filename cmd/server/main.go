package main

import (
	"github.com/kevindurb/togo/internal/app"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	app := app.App{}

	mux.Handle("GET /", http.RedirectHandler("/todos", http.StatusMovedPermanently))
	mux.HandleFunc("GET /todos", app.ListTodos)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
