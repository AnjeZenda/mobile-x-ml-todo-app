package pkg

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"auth/internal/adapter/api"
)

type App struct {
	DB     *sql.DB
	API    http.Handler
	Logger *log.Logger
}

func New(db *sql.DB) *App {
	logger := log.New(os.Stdout, "Auth-service: ", log.LstdFlags)

	apiHandler := api.NewAPI(db)

	return &App{
		DB:     db,
		API:    apiHandler,
		Logger: logger,
	}
}

func (a *App) Run(port string) error {
	if port == "" {
		port = "8080"
	}
	a.Logger.Printf("Запущен сервис статистики на порту %s", port)
	return http.ListenAndServe(":"+port, a.API)
}
