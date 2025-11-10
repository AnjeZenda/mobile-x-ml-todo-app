package cases

import (
	"auth/internal/adapter/db"
	"auth/internal/entities"
	"database/sql"
	"encoding/json"
	"net/http"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var details entities.User

	if err := json.NewDecoder(r.Body).Decode(&details); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ищем пользователя в БД по email - так как email уникален скорее всего
	user, err := db.GetUserByUsername(h.db, details.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			h.Logger.Println("Неверное имя пользователя или пароль")
			http.Error(w, "Неверное имя пользователя или пароль", http.StatusUnauthorized)
		} else {
			http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		}
		return
	}

	// Проверяем пароль ---
	if user.Password != details.Password { // < --  Потом надо будет хеширование добавить
		h.Logger.Println("Неверное имя пользователя или пароль")
		http.Error(w, "Неверное имя пользователя или пароль", http.StatusUnauthorized)
		return
	}

	h.Logger.Printf(" Вход пользователя: %s", details.Email)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":           "OK, user login !",
		"username_received": details.Username + "/" + details.Email,
	})
}
