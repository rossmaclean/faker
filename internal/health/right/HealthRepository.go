package healthright

import (
	healthcoremodel "faker/internal/health/core/model"
	healthrightmongo "faker/internal/health/right/mongo"
)

type HealthRepository interface {
	InitDb()
	SaveHealth(health healthcoremodel.HealthResponse) error
}

var healthRepository HealthRepository

func GetHealthRepository() HealthRepository {
	if healthRepository != nil {
		return healthRepository
	}
	healthRepository = &healthrightmongo.MongoHealthRepository{}
	healthRepository.InitDb()
	return healthRepository
}
