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
