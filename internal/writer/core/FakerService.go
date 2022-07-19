package writer

import (
	writerright "faker/internal/writer/right"
)

type FakerService interface {
	GenerateAndSavePeople(amount int) error
}

var fakerService FakerService

func GetFakerService() FakerService {
	if fakerService != nil {
		return fakerService
	}
	fakerService = &DefaultFakerService{
		GeneratorService: GetGeneratorService(),
		PeopleRepository: writerright.GetPeopleRepository(),
	}
	return fakerService
}
