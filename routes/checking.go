package routes

import (
	"serasa-hotel/handlers"

	"github.com/go-chi/chi/v5"
)

func routesChecking(r chi.Router) {
	r.Post("/", handlers.CreateChecking)
	r.Get("/", handlers.ListCheckings)
	// r.Get("/{id}", handlers.GetCustomer)
	r.Put("/{id}", handlers.UpdateChecking)    // pensar melhor
	r.Put("/checkout/{id}", handlers.Checkout) // pensar melhor

	// r.Post("/", handlers.CreateCustomer)
	// r.Delete("/{id}", handlers.DeleteCustomer)
}
