package usecase

import (
	"log"
	"serasa-hotel/db"
	"serasa-hotel/domain/repository"
	"serasa-hotel/models"
)

func GetChecking(checkingID int64) (checking models.CheckingComplete, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
		return checking, err
	}
	repo := repository.ConnectionRepository(conn)
	checking, err = repo.GetChecking(int64(checkingID))
	if err != nil {
		log.Fatal("Erro ao pegar informações de checking")
	}
	return
}

func GetBooking(bookingID int64) (booking models.Booking, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
		return booking, err
	}
	repo := repository.ConnectionRepository(conn)
	booking = repo.GetBooking(bookingID)
	return
}

func GenerateBill(checkingID int64) int64 {
	checking, err := GetChecking(checkingID)
	if err != nil {
		log.Println("Erro ao pegar dados de checking")
		return 0
	}
	booking, err := GetBooking(*checking.BookingId)
	if err != nil {
		log.Println("Erro ao pegar dados de reserva")
		return 0
	}
	extraHour := ExtraHour(checking, booking)
	totalValue := CalculateTotal(checking, booking.Parking, extraHour)

	conn, err := db.OpenConnection()
	if err != nil {
		log.Println("Erro ao abrir conexão")
	}
	if err != nil {
		log.Println("Erro com conexão")
		return 0
	}
	repo := repository.ConnectionRepository(conn)
	bill := models.Bill{BookingId: &booking.ID, ExtraHour: &extraHour, TotalValue: &totalValue}
	row, err := repo.InsertBill(bill)
	if err != nil {
		log.Println("Erro ao inserir conta")
	}
	return row
}

func CalculateTotal(checking models.CheckingComplete, bookingParking bool, extraHour bool) float32 {
	totalValueParking := float32(0.0)
	if bookingParking {
		totalValueParking = ParkingTotalValue(*checking.CheckingDatetime, *checking.CheckoutDatetime)
		totalValueParking += CalculateExtraParkingHour(extraHour)
	}
	totalValueRoom := RoomTotalValue(*checking.CheckingDatetime, *checking.CheckoutDatetime)
	totalValueRoom += CalculateExtraRoomHour(extraHour)
	total := totalValueParking + totalValueRoom
	return total
}

func ExtraHour(checking models.CheckingComplete, booking models.Booking) bool {
	_, BookingDateEnd := ConvertDBDateToDateTime(booking.StartDatetime, booking.EndDatetime)
	_, CheckingDateEnd := ConvertDBDateToDateTime(*checking.CheckingDatetime, *checking.CheckoutDatetime)
	diff := CheckingDateEnd.Sub(BookingDateEnd)
	diffHours := diff.Hours() / 24
	if diffHours <= 0 {
		return false
	}
	return true
}
