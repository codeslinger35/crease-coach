package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", http.NotFound) // Catch-all route
	r.HandleFunc("/health", app.health)
	r.HandleFunc("/init", app.init)
	r.HandleFunc("/save", app.save)

	r.HandleFunc("/goalies", app.goalieHandler)
	r.HandleFunc("/goalies/{id}", app.goalieByIdHandler)

	return r
}
