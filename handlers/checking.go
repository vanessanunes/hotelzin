package handlers

import (
	"encoding/json"
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
	}
	defer conn.Close()
	repo := repository.ConnectionRepository(conn)
	id, err := repo.InsertChecking(checking)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusInternalServerError, err)
	}
	repo.UpdateStatus("checking", id)
	var resp map[string]any
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusInternalServerError, err)
	} else {
		resp = map[string]any{
			"message": "Checking efetuado com sucesso!",
		}
		response.ResponseJson(w, http.StatusCreated, resp)
	}
}

func ListCheckings(w http.ResponseWriter, r *http.Request) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
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
