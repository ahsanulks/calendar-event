package repository_test

import (
	"context"
	"event-calendar/internal/app/database"
	"event-calendar/internal/app/domain/schedule/entity"
	schedule "event-calendar/internal/app/domain/schedule/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestScheduleRepository_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	dbGorm, _ := gorm.Open(dialector, &gorm.Config{})
	dbMock := database.Database{
		Conn: dbGorm,
	}
	type args struct {
		schedule *entity.Schedule
	}
	tests := []struct {
		name          string
		sr            *schedule.ScheduleRepository
		args          args
		wantErr       bool
		expectedQuery func(mock sqlmock.Sqlmock)
	}{
		{
			name: "error when schedule is null",
			args: args{
				nil,
			},
			sr:      schedule.NewScheduleRepository(&dbMock),
			wantErr: true,
			expectedQuery: func(m sqlmock.Sqlmock) {
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectedQuery(mock)
			_, err := tt.sr.Create(context.Background(), tt.args.schedule)
			if (err != nil) != tt.wantErr {
				t.Errorf("SchedulerUsecase.CreateSchedule() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
