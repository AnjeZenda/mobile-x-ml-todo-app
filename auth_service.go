package auth_service

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
}

type Details struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// .../register
func register(w http.ResponseWriter, r *http.Request) {
	var details Details

	if err := json.NewDecoder(r.Body).Decode(&details); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Регистрация нового пользователя: %s", details.Username)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message":           "OK, user registration!",
		"username_received": details.Username,
	})
}

// .../login
func login(w http.ResponseWriter, r *http.Request) {
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
