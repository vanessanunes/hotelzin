package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"serasa-hotel/db"
	"serasa-hotel/domain/repository"
	"serasa-hotel/models"
	"serasa-hotel/response"
)

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment

	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	if err = payment.Validated(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	repo := repository.ConnectionRepository(conn)
	lastId, err := repo.InsertPayment(payment)
	var resp map[string]any
	if err != nil {
		log.Println(err)
		response.ResponseJson(w, http.StatusCreated, "Erro ao realizar pagamentos. Por favor, verifique os dados cadastrados!")
	} else {
		resp = map[string]any{
			"message": fmt.Sprintf("Pagemento efetuado com sucesso! ID: %d", lastId),
		}
		response.ResponseJson(w, http.StatusCreated, resp)
	}
}
