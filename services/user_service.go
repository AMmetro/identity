package services

import (
	"errors"

	"github.com/AMmetro/identity/internal/repositories"
	"github.com/AMmetro/identity/models"
	"github.com/AMmetro/identity/utils"
)

func RegisterUser(u *models.User) error {
	hashed, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashed
	return repositories.CreateUser(u)
}

func AuthenticateUser(email, password string) (int64, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return 0, errors.New("credentials invalid")
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return 0, errors.New("credentials invalid")
	}
	return user.ID, nil
}
