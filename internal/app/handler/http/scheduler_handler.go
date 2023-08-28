package http

import (
	"context"
	"event-calendar/internal/app/domain/schedule/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SchedulerHandler struct {
	scheduleCreator ScheduleCreator
}

type ScheduleCreator interface {
	CreateSchedule(ctx context.Context, schedule *dto.ScheduleRequestDto) (schedulerId int, err error)
}

func NewSchedulerHandler(scheduleCreator ScheduleCreator) *SchedulerHandler {
	return &SchedulerHandler{
		scheduleCreator: scheduleCreator,
	}
}

func (h *SchedulerHandler) CreateSchedule(c *gin.Context) {
	var requestBody dto.ScheduleRequestDto
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
		return
	}

	id, err := h.scheduleCreator.CreateSchedule(context.Background(), &requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
