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
	"strconv"

	"github.com/go-chi/chi/v5"
)

//	@Summary		Create new checking
//	@Description	Creates a new checking of a reservation already made in booking endpoint
//	@Tags			checking
//	@Accept			json
//	@Produce		json
//	@Param			checking	body		models.Checking	true	"checking"
//	@Success		200			{string}	string			"ok"
//	@Failure		500			{string}	string			"error"
//	@Router			/checking [post]
func CreateChecking(w http.ResponseWriter, r *http.Request) {
	var checking models.Checking
	err := json.NewDecoder(r.Body).Decode(&checking)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	if err = checking.Validated(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}
	isCheckingAlreadyDone, _ := usecase.IsCheckingAlreadyDone(*checking.BookingId)
	if isCheckingAlreadyDone {
		log.Printf("Esse checking já foi feito! Impossivel fazer novamente!")
		response.ResponseJson(w, http.StatusConflict, "Esse checking já foi feito! Impossivel fazer novamente!")
		return
	}
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
		response.ResponseJson(w, http.StatusInternalServerError, "Erro com a conexão com banco de dados. Por favor, tente mais tarde.")
	}
	repo := repository.ConnectionRepository(conn)
	id, err := repo.InsertChecking(checking)
	if err != nil {
		log.Println(err)
		response.ResponseJson(w, http.StatusInternalServerError, "Erro ao tentar fazer checking. Por favor, tente mais tarde.")
	}
	repo.UpdateStatus("checking", *checking.BookingId)
	var resp map[string]any
	if err != nil {
		log.Println(err)
		response.ResponseJson(w, http.StatusInternalServerError, "Erro ao tentar fazer Update. Por favor, verifique as informações e tente mais tarde.")
	} else {
		resp = map[string]any{
			"message": fmt.Sprintf("Checking %d efetuado com sucesso!", id),
		}
		response.ResponseJson(w, http.StatusCreated, resp)
	}
	defer conn.Close()
}

//	@Summary		Get list of checking
//	@Description	Returns a list of all checkouts ever made
//
//	@Tags			checking
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.CheckingComplete
//	@Failure		500	{string}	string	"error"
//	@Router			/checking [get]
func ListCheckings(w http.ResponseWriter, r *http.Request) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
		response.ResponseJson(w, http.StatusInternalServerError, "Erro com a conexão com banco de dados. Por favor, tente mais tarde.")
	}
	defer conn.Close()

	repo := repository.ConnectionRepository(conn)
	checkings, err := repo.GetAllCheckings()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	response.ResponseJson(w, http.StatusOK, checkings)
}

//	@Summary		Create new checkout
//	@Description	Enter the day and time of checkout. It can only be effective after checking.
//	@Tags			checkout
//	@Accept			json
//	@Produce		json
//	@Param			checking	query		string			true	"2023-05-20 20:00:00"
//	@Param			checkout	body		models.Checkout	true	"checkout"
//	@Success		200			{string}	string			"ok"
//	@Failure		500			{string}	string			"error"
//	@Router			/checkout [patch]
func Checkout(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	var checkout models.Checkout
	err = json.NewDecoder(r.Body).Decode(&checkout)
	if err != nil {
		log.Printf("Erro ao fazer parse do de data: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	isCheckoutAlreadyDone, err := usecase.IsCheckoutAlreadyDone(int64(id))
	if isCheckoutAlreadyDone {
		log.Printf("Esse checkout já foi feito! Impossivel fazer novamente!")
		response.ResponseJson(w, http.StatusConflict, "Esse checkout já foi feito! Impossivel fazer novamente!")
		return
	}
	if err != nil {
		log.Print(err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	repo := repository.ConnectionRepository(conn)
	rows, err := repo.UpdateCheckout(int64(id), checkout)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	checking, err := repo.GetChecking(int64(id))
	repo.UpdateStatus("checkout", *checking.BookingId)
	usecase.GenerateBill(int64(id))
	if rows > 1 {
		log.Printf("Erros: foram atualizadas %d registros", rows)
	}
	resp := map[string]any{
		"message": "Checkout realizado com sucesso!",
	}
	response.ResponseJson(w, http.StatusOK, resp)
	defer conn.Close()
}
