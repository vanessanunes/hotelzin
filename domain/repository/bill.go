package repository

import (
	"log"
	"serasa-hotel/models"
)

func (repo Connection) InsertBill(bill models.Bill) (id int64, err error) {
	sql := `INSERT INTO bill (booking_id, extra_hour, total_value) VALUES ($1, $2, $3) RETURNING id`
	err = repo.db.QueryRow(sql, bill.BookingId, bill.ExtraHour, bill.TotalValue).Scan(&id)
	if err != nil {
		log.Print(err)
	}
	defer repo.db.Close()
	return
}

func (repo Connection) GetAllBills() (bills []models.Bill, err error) {
	sql := `SELECT * FROM bill`
	rows, err := repo.db.Query(sql)
	if err != nil {
		log.Println(err)
		return bills, err
	}
	var bill models.Bill
	for rows.Next() {
		err = rows.Scan(&bill.ID, &bill.BookingId, &bill.TotalValue, &bill.ExtraHour)
		if err != nil {
			log.Printf("Erro ao pegar clientes: %v", err)
		}
		bills = append(bills, bill)
	}
	return
}

func (repo Connection) GetBill(bookingId int64) (bill models.Bill, err error) {
	sql := `SELECT * FROM bill where booking_id = $1`
	row := repo.db.QueryRow(sql, bookingId)
	err = row.Scan(&bill.ID, &bill.BookingId, &bill.TotalValue, &bill.ExtraHour)
	if err != nil {
		log.Println(err)
		return bill, err
	}
	return
}
