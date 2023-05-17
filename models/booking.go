package models

import "time"

const (
	Booked   = 0
	Checkin  = 1
	Checkout = 2
)

type Booking struct {
	ID               int64     `json:"id"`
	CustomerID       int64     `json:"customer_id"`
	RoomID           int64     `json:"room_id"`
	StartedDatetime  time.Time `json:"started_datetime"`
	FinishedDatetime time.Time `json:"finished_datetime"`
	Status           bool      `json:"status"`
}
