package familyInfo

import "fmt"

type Persons struct {
	name    string
	surname string
	age     float64
	height  float64
}

type Pets struct {
	name   string
	breed  string
	age    float64
	height float64
}

func PersonsList() []Persons {
	woman := Persons{"Виктория", "Быкова", 39, 172}
	man := Persons{"Владимир", "Зуев", 42, 183}
	return []Persons{woman, man}
}
func PetsList() []Pets {
	cat := Pets{"Рыжик", "мейн-кун", 3, 24}
	turtle := Pets{"Донателло", "красноухая", 5, 6}
	return []Pets{cat, turtle}
}

func ShowPersons(persons []Persons) {
	fmt.Println("Семья:")
	for _, person := range persons {
		fmt.Print(person.name, " ", person.surname, " ", person.age, " ", person.height, "см\n")
	}
	fmt.Print("\n")
}
func ShowPets(pets []Pets) {
	fmt.Println("Домашние животные:")
	for _, pet := range pets {
		fmt.Print(pet.name, " ", pet.breed, " ", pet.age, " ", pet.height, "см\n")
	}
	fmt.Print("\n")
}
