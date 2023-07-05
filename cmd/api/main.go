package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"

	"event-calendar/internal/app/database"
	scheduleRepo "event-calendar/internal/app/domain/schedule/repository"
	"event-calendar/internal/app/route"
)

func main() {
	gotenv.Load()

	db := database.NewDBConnection()
	scheduleRepo.NewScheduleRepository(db)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	route.NewRoute(r)

	r.Run(":" + os.Getenv("APP_PORT"))
}
