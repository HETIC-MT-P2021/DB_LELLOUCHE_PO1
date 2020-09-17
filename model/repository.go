package model

import "database/sql"

//Repository struct to store db connection
type Repository struct {
	Conn *sql.DB
}
