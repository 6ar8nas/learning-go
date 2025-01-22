package api

import (
	"6ar8nas/test-app/handlers"
	"6ar8nas/test-app/middleware"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type ApiServer struct {
	*http.Server
}

func InitApiServer(db *sql.DB) *ApiServer {
	router := http.NewServeMux()

	handlers.RegisterAuth(router)
	handlers.RegisterTasks(router)

	return &ApiServer{
		Server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", Env.Host, Env.Port),
			Handler: middleware.Logging(router),
		},
	}
}

func (s *ApiServer) Start() error {
	log.Printf("Starting server on port %s\n", s.Addr)
	return s.ListenAndServe()
}
