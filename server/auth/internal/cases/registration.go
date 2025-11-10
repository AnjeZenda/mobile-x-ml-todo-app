package cases

import (
	"auth/internal/adapter/db"
	"auth/internal/entities"
	"encoding/json"
	"net/http"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var details entities.User

	if err := json.NewDecoder(r.Body).Decode(&details); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Logger.Printf("Регистрация нового пользователя: %s", details.Username)

	if details.Username == "" || details.Password == "" {
		http.Error(w, "Ошибка: имя пользователя и пароль обязательны", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message":           "OK, user /register работает !",
		"username_received": details.Username,
	})

	err := db.CreateUser(h.db, details.Username, details.Password, details.Email)
	if err != nil {
		if err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"` {
			http.Error(w, "Пользователь с таким именем уже существует", http.StatusConflict)
		} else {
			http.Error(w, "Не удалось создать пользователя", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Пользователь успешно создан", http.StatusCreated)
}
