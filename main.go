package main

import (
	"6ar8nas/test-app/api"
	"6ar8nas/test-app/database"
	"fmt"
	"net/http"
)

func main() {
	db, err := database.NewService()
	if err != nil {
		panic(fmt.Sprintf("database connection error: %v", err))
	}
	defer db.Close()

	server := api.InitApiServer(db)

	done := make(chan bool)
	go server.GracefulShutdown(done)

	if err := server.Start(); err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %v", err))
	}

	<-done
}
