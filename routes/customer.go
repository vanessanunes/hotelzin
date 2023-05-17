package routes

import (
	"serasa-hotel/handlers"

	"github.com/go-chi/chi/v5"
)

func routesCustomers(r chi.Router) {
	r.Get("/", handlers.ListCustomer)
	r.Get("/{id}", handlers.GetCustomer)
	r.Put("/{id}", handlers.UpdateCustomer)
	r.Post("/", handlers.CreateCustomer)
	r.Delete("/{id}", handlers.DeleteCustomer)
}
