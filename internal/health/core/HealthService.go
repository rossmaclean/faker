package healthcore

import (
	"errors"
	healthcoremodel "faker/internal/health/core/model"
	healthright "faker/internal/health/right"
	writerright "faker/internal/writer/right"
	"log"
)

func GetHealth() (healthcoremodel.HealthResponse, error) {
	var healthStatuses []healthcoremodel.HealthStatus
	healthStatuses = append(healthStatuses, getDatabaseHealth())

	isErr := false
	for _, status := range healthStatuses {
		if status.Status == "down" {
			isErr = true
		}
	}

	if isErr {
		return healthcoremodel.HealthResponse{HealthStatuses: healthStatuses}, errors.New("")
	}

	health := healthcoremodel.HealthResponse{HealthStatuses: healthStatuses}
	err := healthright.GetHealthRepository().SaveHealth(health)
	if err != nil {
		log.Println(err)
	}
	return health, nil
}

func getDatabaseHealth() healthcoremodel.HealthStatus {
	status := healthcoremodel.HealthStatus{
		System: "database",
		Status: "up",
	}
	err := writerright.GetPeopleRepository().Ping()
	if err != nil {
		status.Status = "down"
	}
	return status
}
