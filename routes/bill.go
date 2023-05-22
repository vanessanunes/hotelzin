package routes

import (
	"serasa-hotel/handlers"

	"github.com/go-chi/chi/v5"
)

func routesBill(r chi.Router) {
	r.Get("/", handlers.ListBill)
	r.Get("/{id}", handlers.GetBill)
}
