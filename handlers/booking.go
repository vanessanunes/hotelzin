package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"serasa-hotel/db"
	"serasa-hotel/domain/repository"
	"serasa-hotel/domain/usecase"
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
	if err = booking.Validated(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	booked, err := usecase.CheckBookingAvailable(int32(booking.RoomID), booking.StartDatetime, booking.EndDatetime)
	if err != nil {
		log.Printf("Quarto %d disponivel!", booking.RoomID)
	}

	if booked != (models.Booking{}) {
		log.Printf("Desculpe, o quarto para essa data já está ocupado. Por favor, tente outra data!")
		response.ResponseJson(w, http.StatusConflict, "Desculpe, o quarto para essa data já está ocupado. Por favor, tente outra data!")
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
