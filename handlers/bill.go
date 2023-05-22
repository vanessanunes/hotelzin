package handlers

import (
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

// @Summary		Get list of bills
// @Description	The bill is automatically generated once checkout is completed. This endpoint return all registry.
// @Tags			bill
// @Accept			json
// @Produce		json
//
// @Param			booking_id	query		integer	true	"1"
// @Success		200			{array}		models.Bill
// @Failure		500			{string}	string	"error"
// @Router			/bill [get]
func ListBill(w http.ResponseWriter, r *http.Request) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	repo := repository.ConnectionRepository(conn)
	bills, err := repo.GetAllBills()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	response.ResponseJson(w, http.StatusOK, bills)
}

// @Summary		Get a bill
// @Description	Get a bill
// @Tags			bill
// @Accept			json
// @Param			id	query	integer	true	"1"
// @Produce		json
// @Success		200	{object}	models.BillWithPayment
// @Failure		500	{string}	string	"error"
// @Router			/bill/{id} [get]
func GetBill(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	repo := repository.ConnectionRepository(conn)
	bill, err := repo.GetBill(int64(id))
	if err != nil {
		log.Printf("Erro ao obter registros: %v. ID: %d", err, id)
		resp := map[string]any{
			"message": fmt.Sprintf("Registro de conta não encontrado: %d", id),
		}
		response.ResponseJson(w, http.StatusInternalServerError, resp)
		return
	}
	payment, err := repo.GetPaymentByBill(*bill.ID)
	if err != nil {
		log.Print(err)
	}
	billWithPayment := models.BillWithPayment{Bill: bill, Payment: payment}
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}
	defer conn.Close()
	response.ResponseJson(w, http.StatusOK, billWithPayment)
}
