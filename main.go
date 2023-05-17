package main

import (
	"fmt"
	"log"
	"net/http"
	"serasa-hotel/configs"
	"serasa-hotel/routes"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()

	if err != nil {
		log.Println(err)
	}

	r := chi.NewRouter()
	routes.GetRoutes(r)

	log.Println("listening...")

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
