package main

import (
	healthleft "faker/internal/health/left"
	"faker/internal/writer/core"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"log"
)

func main() {
	log.Println("Starting Application")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/api/v1/health", healthleft.HealthHandler)

	c := cron.New()
	c.AddFunc("*/1 * * * *", func() {
		core.GenerateAndSavePeople(500)
	})

	c.Start()

	err := router.Run(":8000")
	if err != nil {
		log.Fatal("Unable to start web server", err)
	}

}
