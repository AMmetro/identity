package repositories

import (
	"github.com/AMmetro/identity/internal/db"
	"github.com/AMmetro/identity/models"
)

func CreateUser(u *models.User) error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, email, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, email)
	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
