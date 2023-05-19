package usecase

import (
	"log"
	"serasa-hotel/db"
	"serasa-hotel/domain/repository"
	"serasa-hotel/models"
)

func CheckBookingAvailable(RoomID int32, DateStart string, DateEnd string) (booking models.Booking, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	repo := repository.ConnectionRepository(conn)
	booking, err = repo.GetBookingByRoom(RoomID, DateStart, DateEnd)
	return
}
