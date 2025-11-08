package main

import (
	"auth/pkg"
	"net/http"
)

func main() {
	app := pkg.New()

	app.Logger.Println("Запускаем сервер регистрации/входа...")

	http.HandleFunc("/registration", app.Handlers.RegisterHandler)
	http.HandleFunc("/login", app.Handlers.LoginHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		app.Logger.Fatal(err)
	}
}
