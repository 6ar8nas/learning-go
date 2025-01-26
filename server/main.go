package main

import (
	"fmt"
	"net/http"

	"github.com/6ar8nas/learning-go/server/api"
	"github.com/6ar8nas/learning-go/server/database"
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
