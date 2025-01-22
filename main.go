package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		value, err := json.Marshal("Pong!")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(value)
	})

	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", router)
}
