package database

import (
	"fmt"
)

type DatabaseConfig struct {
	DriverName       string
	ConnectionString string
	MigrationsPath   string
}

var Env = func() *DatabaseConfig {
	var (
		driver         = "postgres"
		usr            = "admin"
		pwd            = "admin"
		host           = "db"
		port           = "5432"
		name           = "postgresdb"
		migrationsPath = "file://database/migrations"
	)

	return &DatabaseConfig{
		ConnectionString: fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", driver, usr, pwd, host, port, name),
		DriverName:       driver,
		MigrationsPath:   migrationsPath,
	}
}()
