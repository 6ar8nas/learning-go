package main

import (
	"6ar8nas/test-app/api"
	"log"
)

func main() {
	server := api.InitApiServer()
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
	defer server.Close()
}
