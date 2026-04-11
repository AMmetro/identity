package services

import (
	"github.com/AMmetro/identity/db"
	"github.com/AMmetro/identity/models"
	"github.com/AMmetro/identity/repositories"
)

func CreateEvent(e *models.Event) error {
	return repositories.CreateEvent(e)
}

func CreateEventAndRegister(e *models.Event, userId int64) (int64, error) {
	// create a transaction so event creation and creator registration are atomic
	tx, err := db.DB.Begin()
	if err != nil {
		return 0, err
	}
	// create event within tx
	if err := repositories.CreateEventTx(tx, e); err != nil {
		tx.Rollback()
		return 0, err
	}
	// create registration within same tx
	regID, err := repositories.CreateRegistrationTx(tx, e.ID, userId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}
	return regID, nil
}

func GetEventByID(id int64) (*models.Event, error) {
	return repositories.GetEventByID(id)
}

func GetAllEvents() ([]models.Event, error) {
	return repositories.GetAllEvents()
}

func UpdateEvent(e *models.Event) error {
	return repositories.UpdateEvent(e)
}

func DeleteEvent(id int64) error {
	return repositories.DeleteEvent(id)
}
