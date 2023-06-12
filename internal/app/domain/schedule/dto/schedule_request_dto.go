package dto

import "time"

type ScheduleRequestDto struct {
	Name      string
	StartTime time.Time
	EndTime   time.Time
}
