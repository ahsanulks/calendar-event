package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"

	"event-calendar/internal/app/route"
)

func main() {
	gotenv.Load()
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	route.NewRoute(r)

	r.Run(":" + os.Getenv("APP_PORT"))
}
