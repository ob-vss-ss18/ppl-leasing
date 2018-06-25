package leasing

import (
	"database/sql"
	"os"
	"log"
	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
}

