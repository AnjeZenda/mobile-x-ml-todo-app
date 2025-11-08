package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	Logger *log.Logger
}

type Details struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func New(logger *log.Logger) *Handler {
	return &Handler{
		Logger: logger,
	}
}

// .../register
func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var details Details

	if err := json.NewDecoder(r.Body).Decode(&details); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Logger.Printf("Регистрация нового пользователя: %s", details.Username)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message":           "OK, user registration!",
		"username_received": details.Username,
	})
}

// .../login
func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var details Details

	if err := json.NewDecoder(r.Body).Decode(&details); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "OK, user login!",
		"token":   "jwt.token.abc123",
	})

}
