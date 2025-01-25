package users

import (
	"6ar8nas/test-app/auth"
	"6ar8nas/test-app/config"
	"6ar8nas/test-app/types"
	"6ar8nas/test-app/utils"
	"net/http"
)

type Handler struct {
	repository types.UserRepository
}

func NewHandler(repository types.UserRepository) *Handler {
	return &Handler{repository: repository}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /users", h.GetUsers)
	router.HandleFunc("GET /login", h.AuthenticateUser)
	router.HandleFunc("POST /register", h.CreateUser)
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.repository.GetUsers()
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, &users); err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var req types.UserAuthRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.repository.GetUserByUsername(req.Username)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	if correct := auth.VerifyPassword(req.Password, user.Password); !correct {
		utils.WriteErrorJSON(w, http.StatusBadRequest, "Wrong login credentials.")
		return
	}

	authToken, err := auth.GenerateToken(user.Id, user.Admin, config.AuthSecret)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, &types.UserAuthResponse{AuthToken: authToken}); err != nil { // TODO: return proper token
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req types.UserAuthRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.repository.CreateUser(types.UserHashedAuthRequest{Username: req.Username, Password: hashedPassword})
	if err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.WriteJSON(w, http.StatusCreated, &user); err != nil {
		utils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}
