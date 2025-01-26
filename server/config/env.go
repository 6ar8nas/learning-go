package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

var (
	Port, _            = strconv.Atoi(os.Getenv("PORT"))
	Database           = os.Getenv("DB_DATABASE")
	Driver             = os.Getenv("DB_DRIVER")
	DBConnectionString = os.Getenv("DB_CONNECTION_STRING")
	AuthSecret         = []byte(os.Getenv("AUTH_SECRET"))
)
