package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wandirahman13/bookings/pkg/config"
	"github.com/wandirahman13/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// use middleware
	mux.Use(middleware.Recoverer)
	mux.Use(CSRFTokenNoSurf)
	mux.Use(SessionLoad)

	// get handlers
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// serve static file
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
