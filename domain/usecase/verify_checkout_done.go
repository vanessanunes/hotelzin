package usecase

import (
	"log"
	"serasa-hotel/db"
	"serasa-hotel/domain/repository"
)

func IsCheckoutAlreadyDone(checkoutID int64) bool {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	repo := repository.ConnectionRepository(conn)
	isCheckout := repo.CheckoutIsDone(checkoutID)
	return isCheckout

}
