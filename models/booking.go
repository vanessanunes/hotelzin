package models

import "errors"

type Booking struct {
	ID            int64  `json:"id"`
	CustomerID    int64  `json:"customer_id"`
	RoomID        int64  `json:"room_id"`
	StartDatetime string `json:"start_datetime"`
	EndDatetime   string `json:"end_datetime"`
	Status        string `json:"status"`
	Parking       bool   `json:"parking"`
}

func (b *Booking) Validated() error {
	if b.CustomerID == 0 {
		return errors.New("Inclua o ID do cliente para realizar a reserva")
	}
	if b.RoomID == 0 {
		return errors.New("Inclua o ID do quarto para realizar a reserva")
	}
	if b.StartDatetime == "" {
		return errors.New("Uma data deve ser incluída para inicio da reserva")
	}
	if b.EndDatetime == "" {
		return errors.New("Uma data deve ser incluída para final da reserva")
	}
	if b.Status == "" {
		b.Status = "reserved"
	}
	return nil
}
