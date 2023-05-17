package db

import (
	"database/sql"
	"fmt"
	"log"
	"serasa-hotel/configs"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open(conf.Driver, sc)
	if err != nil {
		log.Println(err)
	}

	if err = conn.Ping(); err != nil {
		log.Println(err)
		conn.Close()
	}

	return conn, err

}
