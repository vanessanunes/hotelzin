package routes

import (
	"github.com/go-chi/chi/v5"
)

func GetRoutes(r chi.Router) {
	r.Route("/customer", routesCustomers)

}
