package task

import (
	"6ar8nas/test-app/types"
	"6ar8nas/test-app/utils"
	"net/http"

	"github.com/google/uuid"
)

type Handler struct {
	repository types.TaskRepository
}

func NewHandler(repository types.TaskRepository) *Handler {
	return &Handler{repository: repository}
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

	task, err := h.repository.GetTaskById(id)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
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

	task, err := h.repository.CreateTask(req)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, &task); err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}
