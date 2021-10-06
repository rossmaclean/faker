package core

import (
	writercoremodel "faker/internal/writer/core/model"
	writerright "faker/internal/writer/right"
	"fmt"
	"github.com/gocarina/gocsv"
	"log"
	"os"
	"strings"
	"syreclabs.com/go/faker"
	"time"
)

func GenerateAndSavePeople(amount int) error {
	for i := 0; i < amount; i++ {
		log.Printf("Generating and saving person %d", i)
		person := getFakePerson()
		writerright.GetPeopleRepository().SavePerson(person)
	}
	log.Printf("Generated and saved %d people", amount)
	return nil
}

func SavePeopleFromFileToDatabase(amount int) {
	people := GetPeopleFromFile()
	for i := 0; i < amount; i++ {
		writerright.GetPeopleRepository().SavePerson(*people[i])
	}
}

func GenerateAndSavePeopleToFile(amount int) {
	people := make([]writercoremodel.Person, 0)
	fPeople, errP := os.Create("./person.csv")
	fAdd, errA := os.Create("./address.csv")
	if errP != nil {
		log.Fatal(errP)
	}
	if errA != nil {
		log.Fatal(errA)
	}
	fPeople.WriteString("person_id, first_name, last_name, date_of_birth, address_id, phone_number, email_address, credit_card_number\n")
	fAdd.WriteString("address_id, building_number, street_address, city, country, post_code\n")
	for i := 0; i < amount; i++ {
		person := getFakePerson()
		people = append(people, person)
		personString := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s\n",
			strings.ReplaceAll(strings.ReplaceAll(person.PersonId, ",", ""), "'", ""),
			strings.ReplaceAll(strings.ReplaceAll(person.FirstName, ",", ""), "'", ""),
			strings.ReplaceAll(strings.ReplaceAll(person.LastName, ",", ""), "'", ""),
			strings.ReplaceAll(strings.ReplaceAll(person.DateOfBirth, ",", ""), "'", ""),
			strings.ReplaceAll(strings.ReplaceAll(person.PhoneNumber, ",", ""), "'", ""),
			strings.ReplaceAll(strings.ReplaceAll(person.EmailAddress, ",", ""), "'", ""),
			strings.ReplaceAll(strings.ReplaceAll(person.CreditCardNumber, ",", ""), "'", ""),
		)
		fPeople.WriteString(personString)

		aString := fmt.Sprintf("%s,%s,%s,%s,%s,%s\n",
			strings.ReplaceAll(strings.ReplaceAll(person.Address.AddressId, ",", ""), "'", ""),
			strings.ReplaceAll(strings.ReplaceAll(person.Address.BuildingNumber, ",", ""), "'", ""),
			strings.ReplaceAll(strings.ReplaceAll(person.Address.StreetAddress, ",", ""), "'", ""),
			strings.ReplaceAll(strings.ReplaceAll(person.Address.City, ",", ""), "'", ""),
			strings.ReplaceAll(strings.ReplaceAll(person.Address.Country, ",", ""), "'", ""),
			strings.ReplaceAll(strings.ReplaceAll(person.Address.PostCode, ",", ""), "'", ""),
		)
		fAdd.WriteString(aString)
	}
	fPeople.Sync()
	fAdd.Sync()
	fPeople.Close()
	fAdd.Close()
	log.Printf("Finished generating %d people", amount)
}

func GetPeopleFromFile() []*writercoremodel.Person {
	startTime := time.Now()

	pc := make(chan []*writercoremodel.Person)
	ac := make(chan []*writercoremodel.Address)

	go getPeopleFromFile(pc)
	go getAddressesFromFile(ac)

	people, addresses := <-pc, <-ac

	runTime := time.Since(startTime)
	log.Printf("Got %d people and %d addresses in %s", len(people), len(addresses), runTime)
	return people
}

func getPeopleFromFile(c chan []*writercoremodel.Person) {
	inP, err := os.Open("./person.csv")
	if err != nil {
		panic(err)
	}
	var people []*writercoremodel.Person

	if err := gocsv.UnmarshalFile(inP, &people); err != nil {
		panic(err)
	}
	c <- people
}

func getAddressesFromFile(c chan []*writercoremodel.Address) {
	inA, err := os.Open("./address.csv")
	if err != nil {
		panic(err)
	}
	var addresses []*writercoremodel.Address

	if err := gocsv.UnmarshalFile(inA, &addresses); err != nil {
		panic(err)
	}
	c <- addresses
}

func getFakePerson() writercoremodel.Person {
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
