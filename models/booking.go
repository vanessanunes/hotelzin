package models

import (
	"errors"
	"fmt"
	"regexp"
)

type Booking struct {
	ID            int64  `json:"booking_id,omitempty" swaggerignore:"true"`
	CustomerID    int64  `json:"customer_id,omitemptyd"`
	RoomID        int64  `json:"room_id"`
	StartDatetime string `json:"start_datetime" example:"2023/05/20"`
	EndDatetime   string `json:"end_datetime" example:"2023/05/20"`
	Status        string `json:"status" collection:"reserved,checking,checkout,canceled"`
	Parking       bool   `json:"parking"`
}

type BookingAllInformations struct {
	BookindID            *int64  `json:"booking_id"`
	CustomerID           *int64  `json:"customer_id"`
	BookingStartDatetime *string `json:"booking_start_datetime"`
	BookingEndDatetime   *string `json:"booking_end_datetime"`
	Status               *string `json:"status"`
	Parking              *bool   `json:"parking"`
	NameCustomer         *string `json:"customer_name"`
	Document             *string `json:"customer_document"`
	PhoneNumber          *string `json:"phone_number"`
	Email                *string `json:"email"`
	RoomID               *int64  `json:"room_id"`
	RoomNumber           *int32  `json:"room_number"`
	Description          *string `json:"room_description"`
	CheckingID           *int64  `json:"checking_id"`
	CheckingDatetime     *string `json:"checking_datetime"`
	CheckoutDatetime     *string `json:"checkout_datetime"`
}

type Hosting struct {
	CheckingID          *int64   `json:"checking_id"`
	Checking            *string  `json:"checking"`
	Checkout            *string  `json:"checkout"`
	Parking             *bool    `json:"parking"`
	Status              *string  `json:"status"`
	RoomNumber          *string  `json:"room_number"`
	BookingID           *int64   `json:"booking_id"`
	BookedStartDatetime *string  `json:"booked_start_datetime"`
	BookedEndDatetime   *string  `json:"booked_end_datetime"`
	Value               *float32 `json:"value"`
}

func (b *Booking) Validated() error {
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

	if b.CustomerID == 0 {
		return errors.New("Inclua o ID do cliente para realizar a reserva")
	}
	if b.RoomID == 0 {
		return errors.New("Inclua o ID do quarto para realizar a reserva")
	}
	if b.StartDatetime == "" {
		return errors.New("Uma data deve ser incluída para inicio da reserva")
	}
	if b.StartDatetime != "" {
		dateFind := re.FindStringSubmatch(b.StartDatetime)
		b.StartDatetime = fmt.Sprintf("%s %s", dateFind[0], "16:30")
	}
	if b.EndDatetime == "" {
		return errors.New("Uma data deve ser incluída para final da reserva")
	}
	if b.EndDatetime != "" {
		dateFind := re.FindStringSubmatch(b.EndDatetime)
		b.EndDatetime = fmt.Sprintf("%s %s", dateFind[0], "16:30")
	}
	if b.Status == "" {
		b.Status = "reserved"
	}
	return nil
}
