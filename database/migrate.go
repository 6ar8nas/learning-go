package database

import (
	"6ar8nas/test-app/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (d *ConnectionPool) Migrate() error {
	instance, err := postgres.WithInstance(d.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	mig, err := migrate.NewWithDatabaseInstance(config.MigrationsPath, config.Driver, instance)
	if err != nil {
		return err
	}

	if err = mig.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
