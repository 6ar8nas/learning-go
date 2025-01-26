package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/6ar8nas/learning-go/api"
	"github.com/6ar8nas/learning-go/database"
	"github.com/6ar8nas/learning-go/server/config"
	"github.com/6ar8nas/learning-go/server/middleware"
	"github.com/6ar8nas/learning-go/server/services/tasks"
	"github.com/6ar8nas/learning-go/server/services/users"
)

type ApiServer struct {
	*api.ApiServer
	Database *database.ConnectionPool
}

func InitApiServer(port int, db *database.ConnectionPool) *ApiServer {
	router := http.NewServeMux()

	usersRepo := users.NewRepository(db)
	usersHandler := users.NewHandler(usersRepo)
	usersHandler.RegisterRoutes(router)

	tasksRepo := tasks.NewRepository(db)
	tasksHandler := tasks.NewHandler(tasksRepo)
	tasksHandler.RegisterRoutes(router)

	return &ApiServer{
		ApiServer: api.NewServer(
			&http.Server{
				Addr:         fmt.Sprintf(":%d", config.Port),
				Handler:      middleware.Logging(middleware.Authenticate(router)),
				IdleTimeout:  time.Minute,
				ReadTimeout:  10 * time.Second,
				WriteTimeout: 30 * time.Second,
			},
		),
		Database: db,
	}
}
