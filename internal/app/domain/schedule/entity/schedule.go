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

func NewSchedule(params map[string]interface{}) (*Schedule, error) {
	// validate params
	return &Schedule{}, nil
	// create validation
	// create datanya
}
