package models

type Bill struct {
	ID        int64   `json:"id"`
	BookingId int64   `json:"booking_id"`
	ExtraHour int64   `json:"extra_hour"`
	Discount  float32 `json:"discount"`
}
