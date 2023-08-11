package repository_test

import (
	"context"
	"event-calendar/internal/app/database"
	"event-calendar/internal/app/domain/schedule/entity"
	"event-calendar/internal/app/domain/schedule/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestScheduleRepository_Create(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	defer db.Close()
	dialector := mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})
	dbGorm, _ := gorm.Open(dialector, &gorm.Config{})
	dbMock := database.Database{
		Conn: dbGorm,
	}
	type args struct {
		schedule *entity.Schedule
	}

	scheduleMock := &entity.Schedule{
		Name:      "Test name schedule",
		StartTime: time.Now(),
		EndTime:   time.Now().Add(1 * time.Hour),
	}

	tests := []struct {
		name          string
		sr            *repository.ScheduleRepository
		args          args
		wantErr       bool
		wantId        int
		expectedQuery func(mock sqlmock.Sqlmock)
	}{
		{
			name: "error when schedule is null",
			args: args{
				nil,
			},
			sr:      repository.NewScheduleRepository(&dbMock),
			wantErr: true,
			wantId:  0,
			expectedQuery: func(m sqlmock.Sqlmock) {
			},
		},
		{
			name:    "success create",
			args:    args{scheduleMock},
			sr:      repository.NewScheduleRepository(&dbMock),
			wantErr: false,
			wantId:  1,
			expectedQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO `schedules` (`name`,`start_time`,`end_time`) VALUES (?,?,?)").WithArgs(scheduleMock.Name, scheduleMock.StartTime, scheduleMock.EndTime).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.MatchExpectationsInOrder(true)
			tt.expectedQuery(mock)
			id, err := tt.sr.Create(context.Background(), tt.args.schedule)
			if (err != nil) != tt.wantErr {
				t.Errorf("SchedulerUsecase.CreateSchedule() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

			if tt.wantId != id {
				t.Error("SchedulerUsecase.CreateSchedule() error id is = 0, want not 0")
			}
		})
	}
}
