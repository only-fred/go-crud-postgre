package control

import (
	"database/sql"
	"fmt"
	"log"
)

func Create(db *sql.DB) {
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

func Read(db *sql.DB) {
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

func Update(db *sql.DB) {
	result, err := db.Exec("UPDATE tbl_user SET username = $1, password = $2 WHERE id = $3", "USERNAME_UPDATED_HERE", "PASSWORD_UPDATED_HERE", 1)
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

func Delete(db *sql.DB) {
	result, err := db.Exec("DELETE FROM tbl_user WHERE id = $1", 1)
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
