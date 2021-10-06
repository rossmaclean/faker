package writercoremodel

type Person struct {
	PersonId         string
	FirstName        string
	LastName         string
	DateOfBirth      string
	Address          Address
	PhoneNumber      string
	EmailAddress     string
	CreditCardNumber string
}

type Address struct {
	AddressId      string
	BuildingNumber string
	StreetAddress  string
	City           string
	Country        string
	PostCode       string
}
