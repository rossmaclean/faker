package writer

import (
	writercoremodel "faker/internal/writer/core/model"
	"syreclabs.com/go/faker"
)

type FakeGeneratorService struct{}

func (r *FakeGeneratorService) GetFakePerson() writercoremodel.Person {
	address := writercoremodel.Address{
		AddressId:      faker.RandomString(9),
		BuildingNumber: faker.Address().BuildingNumber(),
		StreetAddress:  faker.Address().StreetAddress(),
		City:           faker.Address().City(),
		Country:        faker.Address().Country(),
		PostCode:       faker.Address().Postcode(),
	}

	return writercoremodel.Person{
		PersonId:         faker.RandomString(9),
		FirstName:        faker.Name().FirstName(),
		LastName:         faker.Name().LastName(),
		DateOfBirth:      faker.Date().Birthday(18, 100).Format("2006-06-01"),
		PhoneNumber:      faker.PhoneNumber().PhoneNumber(),
		EmailAddress:     faker.Internet().Email(),
		CreditCardNumber: faker.Finance().CreditCard(faker.CC_MASTERCARD, faker.CC_VISA),
		Address:          address,
	}
}
