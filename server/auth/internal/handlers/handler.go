package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	Logger *log.Logger
}

func New(logger *log.Logger) *Handler {
	return &Handler{Logger: logger}
}
func (h *Handler) SomeAuthHandler(w http.ResponseWriter, r *http.Request) {

	h.Logger.Println("SomeAuthHandler called")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{

		"message": "Auth handler works!",
	})
}
