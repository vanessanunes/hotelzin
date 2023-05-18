package models

type Checking struct {
	ID               int64  `json:"id"`
	BookingId        int64  `json:"booking_id"`
	CheckingDatetime string `json:"checking_datetime,omitempty"`
	CheckoutDatetime string `json:"checkout_datetime,omitempty"`
}
