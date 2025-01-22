package handlers

import (
	"6ar8nas/test-app/handlers/tasks"
	"net/http"
)

func RegisterTasks(router *http.ServeMux) {
	router.HandleFunc("GET /task/{taskId}", tasks.GetTask)
	router.HandleFunc("POST /task", tasks.PostTask)
}
