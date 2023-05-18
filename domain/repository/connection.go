package repository

import "database/sql"

type Connection struct {
	db *sql.DB
}

func ConnectionRepository(db *sql.DB) *Connection {
	return &Connection{db}
}
