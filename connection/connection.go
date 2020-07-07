package connection

import (
	"database/sql"
	"fmt"
	"log"
)

func Connection() *sql.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "<USER>"
		password = "<PASSWORD>"
		dbname   = "<DBNAME>"
	)

	dataSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Connected!")

	return db
}
