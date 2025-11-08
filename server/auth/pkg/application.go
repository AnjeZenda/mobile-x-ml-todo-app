package pkg

import (
	"auth/internal/handlers"
	"log"
	"os"
)

type App struct {
	Handlers *handlers.Handler
	Logger   *log.Logger
}

func New() *App {

	logger := log.New(os.Stdout, "Auth-service: ", log.LstdFlags)

	handler := handlers.New(logger)

	return &App{
		Handlers: handler,
		Logger:   logger,
	}
}
