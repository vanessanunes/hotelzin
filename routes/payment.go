package routes

import (
	"serasa-hotel/handlers"

	"github.com/go-chi/chi/v5"
)

func routesPayment(r chi.Router) {
	r.Post("/", handlers.CreatePayment)
}
