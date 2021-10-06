package healthleft

import (
	healthcore "faker/internal/health/core"
	"github.com/gin-gonic/gin"
	"log"
)

func HealthHandler(c *gin.Context) {
	health, err := healthcore.GetHealth()
	if err != nil {
		log.Printf("Health: %s", health)
		c.JSON(500, health)
		return
	}
	c.JSON(200, health)
}
