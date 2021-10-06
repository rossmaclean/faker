package writerright

import (
	writercoremodel "faker/internal/writer/core/model"
	writerrightmongo "faker/internal/writer/right/mongo"
)

type PeopleRepository interface {
	InitDb()
	Ping() error
	SavePerson(person writercoremodel.Person) error
}

var peopleRepository PeopleRepository

func GetPeopleRepository() PeopleRepository {
	if peopleRepository != nil {
		return peopleRepository
	}
	peopleRepository = &writerrightmongo.MongoPeopleRepository{}
	peopleRepository.InitDb()
	return peopleRepository
}
