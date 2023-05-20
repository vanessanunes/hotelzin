package routes

import (
	"serasa-hotel/handlers"

	"github.com/go-chi/chi/v5"
)

func routesChecking(r chi.Router) {
	r.Post("/", handlers.CreateChecking)
	r.Get("/", handlers.ListCheckings)
}
