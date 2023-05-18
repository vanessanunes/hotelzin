package models

type Bill struct {
	ID         int64   `json:"id"`
	BookingId  int64   `json:"booking_id"`
	ExtraHour  int32   `json:"extra_hour"`
	TotalValue float32 `json:"total_value"`
}
