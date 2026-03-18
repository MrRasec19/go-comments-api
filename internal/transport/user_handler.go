package transport

import (
	"encoding/json"
	"go-comments-api/internal/model"
	"go-comments-api/internal/service"
	"net/http"
	"strconv"
	"strings"
)

type UserHandler struct {
	service *service.Service
}

func New(s *service.Service) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users, err := h.service.GetAllUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	case http.MethodPost:
		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(user.Password) == "" {
			http.Error(w, "password is required", http.StatusBadRequest)
			return
		}
		created, err := h.service.CreateUser(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		created.Password = ""
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(created)
	default:
		http.Error(w, "Metodo no disponible", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) HandleUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Id no encontrado", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		user, err := h.service.GetBooksByID(id)
		if err != nil {
			http.Error(w, "id no encontrado", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	case http.MethodPut:
		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "input invalido", http.StatusBadRequest)
			return
		}

		updated, err := h.service.Updateuser(id, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updated)
	case http.MethodDelete:
		if err := h.service.Deleteuser(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Metodo no disponible", http.StatusMethodNotAllowed)
	}
}
