package entity

import "time"

type Event struct {
	Id          int
	Name        string
	InvitorId   int
	InviteesIds []int
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewEvent(params map[string]interface{}) (*Event, error) {
	// validate params

	return &Event{}, nil
	// create validation
	// create datanya
}
