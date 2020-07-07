package main

import (
	"fmt"

	"./connection"
	"./control"

	_ "github.com/lib/pq"
)

func main() {
	db := connection.Connection()

	fmt.Print("\n\n[1] C\n[2] R\n[3] U\n[4] D\n->")
	option := 0
	fmt.Scanf("%d", &option)

	switch option {
	case 1:
		control.Create(db)

		db.Close()
		break

	case 2:
		control.Read(db)

		db.Close()
		break

	case 3:
		control.Update(db)

		db.Close()
		break

	case 4:
		control.Delete(db)

		db.Close()
		break
	}
}
