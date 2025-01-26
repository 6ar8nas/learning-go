package tasks

import (
	"net/http"

	"github.com/6ar8nas/learning-go/server/types"
	sharedTypes "github.com/6ar8nas/learning-go/shared/types"
	sharedUtils "github.com/6ar8nas/learning-go/shared/utils"
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
	userId := sharedUtils.GetContextValue(ctx, types.ContextKeyUserId).(uuid.UUID)
	isAdmin := sharedUtils.GetContextValue(ctx, types.ContextKeyIsAdmin).(bool)
	tasks, err := h.repository.GetTasks(userId, isAdmin)
	if err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := sharedUtils.WriteJSON(w, http.StatusOK, &tasks); err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusBadRequest, "Expected id to be a valid UUID")
		return
	}

	ctx := r.Context()
	userId := sharedUtils.GetContextValue(ctx, types.ContextKeyUserId).(uuid.UUID)
	isAdmin := sharedUtils.GetContextValue(ctx, types.ContextKeyIsAdmin).(bool)
	task, err := h.repository.GetTaskById(id, userId, isAdmin)
	if err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	if err := sharedUtils.WriteJSON(w, http.StatusOK, &task); err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req sharedTypes.TaskCreateRequest
	if err := sharedUtils.ParseJSON(r, &req); err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	userId := sharedUtils.GetContextValue(r.Context(), types.ContextKeyUserId).(uuid.UUID)
	task, err := h.repository.CreateTask(userId, req)
	if err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := sharedUtils.WriteJSON(w, http.StatusCreated, &task); err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusBadRequest, "Expected id to be a valid UUID")
		return
	}

	var req sharedTypes.TaskUpdateRequest
	if err := sharedUtils.ParseJSON(r, &req); err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	userId := sharedUtils.GetContextValue(ctx, types.ContextKeyUserId).(uuid.UUID)
	isAdmin := sharedUtils.GetContextValue(ctx, types.ContextKeyIsAdmin).(bool)
	task, err := h.repository.UpdateTask(id, userId, isAdmin, req)
	if err != nil {
		switch err {
		case sharedTypes.ErrorNotFound:
			sharedUtils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
		default:
			sharedUtils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := sharedUtils.WriteJSON(w, http.StatusOK, &task); err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}
