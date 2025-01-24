package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDatabaseConnection() (*sql.DB, error) {
	db, err := sql.Open(Env.DriverName, Env.ConnectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return db, nil
}
