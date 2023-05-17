package models

type Room struct {
	ID          int64  `json:"id"`
	RoomNumber  int    `json:"room_number"`
	Description string `json:"room_description"`
}
