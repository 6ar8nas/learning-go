package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type ConnectionPool struct {
	*sql.DB
	DatabaseName string
}

func NewConnection(driver, connectionString, database string) (*ConnectionPool, error) {
	db, err := sql.Open(driver, connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Printf("Connected to %s database.", database)
	return &ConnectionPool{
		DB:           db,
		DatabaseName: database,
	}, nil
}

func (s *ConnectionPool) Close() error {
	log.Printf("Disconnected from %s database.", s.DatabaseName)
	return s.DB.Close()
}
