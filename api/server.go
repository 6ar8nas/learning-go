package api

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type ApiServer struct {
	*http.Server
}

func NewServer(server *http.Server) *ApiServer {
	return &ApiServer{Server: server}
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
