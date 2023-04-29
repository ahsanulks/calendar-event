package main

import (
	"github.com/gin-gonic/gin"

	"event-calendar/internal/app/route"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	route.NewRoute(r)

	r.Run(":3000")
}
