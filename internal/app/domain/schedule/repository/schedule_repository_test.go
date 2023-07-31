package repository_test

import (
	"context"
	"event-calendar/internal/app/database"
	"event-calendar/internal/app/domain/schedule/entity"
	"event-calendar/internal/app/domain/schedule/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestScheduleRepository_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
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

	scheduleVar := &entity.Schedule{
		Name:      "Test name schedule",
		StartTime: time.Now(),
		EndTime:   time.Now().Add(1 * time.Hour),
	}

	tests := []struct {
		name          string
		sr            *repository.ScheduleRepository
		args          args
		wantErr       bool
		expectedQuery func(mock sqlmock.Sqlmock)
	}{
		{
			name: "error when schedule is null",
			args: args{
				nil,
			},
			sr:      repository.NewScheduleRepository(&dbMock),
			wantErr: true,
			expectedQuery: func(m sqlmock.Sqlmock) {
			},
		},
		// TODO: Fix this
		// {
		// 	name: "success create",
		// 	args: args{
		// 		scheduleVar,
		// 	},
		// 	sr:      repository.NewScheduleRepository(&dbMock),
		// 	wantErr: false,
		// 	expectedQuery: func(m sqlmock.Sqlmock) {
		// 		m.ExpectBegin()
		// 		m.ExpectExec("INSERT INTO schedules (name, start_time, end_time) values (?, ?, ?) RETURNING \"id\"").WithArgs(
		// 			scheduleVar.Name,
		// 			scheduleVar.StartTime,
		// 			scheduleVar.EndTime,
		// 		).WillReturnResult(sqlmock.NewResult(1, 1))
		// 		m.ExpectCommit()
		// 	},
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expectedQuery(mock)
			_, err := tt.sr.Create(context.Background(), tt.args.schedule)
			if (err != nil) != tt.wantErr {
				t.Errorf("SchedulerUsecase.CreateSchedule() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
			// if !tt.wantErr {
			// 	if tt.args.schedule.Id == 0 {
			// 		t.Error("SchedulerUsecase.CreateSchedule() error id is = 0, want not 0")
			// 	}
			// }
		})
	}
}
