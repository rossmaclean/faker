package writer

import "fmt"

type FakeFakerService struct{}

func (r *FakeFakerService) GenerateAndSavePeople(amount int) error {
	fmt.Println(amount)
	return nil
}
