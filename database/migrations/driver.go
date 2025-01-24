package migrations

import (
	"6ar8nas/test-app/database"
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type DriverDB struct {
	db *sql.DB
}

func NewDriverDB(database *sql.DB) *DriverDB {
	return &DriverDB{db: database}
}

func (d *DriverDB) Migrate() error {
	instance, err := postgres.WithInstance(d.db, &postgres.Config{})
	if err != nil {
		return err
	}

	mig, err := migrate.NewWithDatabaseInstance(database.Env.MigrationsPath, database.Env.DriverName, instance)
	if err != nil {
		return err
	}

	if err = mig.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
