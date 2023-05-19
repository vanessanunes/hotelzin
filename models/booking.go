package models

import (
	"errors"
	"fmt"
	"regexp"
)

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
