package api

import (
	"auth/internal/cases"
	"database/sql"
	"log"
	"net/http"
	"os"
)

func NewAPI(db *sql.DB) http.Handler {

	logger := log.New(os.Stdout, "AuthAPI: ", log.LstdFlags)

	h := cases.New(logger, db)

	mux := http.NewServeMux()
	mux.HandleFunc("/auth", h.AuthHandler) //< -- Для проверки
	mux.HandleFunc("/register", h.Register)
	mux.HandleFunc("/login", h.Login)

	return mux
}
