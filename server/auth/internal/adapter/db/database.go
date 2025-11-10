package db

import (
	"auth/internal/entities"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) { //--------------------------------Тут все подключение к бд

	dsn := os.Getenv("---") // <-------------------Здесь через переменную окружения будет передаваться строка подключения к БД
	if dsn == "" {
		log.Fatal("Переменная окружения не установлена")
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Успешное подключение к PostgreSQL!")

	return db, nil
}

// Создание нового пользователя в бд
func CreateUser(db *sql.DB, username, password string, email string) error {
	query := `INSERT INTO user (name, password, email) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, username, password, email)
	return err
}

// GetUserByUsername находит пользователя по имени
func GetUserByUsername(db *sql.DB, email string) (*entities.User, error) {
	user := &entities.User{}
	query := `SELECT id, name, password, email FROM user WHERE email = $1`

	err := db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
