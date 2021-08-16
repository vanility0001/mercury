package database

import (
	"database/sql"
	"log"
)

func ConnectToDb(databaseURL string) *sql.DB {
	con, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalln()
	}

	return con
}
