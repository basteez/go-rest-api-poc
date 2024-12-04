package models

import (
	"database/sql"
	"fmt"
	"time"

	"bstz.it/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"date_time"`
	UserID      int       `json:"user_id"`
}

func (event *Event) Save() error {
	query := `
		INSERT INTO events(name, description, location, date_time, user_id) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	err = statement.QueryRow(event.Name, event.Description, event.Location, event.DateTime, event.UserID).Scan(&event.ID)
	if err != nil {
		return err
	}

	return err
}

func (event *Event) Update() error {
	query := `
	UPDATE events 
	SET name = $1, description = $2, location = $3, date_time = $4
	WHERE id = $5`

	args := []any{
		event.Name,
		event.Description,
		event.Location,
		event.DateTime,
		event.ID,
	}

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(args...)
	if err != nil {
		return err
	}

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = $1`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("event with id %d not found: %w", id, sql.ErrNoRows)
		}
		return nil, err
	}

	return &event, nil

}

func (event *Event) Delete() error {
	query := `DELETE FROM events WHERE id = $1`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(event.ID)

	return err

}
