package main

import (
	"auth/internal/adapter/db"
	"auth/pkg"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден, используются переменные окружения")
	}

	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer database.Close()

	app := pkg.New(database)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Сервер запускается на порту %s...", port)
	if err := app.Run(port); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
