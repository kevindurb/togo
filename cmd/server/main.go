package main

import (
	"log"
	"net/http"

	"github.com/kevindurb/togo/internal/app"
)

func main() {
	mux := http.NewServeMux()

	app := app.New()

	mux.Handle("GET /", http.RedirectHandler("/todos", http.StatusMovedPermanently))
	mux.HandleFunc("GET /todos", app.ListTodos)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
