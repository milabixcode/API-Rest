package models

import (
	"time"

	"API.com/api-test/db"
)

type Event struct {
	ID          int64     `binding: "required"`
	Name        string    `binding: "required"`
	Description string    `binding: "required"`
	Location    string    `binding: "required"`
	DateTime    time.Time `binding: "required"`
	UserID      int
}

var events = []Event{}

// Método que pertence ao tipo Event
func (e Event) Save() error {
	query := `
	INSERT INTO events(name, descrption, location, dateTime, user_id)
	VALUES(?, ?, ?, ?, ?)`
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

func GetAllEvents() []Event {
	return events
}
