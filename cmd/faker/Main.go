package main

import (
	healthleft "faker/internal/health/left"
	"faker/internal/writer/core"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"log"
)

var cronRunning = false
var skippedCount = 0
var totalSkipped = 0

func main() {
	log.Println("Starting Application")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/api/v1/health", healthleft.HealthHandler)

	c := cron.New()
	c.AddFunc("*/1 * * * *", func() {
		if cronRunning {
			log.Printf("Cron already running, skipping, already skipped: %d, total skipped: %d",
				skippedCount, totalSkipped)
			skippedCount = skippedCount + 1
			totalSkipped = totalSkipped + 1
		} else {
			skippedCount = 0
			cronRunning = true
			core.GenerateAndSavePeople(50)
			cronRunning = false
		}
	})

	c.Start()

	err := router.Run(":8000")
	if err != nil {
		log.Fatal("Unable to start web server", err)
	}

}
