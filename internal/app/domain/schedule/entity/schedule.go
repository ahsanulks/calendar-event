package entity

import "time"

type Schedule struct {
	Id        int
	Name      string
	StartTime time.Time
	EndTime   *time.Time
}

func (s *Schedule) isStarted() bool {
	return false
}
