package services

import (
	"errors"

	"github.com/AMmetro/identity/internal/repositories"
	"github.com/AMmetro/identity/models"
)

var allowedStatuses = map[string]bool{
	"pending":   true,
	"progress":  true,
	"success":   true,
	"cancelled": true,
}

func statusValidation(status string) bool {
	return allowedStatuses[status]
}

func UpdateRegistrationStatus(eventId, userId int64, status string) error {
	var isValidStatus = statusValidation(status)
	if !isValidStatus {
		return errors.New("invalid status for update registration")
	}
	currentStatus, err := repositories.GetRegistrationByIds(eventId, userId)
	if err != nil {
		return err
	}
	if currentStatus == status {
		return errors.New("status is up to date")
	}
	return repositories.UpdateRegistrationStatus(eventId, userId, status)
}

func GetAllRegistrations() ([]models.Registration, error) {
	return repositories.GetAllRegistrations()
}
