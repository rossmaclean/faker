package writer

import (
	writercoremodel "faker/internal/writer/core/model"
)

type GeneratorService interface {
	GetFakePerson() writercoremodel.Person
}

var generatorService GeneratorService

func GetGeneratorService() GeneratorService {
	if generatorService != nil {
		return generatorService
	}
	generatorService = &DefaultGeneratorService{}
	return generatorService
}
