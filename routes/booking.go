package routes

import (
	"serasa-hotel/handlers"

	"github.com/go-chi/chi/v5"
)

func routesBooking(r chi.Router) {
	r.Post("/", handlers.CreateBooking)
	// r.Get("/", handlers.ListCheckings)
	// r.Get("/{id}", handlers.GetCustomer)
	// r.Put("/{id}", handlers.UpdateChecking)
	// r.Post("/", handlers.CreateCustomer)
	// r.Delete("/{id}", handlers.DeleteCustomer)
}
