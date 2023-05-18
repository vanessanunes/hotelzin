package usecase

import (
	"fmt"
	"log"
	"serasa-hotel/db"
	"serasa-hotel/domain/repository"
	"serasa-hotel/models"
)

func GetChecking(checkingID int) (checking models.Checking) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	repo := repository.ConnectionRepository(conn)
	checking, err = repo.GetChecking(int64(checkingID))
	if err != nil {
		log.Fatal("Erro ao pegar informações de checking")
	}
	return checking
}

func GetBooking(bookingID int) (booking models.Booking) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	repo := repository.ConnectionRepository(conn)
	booking, err = repo.GetBooking(int(bookingID))
	if err != nil {
		log.Fatal("Erro ao pegar informações de booking")
	}
	return booking
}

func GenerateBill(checkingID int) int64 {
	checking := GetChecking(checkingID)
	booking := GetBooking(int(checking.BookingId))
	extraHour := ExtraHour(checking, booking)
	totalValue := CalculateTotal(checking, booking.Parking, extraHour)

	conn, err := db.OpenConnection()
	if err != nil {
		log.Println("Erro ao abrir conexão")
	}
	repo := repository.ConnectionRepository(conn)
	var bill models.Bill
	bill.BookingId = booking.ID
	bill.ExtraHour = extraHour
	bill.TotalValue = totalValue
	row, err := repo.InsertBill(bill)
	if err != nil {
		log.Println("Erro ao abrir conexão")
	}
	return row
}

func CalculateTotal(checking models.Checking, bookingParking bool, extraHour int32) float32 {
	totalValueParking := float32(0.0)
	if bookingParking {
		totalValueParking = ParkingTotalValue(checking.CheckingDatetime, checking.CheckoutDatetime)
		totalValueParking += CalculateExtraParkingHour(extraHour)
	}
	totalValueRoom := RoomTotalValue(checking.CheckingDatetime, checking.CheckoutDatetime)
	totalValueRoom += CalculateExtraRoomHour(extraHour)
	total := totalValueParking + totalValueRoom
	return total
}

func ExtraHour(checking models.Checking, booking models.Booking) int32 {
	_, BookingDateEnd := ConvertDBDateToDateTime(booking.StartDatetime, booking.EndDatetime)
	_, CheckingDateEnd := ConvertDBDateToDateTime(checking.CheckingDatetime, checking.CheckoutDatetime)
	diff := CheckingDateEnd.Sub(BookingDateEnd)
	diffHours := diff.Hours() / 24
	fmt.Println(diffHours)
	if diff.Hours()/24 <= 0 {
		return 0
	}
	return 1
}
