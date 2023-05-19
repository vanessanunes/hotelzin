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
	"strconv"

	"github.com/go-chi/chi/v5"
)

func ListCustomer(w http.ResponseWriter, r *http.Request) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	repo := repository.ConnectionRepository(conn)
	customers, err := repo.GetAllCustomer()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	response.ResponseJson(w, http.StatusOK, customers)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	repo := repository.ConnectionRepository(conn)
	customers, err := repo.GetCustomer(int64(id))
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	response.ResponseJson(w, http.StatusOK, customers)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	var customer models.Customer
	err = json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		log.Printf("Erro ao fazer parse do cliente: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	repo := repository.ConnectionRepository(conn)
	rows, err := repo.UpdateCustomer(int64(id), customer)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if rows > 1 {
		log.Printf("Erros: foram atualizadas %d registros", rows)
	}

	resp := map[string]any{
		"message": "dados atualizados com sucesso!",
	}

	response.ResponseJson(w, http.StatusOK, resp)

}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if err = customer.Validated(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	repo := repository.ConnectionRepository(conn)
	lastId, err := repo.InsertCustomer(customer)

	var resp map[string]any

	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusCreated, err)
	} else {
		resp = map[string]any{
			"message": fmt.Sprintf("Cliente inserido com sucesso! ID: %d", lastId),
		}
		response.ResponseJson(w, http.StatusCreated, resp)
	}
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()
	repo := repository.ConnectionRepository(conn)
	rows, err := repo.DeleteCustomer(int64(id))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Erros: foram removidos %d registros", rows)
	}

	resp := map[string]any{
		"message": "dados removidos com suscesso!",
	}
	response.ResponseJson(w, http.StatusOK, resp)
}
