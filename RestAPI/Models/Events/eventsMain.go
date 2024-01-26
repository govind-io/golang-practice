package event

import (
	"time"
)

var events = []Event{}

type Event struct {
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	Location    string    `json:"loc"`
	UserId      int       `json:"uid"`
	DateTime    time.Time `json:"date"`
}

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}

func NewEvent(name string, description string, location string, userId int) *Event {
	return &Event{
		Name:        name,
		Description: description,
		Location:    location,
		UserId:      userId,
		DateTime:    time.Now(),
	}
}
