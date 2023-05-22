package routes

import (
	_ "serasa-hotel/docs"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func GetRoutes(r chi.Router) {
	r.Route("/customer", routesCustomers)
	r.Route("/booking", routesBooking)
	r.Route("/checking", routesChecking)
	r.Route("/checkout", routesCheckout)
	r.Route("/payment", routesPayment)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:9000/swagger/doc.json"), //The url pointing to API definition
	))

}
