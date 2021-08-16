package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func ConnectToDb(databaseURL string) {
	con, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalln()
	}

	DB = con
}
