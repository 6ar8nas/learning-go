package main

import (
	"6ar8nas/test-app/handlers"
	"6ar8nas/test-app/middleware"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /task/{taskId}", handlers.GetTask)
	router.HandleFunc("POST /task", handlers.PostTask)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: middleware.Logging(router),
	}

	log.Println("Starting server on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
