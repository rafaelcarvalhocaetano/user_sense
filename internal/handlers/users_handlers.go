package handlers

import (
	"encoding/json"
	"net/http"

	"go-typesense-app/internal/models"
	"go-typesense-app/internal/users"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserHandlerInterface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	SearchUsers(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
}

type UserHandlerImpl struct {
	users users.UserModuleInterface
}

func NewUserHandler(users users.UserModuleInterface) *UserHandlerImpl {
	return &UserHandlerImpl{users: users}
}

func (h *UserHandlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.users.CreateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

func (h *UserHandlerImpl) GetUsers(w http.ResponseWriter, _ *http.Request) {
	allUsers, err := h.users.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(allUsers)
}

func (h *UserHandlerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updateData models.User
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user, err := h.users.UpdateUser(id, &updateData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}

func (h *UserHandlerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.users.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandlerImpl) SearchUsers(w http.ResponseWriter, r *http.Request) {
	var searchRequest struct {
		Query string `json:"query"`
	}

	if err := json.NewDecoder(r.Body).Decode(&searchRequest); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if searchRequest.Query == "" {
		http.Error(w, "Query field is required", http.StatusBadRequest)
		return
	}

	allUsers, err := h.users.SearchUsers(searchRequest.Query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"query":   searchRequest.Query,
		"results": allUsers,
		"count":   len(allUsers),
	})
}

func (h *UserHandlerImpl) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id := uuid.MustParse(idStr)

	user, err := h.users.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
