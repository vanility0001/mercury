package database

import (
	"database/sql"
	"log"
)

func QueryUser(con *sql.DB, name string) User {
	var user User
	err := con.QueryRow("SELECT * FROM fake_table WHERE name=$1", name).
		Scan(
			&user.Name,
			&user.Password,
			&user.Country,
			&user.UserID,
			&user.Address,
		)
	if err != nil {
		log.Fatalln(err)
	}

	return user
}
