package repository

import (
	"log"
	"serasa-hotel/db"
	"serasa-hotel/models"
)

func InsertRoom(room models.Room) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO room (room_number, description) VALUES($1, $2) RETURNING id`

	err = conn.QueryRow(sql, room.RoomNumber, room.Description).Scan(&id)

	return
}

func GetRoom(id int64) (room models.Room, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM room where id=$1`, id)

	err = row.Scan(&room.ID, &room.RoomNumber, &room.Description)

	return
}

func GetAllRoom() (rooms []models.Room, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM room`)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var room models.Room
		err = rows.Scan(&room.ID, &room.RoomNumber, &room.Description)
		if err != nil {
			log.Println(err)
		}
		rooms = append(rooms, room)
	}

	return
}

func UpdateRoom(id int64, room models.Room) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE room SET room_number=$1, description=$2 WHERE id=$3`, room.RoomNumber, room.Description, id)
	if err != nil {
		log.Println(err)
	}
	return res.RowsAffected()
}

func DeleteRoom(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM room WHERE id = $1`, id)
	if err != nil {
		log.Println(err)
	}

	return res.RowsAffected()
}
