package main

import (
	healthleft "faker/internal/health/left"
	"faker/internal/writer/core"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"log"
)

func main() {
	log.Println("Starting Application")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/api/v1/health", healthleft.HealthHandler)

	err := router.Run(":8000")
	if err != nil {
		log.Fatal("Unable to start web server", err)
	}
	log.Println("Application Running")

	gocron.Every(1).Second().Do(core.GenerateAndSavePeople(1), 1, "hello")
	<-gocron.Start()
}
