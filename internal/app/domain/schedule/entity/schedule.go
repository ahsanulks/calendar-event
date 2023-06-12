package entity

import (
	"errors"
	"fmt"
	"time"
)

type Schedule struct {
	Id        int
	Name      string
	StartTime time.Time
	EndTime   time.Time
}

func NewSchedule(name string, startTime, endTime time.Time) (schedule *Schedule, err error) {
	schedule = &Schedule{
		Name:      name,
		StartTime: startTime,
		EndTime:   endTime,
	}
	err = schedule.validateCreateSchedule()
	if err != nil {
		schedule = nil
	}
	return
}

func (schedule *Schedule) validateCreateSchedule() (err error) {
	if schedule.StartTime.After(schedule.EndTime) {
		err = errors.New("start time must before end time")
		return
	}

	if schedule.Name == "" {
		err = errors.New("schedule name is required")
		return
	}
	if !schedule.StartTime.After(time.Now()) {
		err = fmt.Errorf("start time must after %s", time.Now().Format("2006-01-02 15:4"))
		return
	}
	return
}
