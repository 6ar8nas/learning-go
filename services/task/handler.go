package task

import (
	"6ar8nas/test-app/types"
	"6ar8nas/test-app/utils"
	"net/http"

	"github.com/google/uuid"
)

type Handler struct {
	cache *utils.Cache[uuid.UUID, types.Task]
}

func NewHandler() *Handler {
	return &Handler{cache: utils.NewCache[uuid.UUID, types.Task]()}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /task/{id}", h.GetTaskById)
	router.HandleFunc("POST /task", h.PostTask)
}

func (h *Handler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "Expected id to be a valid UUID")
		return
	}

	task, exists := h.cache.Get(id)
	if !exists {
		utils.WriteErrorJSON(w, http.StatusNotFound, "Expected task doesn't exist")
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, &task); err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) PostTask(w http.ResponseWriter, r *http.Request) {
	var req types.TaskCreateRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	task := types.Task{Id: uuid.New(), Type: req.Type, Status: types.Scheduled, Result: ""}
	h.cache.Set(task.Id, task)

	if err := utils.WriteJSON(w, http.StatusOK, &task); err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}
