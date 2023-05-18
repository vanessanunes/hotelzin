package repository

import (
	"log"
	"serasa-hotel/models"
)

func (repo Connection) InsertBooking(booking models.Booking) (id int64, err error) {
	sql := `INSERT INTO booking (customer_id, room_id, start_datetime, end_datetime, status, parking) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err = repo.db.QueryRow(sql, booking.CustomerID, booking.RoomID, booking.StartDatetime, booking.EndDatetime, booking.Status).Scan(&id)
	if err != nil {
		log.Print(err)
	}
	defer repo.db.Close()
	return
}

func (repo Connection) GetBooking(booking_id int) (booking models.Booking, err error) {
	sql := `SELECT * FROM booking WHERE id = $1`
	row := repo.db.QueryRow(sql, booking_id)
	defer repo.db.Close()

	err = row.Scan(&booking.ID, &booking.CustomerID, &booking.RoomID, &booking.StartDatetime, &booking.EndDatetime, &booking.Status, &booking.Parking)
	if err != nil {
		log.Println(err)
	}
	return

}
