package models

import (
	"time"

	"example.com/rest-api/db"
	"github.com/bytedance/gopkg/util/logger"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date" binding:"required"`
	UserID      int       `json:"user_id"`
}

var events = []Event{}

func (e *Event) Save() (int64, error) {
	query := `
	INSERT INTO events (name, description, location, dateTime, user_id) 
	VALUES (?,?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.DateTime)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Error(err)
		return 0, err
	}

	e.ID = id

	return id, nil
}

func GetAllEvents() []Event {
	return events
}
