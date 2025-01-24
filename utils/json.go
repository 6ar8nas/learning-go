package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, value any) error {
	if r.Body == nil {
		return fmt.Errorf("request body is missing")
	}
	return json.NewDecoder(r.Body).Decode(value)
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

func WriteErrorJSON(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}
