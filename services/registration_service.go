package services

import (
	"github.com/AMmetro/identity/models"
	"github.com/AMmetro/identity/repositories"
)

func RegisterForEvent(eventId, userId int64) (int64, error) {
	return repositories.CreateRegistration(eventId, userId)
}

func UpdateRegistrationStatus(eventId, userId int64, status string) error {
	return repositories.UpdateRegistrationStatus(eventId, userId, status)
}

func GetAllRegistrations() ([]models.Registration, error) {
	return repositories.GetAllRegistrations()
}
