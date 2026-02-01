package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/terryluciano/templ-test/internal/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", handlers.HomeHandler)

	http.ListenAndServe(":3000", r)
}
