package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/terryluciano/templ-test/internal/config"
	"github.com/terryluciano/templ-test/internal/database"
	"github.com/terryluciano/templ-test/internal/handler"
)

func main() {
	config.LoadConfig()

	database.ConnectDatabase()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	// In the future I want to try chi's http rate limiter lib "https://github.com/go-chi/httprate"

	r.Get("/", handler.HomeHandler)

	port := fmt.Sprintf(":%s", config.Config.SERVER_PORT)

	log.Printf("Server is running on Port %s", port)
	http.ListenAndServe(port, r)
}
