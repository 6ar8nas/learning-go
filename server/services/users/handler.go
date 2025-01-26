package users

import (
	"net/http"

	"github.com/6ar8nas/learning-go/auth"
	"github.com/6ar8nas/learning-go/server/config"
	"github.com/6ar8nas/learning-go/server/types"
	sharedTypes "github.com/6ar8nas/learning-go/shared/types"
	sharedUtils "github.com/6ar8nas/learning-go/shared/utils"
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
		sharedUtils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := sharedUtils.WriteJSON(w, http.StatusOK, &users); err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var req sharedTypes.UserAuthRequest
	if err := sharedUtils.ParseJSON(r, &req); err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.repository.GetUserByUsername(req.Username)
	if err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	if correct := VerifyPassword(req.Password, user.Password); !correct {
		sharedUtils.WriteErrorJSON(w, http.StatusBadRequest, "Wrong login credentials.")
		return
	}

	claims := make(map[string]interface{})
	claims[types.ClaimsKeyUserId] = user.Id
	claims[types.ClaimsKeyIsAdmin] = user.Admin
	authToken, err := auth.GenerateToken(claims, config.AuthSecret)
	if err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := sharedUtils.WriteJSON(w, http.StatusOK, &sharedTypes.UserAuthResponse{AuthToken: authToken}); err != nil { // TODO: return proper token
		sharedUtils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req sharedTypes.UserAuthRequest
	if err := sharedUtils.ParseJSON(r, &req); err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.repository.CreateUser(types.UserHashedAuthRequest{Username: req.Username, Password: hashedPassword})
	if err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := sharedUtils.WriteJSON(w, http.StatusCreated, &user); err != nil {
		sharedUtils.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
}
