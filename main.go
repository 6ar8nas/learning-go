package main

import (
	"6ar8nas/test-app/api"
	"6ar8nas/test-app/database"
	"6ar8nas/test-app/database/migrations"
	"log"
)

func main() {
	db, err := database.InitDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := migrations.NewDriverDB(db).Migrate(); err != nil {
		log.Fatal(err)
	}

	server := api.InitApiServer(db)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
	defer server.Close()
}
