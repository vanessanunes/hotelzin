package models

type Booking struct {
	ID            int64  `json:"id"`
	CustomerID    int64  `json:"customer_id"`
	RoomID        int64  `json:"room_id"`
	StartDatetime string `json:"start_datetime"`
	EndDatetime   string `json:"end_datetime"`
	Status        string `json:"status"`
	Parking       bool   `json:"parking"`
}
