package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	Username         = os.Getenv("DB_USERNAME")
	Password         = os.Getenv("DB_PASSWORD")
	Host             = os.Getenv("DB_HOST")
	Port             = os.Getenv("PORT")
	DatabasePort     = os.Getenv("DB_PORT")
	Database         = os.Getenv("DB_DATABASE")
	Schema           = os.Getenv("DB_SCHEMA")
	Driver           = os.Getenv("DB_DRIVER")
	MigrationsPath   = os.Getenv("DB_MIGRATIONS")
	ConnectionString = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", Driver, Username, Password, Host, DatabasePort, Database, Schema)
)
