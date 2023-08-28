package http

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSchedulerHandler_CreateSchedule(t *testing.T) {
	type fields struct {
		scheduleCreator ScheduleCreator
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &SchedulerHandler{
				scheduleCreator: tt.fields.scheduleCreator,
			}
			h.CreateSchedule(tt.args.c)
		})
	}
}
