package repository

import (
	"context"
	"errors"
	"event-calendar/internal/app/database"
	"event-calendar/internal/app/domain/schedule/entity"
)

type ScheduleRepository struct {
	db *database.Database
}

func NewScheduleRepository(db *database.Database) *ScheduleRepository {
	return &ScheduleRepository{
		db: db,
	}
}

func (r *ScheduleRepository) Create(ctx context.Context, schedule *entity.Schedule) (int, error) {
	if schedule == nil {
		return 0, errors.New("error")
	}
	err := r.db.Conn.WithContext(ctx).Create(schedule).Error
	return schedule.Id, err
}
