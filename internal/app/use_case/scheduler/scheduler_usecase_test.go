package scheduler_test

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

type fakeEmailSender struct {
	called bool
}

func (fakeEmailSender *fakeEmailSender) SendBatchTo(ctx context.Context, guestsEmail []string, schedule *entity.Schedule) error {
	fakeEmailSender.called = true
	if guestsEmail[0] == "test@mail.com" {
		return nil
	}
	return errors.New("error")
}

func TestSchedulerUsecase_CreateSchedule(t *testing.T) {
	type args struct {
		schedule *dto.ScheduleRequestDto
	}
	tests := []struct {
		name            string
		su              *scheduler.SchedulerUsecase
		args            args
		wantErr         bool
		wantSchedulerId int
		wantSendEmail   bool
	}{
		{
			name: "error when endTime < startTime",
			args: args{
				&dto.ScheduleRequestDto{
					Name:      "testing schedule",
					StartTime: time.Now(),
					EndTime:   time.Now().Add(-2 * time.Hour),
				},
			},
			wantErr: true,
		},
		{
			name: "error when name empty",
			args: args{
				&dto.ScheduleRequestDto{
					Name:      "",
					StartTime: time.Now(),
					EndTime:   time.Now().Add(1 * time.Hour),
				},
			},
			wantErr: true,
		},
		{
			name: "error when start time < time now",
			args: args{
				&dto.ScheduleRequestDto{
					Name:      "test schedule",
					StartTime: time.Now().Add(-1 * time.Hour),
					EndTime:   time.Now().Add(1 * time.Hour),
				},
			},
			wantErr: true,
		},
		{
			name: "error when hitting database",
			args: args{
				&dto.ScheduleRequestDto{
					Name:      "test schedule",
					StartTime: time.Now().Add(2 * time.Hour),
					EndTime:   time.Now().Add(1 * time.Hour),
				},
			},
			wantErr: true,
		},
		{
			name: "error when sending invitation",
			args: args{
				&dto.ScheduleRequestDto{
					Name:        "testing",
					StartTime:   time.Now().Add(1 * time.Hour),
					EndTime:     time.Now().Add(2 * time.Hour),
					GuestsEmail: []string{"assd@mail.com"},
				},
			},
			wantErr:         true,
			wantSendEmail:   true,
			wantSchedulerId: 0,
		},
		{
			name: "success without sending invitation",
			args: args{
				&dto.ScheduleRequestDto{
					Name:      "testing",
					StartTime: time.Now().Add(1 * time.Hour),
					EndTime:   time.Now().Add(2 * time.Hour),
				},
			},
			wantErr:         false,
			wantSchedulerId: 1,
		},
		{
			name: "success with sending invitation",
			args: args{
				&dto.ScheduleRequestDto{
					Name:        "testing",
					StartTime:   time.Now().Add(1 * time.Hour),
					EndTime:     time.Now().Add(2 * time.Hour),
					GuestsEmail: []string{"test@mail.com"},
				},
			},
			wantErr:         false,
			wantSchedulerId: 1,
			wantSendEmail:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emailSender := new(fakeEmailSender)
			su := scheduler.NewSchedulerUsecase(new(fakeWriter), emailSender)
			schedulerId, err := su.CreateSchedule(context.Background(), tt.args.schedule)
			if (err != nil) != tt.wantErr {
				t.Errorf("SchedulerUsecase.CreateSchedule() error = %v, wantErr %v", err, tt.wantErr)
			}
			if schedulerId != tt.wantSchedulerId {
				t.Errorf("SchedulerUsecase.CreateSchedule() schedulerId = %d, schedulerId %d", schedulerId, tt.wantSchedulerId)
			}

			if emailSender.called != tt.wantSendEmail {
				t.Error("SchedulerUsecase.CreateSchedule() email not called")
			}
		})
	}
}
