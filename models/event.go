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
	UserID      int
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

func CreateEvents() []Event {
	var events []Event = []Event{}
	return events
}
