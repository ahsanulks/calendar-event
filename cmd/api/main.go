package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"

	"event-calendar/internal/app/route"
	"event-calendar/internal/app/database"
)

func main() {
	gotenv.Load()

	db := database.NewDBConnection()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	route.NewRoute(r)

	

	r.Run(":" + os.Getenv("APP_PORT"))
}
