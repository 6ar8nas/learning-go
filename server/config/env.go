package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	Port             = os.Getenv("PORT")
	Database         = os.Getenv("DB_DATABASE")
	Driver           = os.Getenv("DB_DRIVER")
	ConnectionString = os.Getenv("DB_CONNECTION_STRING")
	AuthSecret       = []byte(os.Getenv("AUTH_SECRET"))
)
