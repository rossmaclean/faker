package writer

import (
	writerright "faker/internal/writer/right"
	"log"
	"time"
)

type DefaultFakerService struct {
	GeneratorService GeneratorService
	PeopleRepository writerright.PeopleRepository
}

func (r *DefaultFakerService) GenerateAndSavePeople(amount int) error {
	startTime := time.Now()
	for i := 0; i < amount; i++ {
		person := r.GeneratorService.GetFakePerson()
		err := r.PeopleRepository.SavePerson(person)
		if err != nil {
			return err
		}
	}
	runTime := time.Since(startTime)
	log.Printf("Generated and saved %d people in %s", amount, runTime)
	return nil
}
