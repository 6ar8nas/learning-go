package database

import (
	"database/sql"
	"log"

	"github.com/6ar8nas/learning-go/server/config"
	_ "github.com/lib/pq"
)

type ConnectionPool struct {
	*sql.DB
}

var dbConn *ConnectionPool

func NewService() (*ConnectionPool, error) {
	if dbConn != nil {
		return dbConn, nil
	}

	db, err := sql.Open(config.Driver, config.ConnectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	dbConn = &ConnectionPool{
		DB: db,
	}
	log.Printf("Connected to %s database.", config.Database)
	return dbConn, nil
}

func (s *ConnectionPool) Close() error {
	log.Printf("Disconnected from %s database.", config.Database)
	return s.DB.Close()
}
