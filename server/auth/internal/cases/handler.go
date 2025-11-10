package cases

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	Logger *log.Logger
	db     *sql.DB
}

func New(logger *log.Logger, db *sql.DB) *Handler {
	return &Handler{
		Logger: logger,
		db:     db,
	}
}

func (h *Handler) AuthHandler(w http.ResponseWriter, r *http.Request) {

	h.Logger.Println("Handler сервиса авторизации работает")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Handler сервиса авторизации работает",
	})
}
