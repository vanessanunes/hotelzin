package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"serasa-hotel/db"
	"serasa-hotel/domain/repository"
	"serasa-hotel/domain/usecase"
	"serasa-hotel/domain/utils"
	"serasa-hotel/models"
	"serasa-hotel/response"
)

// @Summary		Create new booking
// @Description	Create a new booking intent at the hotel.
// @Tags			booking
// @Accept			json
// @Produce		json
// @Param			booking	body		models.Booking	true	"booking"
// @Success		200		{string}	string			"ok"
// @Failure		500		{string}	string			"error"
// @Router			/booking [post]
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		response.ResponseJson(w, http.StatusInternalServerError, "Erro nas informações. Por favor, verifique e tente novamente.")
		return
	}
	if err = booking.Validated(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	booked, err := usecase.CheckBookingAvailable(int32(booking.RoomID), booking.StartDatetime, booking.EndDatetime)
	if booked != (models.Booking{}) {
		log.Printf("Desculpe, o quarto para essa data já está ocupado. Por favor, tente outra data!")
		resp := map[string]any{
			"message": fmt.Sprintf("Desculpe, o quarto para essa data já está ocupado. Por favor, tente outra data!"),
		}
		response.ResponseJson(w, http.StatusNotFound, resp)
		return
	}

	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
		response.ResponseJson(w, http.StatusConflict, "Erro de conexão com o banco de dados, por favor, tente mais tarde.")
	}
	defer conn.Close()
	repo := repository.ConnectionRepository(conn)
	customer, err := repo.GetCustomer(booking.CustomerID)
	if err != nil {
		resp := map[string]any{
			"message":  fmt.Sprintf("ID de %d cliente não encontrado. Por favor, verifique e tente novamente.", booking.CustomerID),
			"customer": customer,
		}
		response.ResponseJson(w, http.StatusNotFound, resp)
		return
	}
	lastID, err := repo.InsertBooking(booking)
	if err != nil {
		resp := map[string]any{
			"message":  fmt.Sprintf("Ocorreu algum problema ao efetuar reserva. Por favor, tente novamente mais tarde."),
			"customer": customer,
		}
		response.ResponseJson(w, http.StatusNotFound, resp)
		return
	}

	var resp map[string]any

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
	} else {
		resp = map[string]any{
			"message":  fmt.Sprintf("Reserva gerada com sucesso! Número de reserva: %d", lastID),
			"customer": customer,
		}
		response.ResponseJson(w, http.StatusCreated, resp)
	}
}

// @Summary		Get list of bookings
// @Description	List all bookings at the hotel. Can use status filter.
// @Tags			booking
// @Accept			json
// @Param			status	query	string	false	"string enums"	Enums(reserved,checking,checkout,canceled)
// @Produce		json
// @Success		200	{array}		models.Booking
// @Failure		500	{string}	string	"error"
// @Router			/booking [get]
func ListBooking(w http.ResponseWriter, r *http.Request) {
	var params utils.BookingParams
	err := utils.Decoder.Decode(&params, r.URL.Query())
	conn, err := db.OpenConnection()
	if err != nil {
		log.Printf("Erro com conexão: %v", err)
		response.ResponseJson(w, http.StatusInternalServerError, "Erro com conexão com o banco de dados, por favor, tente novamente.")
		return
	}
	defer conn.Close()
	repo := repository.ConnectionRepository(conn)
	customers, err := repo.GetAllBooking(params)
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	response.ResponseJson(w, http.StatusOK, customers)
}
