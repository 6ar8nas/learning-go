package api

import (
	"6ar8nas/test-app/middleware"
	"6ar8nas/test-app/services/task"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type ApiServer struct {
	*http.Server
	db *sql.DB
}

func InitApiServer(db *sql.DB) *ApiServer {
	router := http.NewServeMux()

	taskRepo := task.NewRepository(db)
	taskHandler := task.NewHandler(taskRepo)
	taskHandler.RegisterRoutes(router)

	return &ApiServer{
		Server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", Env.Host, Env.Port),
			Handler: middleware.Logging(router),
		},
		db: db,
	}
}

func (s *ApiServer) Start() error {
	log.Println("Starting server on port", s.Addr)
	return s.ListenAndServe()
}
