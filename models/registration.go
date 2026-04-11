package models

import (
	"github.com/AMmetro/identity/db"
)

type Registration struct {
	ID       int
	Status   string
	Event_id int64
	User_id  int64
}

func (e Event) RegisterForEvent(userId int64) (int64, error) {
	query := `INSERT INTO registrations (event_id, user_id) VALUES (?, ?)`

	result, err := db.DB.Exec(query, e.ID, userId)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (e Event) UpdateRegistration(userId int64, status string) error {
	query := `UPDATE registrations SET status = ? WHERE event_id = ? AND user_id = ?`
	_, err := db.DB.Exec(query, status, e.ID, userId)
	return err
}

func GetAllRegistrations() ([]Registration, error) {
	// select columns explicitly in the same order as rows.Scan below
	query := `SELECT id, status, event_id, user_id FROM registrations`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var registrations []Registration
	for rows.Next() {
		var reg Registration
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
