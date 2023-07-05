package schedule_test

import (
	"context"
	"errors"
	"event-calendar/internal/app/domain/schedule/dto"
	"event-calendar/internal/app/domain/schedule/entity"
	"event-calendar/internal/app/use_case/scheduler"
	"testing"
	"time"
)

type fakeWriter struct{}

func (writer *fakeWriter) Create(ctx context.Context, schedule *entity.Schedule) (int, error) {
	if schedule.Name == "testing" {
		return 1, nil
	}
	return 0, errors.New("error")
}

func TestScheduleRepository_Create(t *testing.T) {
	type args struct {
		schedule *entity.Schedule
	}
	tests := []struct {
		name            string
		sr              *schedule.ScheduleRepository
		args            args
		wantErr         bool
	}{
		{
			name: "error when schedule is null",
			args: args{
				nil,
			},
			su:      scheduler.NewSchedulerUsecase(new(fakeWriter)),
			wantErr: true,
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			schedulerId, err := tt.sr.Create(context.Background(), tt.args.schedule)
			if (err != nil) != tt.wantErr {
				t.Errorf("SchedulerUsecase.CreateSchedule() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
