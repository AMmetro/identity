package models

import (
	"time"

	"github.com/AMmetro/identity/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	/*
	* alternative without Prepare
	 */
	// 	result, err := db.DB.Exec(
	// 	query,
	// 	e.Name,
	// 	e.Description,
	// 	e.Location,
	// 	e.DateTime,
	// 	e.UserID,
	// )

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

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events where id = ?`
	row := db.DB.QueryRow(query, id) // rows equals obgect type *sql.Rows = cursor (stream)
	// not need close
	// defer row.Close() - it will be closed automatically
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query) // rows equals obgect type *sql.Rows = cursor (stream)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	// update with SQLX
	for rows.Next() { // return false when rows is empty
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (event Event) UpdateEvent() error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, date_time = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) DeleteEvent() error {
	query := `Delete from events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	return err
}

func (e Event) RegisterForEvent(userId int64) (int64, error) {
	query := `INSERT INTO registrations (event_id, user_id) VALUES (?, ?)`

	result, err := db.DB.Exec(query, e.ID, userId)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (e Event) CancelRegistration(userId int64) error {
	query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`
	_, err := db.DB.Exec(query, e.ID, userId)
	return err
}
