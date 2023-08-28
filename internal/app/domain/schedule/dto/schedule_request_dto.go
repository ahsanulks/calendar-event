package dto

import "time"

type ScheduleRequestDto struct {
	Name        string    `json:"name" binding:"required"`
	StartTime   time.Time `json:"startTime" binding:"required"`
	EndTime     time.Time `json:"endTime" binding:"required"`
	GuestsEmail []string  `json:"guests" binding:"email"`
}
