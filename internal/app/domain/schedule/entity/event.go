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
