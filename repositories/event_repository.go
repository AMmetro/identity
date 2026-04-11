package repositories

import (
	"database/sql"

	"github.com/AMmetro/identity/db"
	"github.com/AMmetro/identity/models"
)

func CreateEvent(e *models.Event) error {
	query := `INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

// CreateEventTx creates an event using the provided transaction and sets e.ID
func CreateEventTx(tx *sql.Tx, e *models.Event) error {
	// Use the provided transaction to execute the insert
	query := `INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)`
	result, err := tx.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetEventByID(id int64) (*models.Event, error) {
	query := `SELECT id, name, description, location, date_time, user_id FROM events where id = ?`
	row := db.DB.QueryRow(query, id)
	var event models.Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func GetAllEvents() ([]models.Event, error) {
	query := `SELECT id, name, description, location, date_time, user_id FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []models.Event
	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func UpdateEvent(e *models.Event) error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, date_time = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func DeleteEvent(id int64) error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}
