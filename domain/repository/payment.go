package repository

import (
	"log"
	"serasa-hotel/models"
)

func (repo Connection) InsertPayment(payment models.Payment) (id int64, err error) {
	sql := `INSERT INTO payment (bill_id, total_value, type_payment, installments) VALUES ($1, $2, $3, $4) RETURNING id`

	err = repo.db.QueryRow(sql, payment.BillID, payment.TotalValue, payment.TypePayment, payment.Installments).Scan(&id)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return
}
