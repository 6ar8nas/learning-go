package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/6ar8nas/learning-go/config"
	"github.com/6ar8nas/learning-go/database"
	"github.com/6ar8nas/learning-go/middleware"
	"github.com/6ar8nas/learning-go/services/tasks"
	"github.com/6ar8nas/learning-go/services/users"
)

type ApiServer struct {
	*http.Server
	db *database.ConnectionPool
}

func InitApiServer(db *database.ConnectionPool) *ApiServer {
	router := http.NewServeMux()

	usersRepo := users.NewRepository(db)
	usersHandler := users.NewHandler(usersRepo)
	usersHandler.RegisterRoutes(router)

	tasksRepo := tasks.NewRepository(db)
	tasksHandler := tasks.NewHandler(tasksRepo)
	tasksHandler.RegisterRoutes(router)

	return &ApiServer{
		Server: &http.Server{
			Addr:         fmt.Sprintf(":%s", config.Port),
			Handler:      middleware.Logging(middleware.Authenticate(router)),
			IdleTimeout:  time.Minute,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
		db: db,
	}
}

func (s *ApiServer) Start() error {
	log.Println("Starting server on port", s.Addr)
	return s.ListenAndServe()
}

func (s *ApiServer) GracefulShutdown(done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for SIGINT
	<-ctx.Done()

	log.Println("Shutting down gracefully")

	// Allowing up to 5 seconds to finish server's requests
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Alert the main thread about a shutdown
	done <- true
}
