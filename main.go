package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db := connection()

	fmt.Print("\n\n[1] C\n[2] R\n[3] U\n[4] D\n->")
	option := 0
	fmt.Scanf("%d", &option)

	switch option {
	case 1:
		create(db)

		db.Close()
		break

	case 2:
		read(db)

		db.Close()
		break

	case 3:
		update(db)

		db.Close()
		break

	case 4:
		delete(db)

		db.Close()
		break
	}
}

func connection() *sql.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "<USER_HERE>"
		password = "<PASSWORD_HERE>"
		dbname   = "<DB_HERE>"
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

func create(db *sql.DB) {
	result, err := db.Exec("INSERT INTO tbl_user (username, password) VALUES ($1, $2)", "USERNAME_HERE", "PASSWORD_HERE")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
	}

	fmt.Print("\nSuccessfully inserted!")
}

func read(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM tbl_user")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			id       int
			username string
			password string
		)

		if err := rows.Scan(&id, &username, &password); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\n||ID:  %d ||Username: %s ||Password: %s ||", id, username, password)
	}
}

func update(db *sql.DB) {
	result, err := db.Exec("UPDATE tbl_user SET username = $1, password = $2 WHERE id = $3", "USERNAME_HERE", "PASSWORD_HERE", 666)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatalf("Expected to affect 1 row, affected %d", rows)
	}

	fmt.Print("\nSuccessfully updated!")
}

func delete(db *sql.DB) {
	result, err := db.Exec("DELETE FROM tbl_user WHERE id = $1", 666)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatalf("Expected to affect 1 row, affected %d", rows)
	}

	fmt.Print("Sucessfully deleted!")
}
