package models

import (
	"github.com/AMmetro/identity/db"
	"github.com/AMmetro/identity/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func FindByEmail(email string) (*User, error) {
	query := `SELECT * FROM users WHERE email = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := db.DB.QueryRow(query, email) // rows equals obgect type *sql.Rows = cursor (stream)

	var user User
	err = row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) ValidateCredentials() (bool, error) {
	res, err := FindByEmail(u.Email)
	if err != nil {
		return false, err
	}

	isValid := utils.CheckPasswordHash(u.Password, res.Password)
	return isValid, nil
}
