package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"serasa-hotel/db"
	"serasa-hotel/domain/repository"
	"serasa-hotel/domain/utils"
	"serasa-hotel/models"
	"serasa-hotel/response"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// @Summary		Get list of customer
// @Description	Get list of customer
// @Tags			customer
// @Accept			json
// @Param			id			query	integer	true	"1"
// @Param			phone		query	integer	true	"1165556989"
// @Param			document	query	integer	true	"40140154588"
// @Produce		json
// @Success		200	{array}		models.Customer
// @Failure		500	{string}	string	"error"
// @Router			/customer [get]
func ListCustomer(w http.ResponseWriter, r *http.Request) {
	var params utils.CustomerParams
	err := utils.Decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Println("Error in GET parameters : ", err)
	}
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	repo := repository.ConnectionRepository(conn)
	customers, err := repo.GetAllCustomer(params)
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	response.ResponseJson(w, http.StatusOK, customers)
}

// @Summary		Get a customer
// @Description	Get a customer
// @Tags			customer
// @Accept			json
// @Param			id	query	integer	true	"1"
// @Produce		json
// @Success		200	{object}	models.CustomerWithHosting
// @Failure		500	{string}	string	"error"
// @Router			/customer/{id} [get]
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	repo := repository.ConnectionRepository(conn)
	customer, err := repo.GetCustomer(int64(id))
	if err != nil {
		log.Printf("Erro ao obter registros: %v. ID: %d", err, id)
		resp := map[string]any{
			"message": fmt.Sprintf("Registro de cliente não encontrado: %d", id),
		}
		response.ResponseJson(w, http.StatusInternalServerError, resp)
		return
	}
	bookings, err := repo.GetInfoBookingHost(int64(id))
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	defer conn.Close()
	hosting := models.CustomerWithHosting{Customer: customer, Hostings: bookings, TotalValue: 0.0}
	response.ResponseJson(w, http.StatusOK, hosting)
}

// @Summary		Update a customer
// @Description	Update a customer
// @Tags			customer
// @Accept			json
// @Param			customer	body	models.Customer	true	"customer"
// @Produce		json
// @Success		200	{array}		models.Customer
// @Failure		500	{string}	string	"error"
// @Router			/customer/ [put]
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

// @Summary		Create new Customer
// @Description	Create new Customer
// @Tags			customer
// @Accept			json
// @Produce		json
// @Param			customer	body		models.Customer	true	"customer"
// @Success		200			{string}	string			"ok"
// @Failure		500			{string}	string			"error"
// @Router			/customer [post]
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
		response.ResponseJson(w, http.StatusInternalServerError, "Erro com a conexão com banco de dados. Por favor, tente mais tarde.")
		return
	}
	defer conn.Close()

	repo := repository.ConnectionRepository(conn)
	lastId, err := repo.InsertCustomer(customer)

	var resp map[string]any

	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusInternalServerError, err)
	} else {
		resp = map[string]any{
			"message": fmt.Sprintf("Cliente inserido com sucesso! ID: %d", lastId),
		}
		response.ResponseJson(w, http.StatusCreated, resp)
	}
}

// @Summary		Delete a customer
// @Description	Delete a customer
// @Tags			customer
// @Accept			json
// @Param			id	query	integer	true	"1"
// @Produce		json
// @Success		200	{string}	string	"ok"
// @Failure		500	{string}	string	"error"
// @Router			/customer/{id} [delete]
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
