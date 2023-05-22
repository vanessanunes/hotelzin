package main

import (
	"fmt"
	"log"
	"net/http"
	"serasa-hotel/configs"
	"serasa-hotel/routes"

	"github.com/go-chi/chi/v5"
)

//	@title			API de Hospedagem
//	@version		1.0.0
//	@contact.name	Vanessa Nunes
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:9000
//	@BasePath		/
func main() {
	err := configs.Load()

	if err != nil {
		log.Println(err)
	}

	r := chi.NewRouter()
	routes.GetRoutes(r)

	log.Println("listening...")
	log.Println(fmt.Sprintf(":%s", configs.GetServerPort()))

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
