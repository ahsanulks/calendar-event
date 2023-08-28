package scheduler

import (
	"context"
	"event-calendar/internal/app/domain/schedule/dto"
	"event-calendar/internal/app/domain/schedule/entity"
)

type ScheduleWriter interface {
	Create(ctx context.Context, schedule *entity.Schedule) (int, error)
}

type EmailSender interface {
	SendBatchTo(ctx context.Context, email []string, schedule *entity.Schedule) error
}

type SchedulerUsecase struct {
	writer      ScheduleWriter
	emailSender EmailSender
}

func NewSchedulerUsecase(writer ScheduleWriter, emailSender EmailSender) *SchedulerUsecase {
	return &SchedulerUsecase{writer: writer, emailSender: emailSender}
}

func (su *SchedulerUsecase) CreateSchedule(ctx context.Context, schedule *dto.ScheduleRequestDto) (schedulerId int, err error) {
	scheduleEntity, err := entity.NewSchedule(schedule.Name, schedule.StartTime, schedule.EndTime)
	if err != nil {
		return
	}

	if len(schedule.GuestsEmail) > 0 {
		err = su.emailSender.SendBatchTo(ctx, schedule.GuestsEmail, scheduleEntity)
		if err != nil {
			return
		}
	}

	schedulerId, err = su.writer.Create(ctx, scheduleEntity)
	return
}
