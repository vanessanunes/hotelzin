package repository

import (
	"log"
	"serasa-hotel/db"
	"serasa-hotel/models"
)

func (repo Connection) InsertChecking(checking models.Checking) (id int64, err error) {
	sql := `INSERT INTO checkin (booking_id, checking_datetime) VALUES ($1, $2) RETURNING id`

	err = repo.db.QueryRow(sql, checking.BookingId, checking.CheckingDatetime).Scan(&id)
	if err != nil {
		log.Print(err)
	}
	defer repo.db.Close()
	return
}

func (repo Connection) UpdateChecking(id int64, checking models.Checking) (int64, error) {
	res, err := repo.db.Exec(`UPDATE checkin SET booking_id=$1, checking_datetime=$2 WHERE id=$3`, checking.BookingId, checking.CheckingDatetime, id)
	if err != nil {
		log.Println(err)
	}

	defer repo.db.Close()
	return res.RowsAffected()
}

func (repo Connection) UpdateCheckout(id int64, checking models.Checking) (int64, error) {
	res, err := repo.db.Exec(`UPDATE checkin SET booking_id=$1, checking_datetime=$2, checkout_datetime=$3 WHERE id=$4`, checking.BookingId, checking.CheckingDatetime, checking.CheckoutDatetime, id)
	if err != nil {
		log.Println(err)
	}

	return res.RowsAffected()
}

func (repo Connection) GetAllCheckings() (checkings []models.Checking, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM checkin`)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var checking models.Checking
		err = rows.Scan(&checking.ID, &checking.BookingId, &checking.CheckingDatetime, &checking.CheckoutDatetime)
		if err != nil {
			log.Println(err)
		}
		checkings = append(checkings, checking)
	}

	return
}