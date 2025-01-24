package api

import (
	"6ar8nas/test-app/middleware"
	"6ar8nas/test-app/services/task"
	"fmt"
	"log"
	"net/http"
)

type ApiServer struct {
	*http.Server
}

func InitApiServer() *ApiServer {
	router := http.NewServeMux()

	taskHandler := task.NewHandler()
	taskHandler.RegisterRoutes(router)

	return &ApiServer{
		Server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", Env.Host, Env.Port),
			Handler: middleware.Logging(router),
		},
	}
}

func (s *ApiServer) Start() error {
	log.Println("Starting server on port", s.Addr)
	return s.ListenAndServe()
}
