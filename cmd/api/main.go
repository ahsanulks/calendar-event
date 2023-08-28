package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"

	"event-calendar/internal/app/database"
	scheduleRepo "event-calendar/internal/app/domain/schedule/repository"

	handler "event-calendar/internal/app/handler/http"
	"event-calendar/internal/app/route"
	"event-calendar/internal/app/use_case/email"
	"event-calendar/internal/app/use_case/scheduler"
)

func main() {
	gotenv.Load()

	db := database.NewDBConnection()
	repo := scheduleRepo.NewScheduleRepository(db)
	emailGenerator := email.EmailGenerator{}

	schedulerUseCase := scheduler.NewSchedulerUsecase(repo, emailGenerator)
	scheduleHandler := handler.NewSchedulerHandler(schedulerUseCase)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.POST("/schedules", scheduleHandler.CreateSchedule)

	route.NewRoute(r)

	r.Run(":" + os.Getenv("APP_PORT"))

}
