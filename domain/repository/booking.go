package repository

import (
	"fmt"
	"log"
	"serasa-hotel/domain/utils"
	"serasa-hotel/models"
)

func (repo Connection) InsertBooking(booking models.Booking) (id int64, err error) {
	sql := `INSERT INTO booking (customer_id, room_id, start_datetime, end_datetime, status, parking) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err = repo.db.QueryRow(sql, booking.CustomerID, booking.RoomID, booking.StartDatetime, booking.EndDatetime, booking.Status, booking.Parking).Scan(&id)
	if err != nil {
		log.Print(err)
	}
	return
}

func (repo Connection) GetBooking(booking_id int64) (booking models.Booking) {
	sql := `SELECT * FROM booking WHERE id = $1`
	row := repo.db.QueryRow(sql, booking_id)
	row.Scan(&booking.ID, &booking.CustomerID, &booking.RoomID, &booking.StartDatetime, &booking.EndDatetime, &booking.Status, &booking.Parking)
	return
}

func (repo Connection) GetAllBooking(bookingParams utils.BookingParams) (bookings []models.BookingAllInformations, err error) {
	sql := `
		SELECT b.id, b.start_datetime, b.end_datetime, b.status, b.parking,
		c.id, c.name, c.document, c.phone_number, c.email,
		r.id, r.room_number, r.description, ch.id, ch.checking_datetime,
		ch.checkout_datetime
		FROM booking b 
		inner join customer c on c.id = b.customer_id
		inner join room r on r.id = b.room_id
		left join checkin ch on ch.booking_id = b.id
	`
	if bookingParams.Status != "" {
		whereToSelect := fmt.Sprintf(" where status = '%s' ", bookingParams.Status)
		sql += whereToSelect
	}
	rows, err := repo.db.Query(sql)
	if err != nil {
		log.Println(err)
		return bookings, err
	}
	for rows.Next() {
		var booking models.BookingAllInformations
		err = rows.Scan(
			&booking.BookindID, &booking.BookingStartDatetime, &booking.BookingEndDatetime, &booking.Status, &booking.Parking,
			&booking.CustomerID, &booking.NameCustomer, &booking.Document, &booking.PhoneNumber, &booking.Email,
			&booking.RoomID, &booking.RoomNumber, &booking.Description, &booking.CheckingID, &booking.CheckingDatetime,
			&booking.CheckoutDatetime,
		)
		if err != nil {
			log.Println(err)
		}
		bookings = append(bookings, booking)
	}
	return
}

func (repo Connection) GetBookingByRoom(roomID int32, DateStart string, DateEnd string) (booking models.Booking, err error) {
	sql := `SELECT * FROM booking WHERE status not in ('canceled') AND (start_datetime <= $1 AND end_datetime >= $2) AND room_id=$3`
	row := repo.db.QueryRow(sql, DateEnd, DateStart, roomID)
	err = row.Scan(&booking.ID, &booking.CustomerID, &booking.RoomID, &booking.StartDatetime, &booking.EndDatetime, &booking.Status, &booking.Parking)
	if err != nil {
		err = row.Scan()
	}
	return
}

func (repo Connection) GetBookingByCustomer(customer_id int64) (bookings []models.Booking, err error) {
	sql := `SELECT * FROM booking WHERE customer_id = $1`
	rows, err := repo.db.Query(sql, customer_id)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var booking models.Booking
		err = rows.Scan(&booking.ID, &booking.CustomerID, &booking.RoomID, &booking.StartDatetime, &booking.EndDatetime, &booking.Status, &booking.Parking)
		if err != nil {
			log.Println(err)
		}
		bookings = append(bookings, booking)
	}
	return
}

func (repo Connection) GetInfoBookingHost(customer_id int64) (hostings []models.Hosting, err error) {
	sql := `select c.id, c.checking_datetime , c.checkout_datetime , b.parking , b.status , r.room_number, b2.total_value,
	b.id, b.start_datetime, b.end_datetime
	from booking b
	left join checkin c on c.booking_id = b.id
	inner join room r on r.id = b.room_id 
	left join bill b2 on b2.booking_id = b.id
	where customer_id = $1`
	rows, err := repo.db.Query(sql, customer_id)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var hosting models.Hosting
		err = rows.Scan(&hosting.CheckingID, &hosting.Checking, &hosting.Checkout, &hosting.Parking, &hosting.Status, &hosting.RoomNumber, &hosting.Value, &hosting.BookingID, &hosting.BookedStartDatetime, &hosting.BookedEndDatetime)
		if err != nil {
			log.Println(err)
		}
		hostings = append(hostings, hosting)
	}
	return
}

func (repo Connection) UpdateStatus(status string, id int64) {
	sql := fmt.Sprintf(`UPDATE booking SET status = '%s' WHERE id = %d`, status, id)
	repo.db.QueryRow(sql)
	return
}
