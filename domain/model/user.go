package model

// User is model for `users` table
type User struct {
	Model
	Name  string `json:"name"`
	Email string `json:"email"`
}
