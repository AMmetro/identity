package models

// User is a domain model. DB access moved to repositories/services.
type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
