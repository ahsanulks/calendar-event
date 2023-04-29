package reminder

import "time"

type Reminder struct {
	Id          int
	Name        string
	Description string
	UserId      int
	IsSent      bool
	ExecutedAt  time.Time
}

type Event struct {
	Id          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
