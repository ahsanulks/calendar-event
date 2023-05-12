package scheduler_test

import (
	"event-calendar/internal/app/use_case/scheduler"
	"testing"
)

func TestSchedulerUsecase_Invite(t *testing.T) {
	tests := []struct {
		name string
		s    *scheduler.SchedulerUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSchedulerUsecase()
			s.Invite()
		})
	}
}
