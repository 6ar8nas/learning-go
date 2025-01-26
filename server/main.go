package main

import (
	"fmt"
	"net/http"

	"github.com/6ar8nas/learning-go/database"
	"github.com/6ar8nas/learning-go/server/api"
	"github.com/6ar8nas/learning-go/server/config"
)

func main() {
	db, err := database.NewConnection(config.Driver, config.ConnectionString, config.Database)
	if err != nil {
		panic(fmt.Sprintf("database connection error: %v", err))
	}
	defer db.Close()

	server := api.InitApiServer(config.Port, db)

	done := make(chan bool)
	go server.GracefulShutdown(done)

	if err := server.Start(); err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %v", err))
	}

	<-done
}
