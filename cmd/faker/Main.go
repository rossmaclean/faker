package main

import (
	"faker/internal"
	healthleft "faker/internal/health/left"
	writer "faker/internal/writer/core"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"log"
	"time"
)

var cronRunning = false
var skippedCount = 0
var totalSkipped = 0

func main() {
	log.Println("Starting Application")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/api/v1/health", healthleft.HealthHandler)

	c := make(chan error)
	go runAsLoop(c)
	err := <-c
	if err != nil {
		log.Println(err)
	}

	err = router.Run(":8000")
	if err != nil {
		log.Fatal("Unable to start web server", err)
	}
}

func runAsLoop(c chan error) {
	fakerService := writer.GetFakerService()
	for i := 0; i < 1000000; i++ {
		err := fakerService.GenerateAndSavePeople(internal.GetFakerProperties().NumToGenerate)
		if err != nil {
			c <- err
		}
		sleepTime := time.Millisecond.Milliseconds() * int64(internal.GetFakerProperties().SleepMillis)

		time.Sleep(time.Duration(sleepTime))
	}
}

func runAsCron(c chan error) {
	cr := cron.New()
	err := cr.AddFunc("*/1 * * * *", runWithConcurrentChecks)
	if err != nil {
		c <- err
		return
	}
	cr.Start()
}

func runWithConcurrentChecks() {
	fakerService := writer.GetFakerService()
	if cronRunning {
		log.Printf("Cron already running, skipping, already skipped: %d, total skipped: %d",
			skippedCount, totalSkipped)
		skippedCount = skippedCount + 1
		totalSkipped = totalSkipped + 1
	} else {
		skippedCount = 0
		cronRunning = true
		err := fakerService.GenerateAndSavePeople(500)
		if err != nil {
			log.Println(err)
		}

		cronRunning = false
	}
}
