package repositories

import (
	"database/sql"

	"github.com/AMmetro/identity/db"
	"github.com/AMmetro/identity/models"
)

func CreateRegistration(eventId, userId int64) (int64, error) {
	query := `INSERT INTO registrations (event_id, user_id) VALUES (?, ?)`
	result, err := db.DB.Exec(query, eventId, userId)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// CreateRegistrationTx inserts a registration using the provided transaction
func CreateRegistrationTx(tx *sql.Tx, eventId, userId int64) (int64, error) {
	query := `INSERT INTO registrations (event_id, user_id) VALUES (?, ?)`
	result, err := tx.Exec(query, eventId, userId)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func UpdateRegistrationStatus(eventId, userId int64, status string) error {
	query := `UPDATE registrations SET status = ? WHERE event_id = ? AND user_id = ?`
	_, err := db.DB.Exec(query, status, eventId, userId)
	return err
}

func GetAllRegistrations() ([]models.Registration, error) {
	query := `SELECT id, status, event_id, user_id FROM registrations`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var registrations []models.Registration
	for rows.Next() {
		var reg models.Registration
		if err := rows.Scan(&reg.ID, &reg.Status, &reg.Event_id, &reg.User_id); err != nil {
			return nil, err
		}
		registrations = append(registrations, reg)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return registrations, nil
}
