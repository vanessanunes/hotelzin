package usecase

import (
	"fmt"
	"log"
	"serasa-hotel/db"
	"serasa-hotel/domain/repository"
)

func IsCheckoutAlreadyDone(checkoutID int64) (bool, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	repo := repository.ConnectionRepository(conn)
	isCheckout, err := repo.IsCheckoutIsDone(checkoutID)
	if err != nil {
		fmt.Print(err)
		return false, err
	}
	return isCheckout, err
}

func IsCheckingAlreadyDone(bookingID int64) (bool, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	repo := repository.ConnectionRepository(conn)
	isCheckout, err := repo.IsCheckingIsDone(bookingID)
	if err != nil {
		fmt.Print(err)
		return false, err
	}
	return isCheckout, err
}
