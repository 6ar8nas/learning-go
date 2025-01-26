package tasks

import (
	"net/http"

	"github.com/6ar8nas/learning-go/server/types"
	"github.com/6ar8nas/learning-go/server/utils"
	"github.com/google/uuid"
)

type Handler struct {
	repository types.TaskRepository
}

func NewHandler(repository types.TaskRepository) *Handler {
	return &Handler{repository: repository}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /tasks", h.GetTasks)
	router.HandleFunc("GET /tasks/{id}", h.GetTaskById)
	router.HandleFunc("POST /tasks", h.CreateTask)
	router.HandleFunc("PATCH /tasks/{id}", h.UpdateTask)
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := utils.GetContextValue(ctx, types.ContextKeyUserId).(uuid.UUID)
	isAdmin := utils.GetContextValue(ctx, types.ContextKeyIsAdmin).(bool)
	tasks, err := h.repository.GetTasks(userId, isAdmin)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, &tasks); err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "Expected id to be a valid UUID")
		return
	}

	ctx := r.Context()
	userId := utils.GetContextValue(ctx, types.ContextKeyUserId).(uuid.UUID)
	isAdmin := utils.GetContextValue(ctx, types.ContextKeyIsAdmin).(bool)
	task, err := h.repository.GetTaskById(id, userId, isAdmin)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, &task); err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req types.TaskCreateRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	userId := utils.GetContextValue(r.Context(), types.ContextKeyUserId).(uuid.UUID)
	task, err := h.repository.CreateTask(userId, req)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.WriteJSON(w, http.StatusCreated, &task); err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "Expected id to be a valid UUID")
		return
	}

	var req types.TaskUpdateRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	userId := utils.GetContextValue(ctx, types.ContextKeyUserId).(uuid.UUID)
	isAdmin := utils.GetContextValue(ctx, types.ContextKeyIsAdmin).(bool)
	task, err := h.repository.UpdateTask(id, userId, isAdmin, req)
	if err != nil {
		switch err {
		case types.ErrorNotFound:
			utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
		default:
			utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, &task); err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}
