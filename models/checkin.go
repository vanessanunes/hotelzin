package models

import "time"

type Checking struct {
	ID               int64     `json:"id"`
	BookingId        int64     `json:"booking_id"`
	CheckingDatetime time.Time `json:"checking_datetime"`
	CheckoutDatetime time.Time `json:"checkout_datetime"`
}
