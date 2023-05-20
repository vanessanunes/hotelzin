package routes

import (
	"serasa-hotel/handlers"

	"github.com/go-chi/chi/v5"
)

func routesCheckout(r chi.Router) {
	r.Patch("/{id}", handlers.Checkout)
}
