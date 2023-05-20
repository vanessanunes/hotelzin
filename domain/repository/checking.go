package repository

import (
	"log"
	"serasa-hotel/models"
)

func (repo Connection) InsertChecking(checking models.Checking) (id int64, err error) {
	sql := `INSERT INTO checkin (booking_id, checking_datetime) VALUES ($1, $2) RETURNING id`
	err = repo.db.QueryRow(sql, checking.BookingId, checking.CheckingDatetime).Scan(&id)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	defer repo.db.Close()
	return
}

func (repo Connection) UpdateCheckout(id int64, checkout models.Checkout) (int64, error) {
	res, err := repo.db.Exec(`UPDATE checkin SET checkout_datetime=$1 WHERE id=$2`, checkout.CheckoutDatetime, id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return res.RowsAffected()
}

func (repo Connection) GetAllCheckings() (checkings []models.CheckingComplete, err error) {
	rows, err := repo.db.Query(`SELECT c.id, c.booking_id, c.checking_datetime, c.checkout_datetime, b.status FROM checkin c inner join booking b on b.id = c.booking_id`)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var checking models.CheckingComplete
		err = rows.Scan(&checking.ID, &checking.BookingId, &checking.CheckingDatetime, &checking.CheckoutDatetime, &checking.Status)
		if err != nil {
			log.Println(err)
		}
		checkings = append(checkings, checking)
	}

	return
}

func (repo Connection) GetChecking(id int64) (checking models.CheckingComplete, err error) {
	row := repo.db.QueryRow(`SELECT * FROM checkin where id=$1`, id)
	err = row.Scan(&checking.ID, &checking.BookingId, &checking.CheckingDatetime, &checking.CheckoutDatetime)
	if err != nil {
		log.Println(err)
		return checking, err
	}
	return
}

func (repo Connection) GetCheckingByBooking(id int64) (checking models.CheckingComplete, err error) {
	row := repo.db.QueryRow(`SELECT * FROM checkin where booking_id=$1`, id)
	defer repo.db.Close()
	err = row.Scan(&checking.ID, &checking.BookingId, &checking.CheckingDatetime, &checking.CheckoutDatetime)
	if err != nil {
		log.Println(err)
		return checking, err
	}
	return
}

func (repo Connection) IsCheckoutIsDone(id int64) (bool, error) {
	getChecking, err := repo.GetChecking(id)
	if err != nil {
		log.Println(err)
		return false, err
	}
	if getChecking.CheckoutDatetime != nil {
		return true, err
	}
	return false, err

}

func (repo Connection) IsCheckingIsDone(bookingId int64) (bool, error) {
	getChecking, err := repo.GetCheckingByBooking(bookingId)
	if err != nil {
		log.Println(err)
		return false, err
	}
	if getChecking.CheckingDatetime != nil {
		return true, err
	}
	return false, err

}
