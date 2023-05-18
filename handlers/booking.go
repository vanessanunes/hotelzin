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

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	repo := repository.ConnectionRepository(conn)
	lastID, err := repo.InsertBooking(booking)

	var resp map[string]any

	if err != nil {
		response.ResponseError(w, http.StatusCreated, err)
	} else {
		resp = map[string]any{
			"message": fmt.Sprintf("Reserva gerada com sucesso! Número de reserva: %d", lastID),
		}
		response.ResponseJson(w, http.StatusCreated, resp)
	}
}
