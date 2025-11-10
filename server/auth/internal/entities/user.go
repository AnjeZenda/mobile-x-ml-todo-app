package entities

import "time"

// модельки по диаграммам

type User struct {
	ID       int    `json:"id"`
	Username string `json:"Username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
type Task struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Timestamp   time.Time `json:"timestamp"`
	Is_finished bool      `json:"is_finished"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
}
