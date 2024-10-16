package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitializeDB(connectionString string) (*sql.DB, error) {
	if db == nil {
		log.Println("initializing db connection for postgres psql info")

		var err error

		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			return nil, err
		}

		// Verify connection
		err = db.Ping()
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
