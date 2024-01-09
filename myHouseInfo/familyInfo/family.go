package familyInfo

import "fmt"

type Persons struct {
	Name    string
	Surname string
	Age     float64
	Height  float64
}

type Pets struct {
	Name  string
	Breed string
	Age   float64
}

func PersonsList() []Persons {
	woman := Persons{"Виктория", "Быкова", 39, 172}
	man := Persons{"Владимир", "Зуев", 42, 183}

	return []Persons{woman, man}
}

func PetsList() []Pets {
	cat := Pets{"Рыжик", "мейн-кун", 3}
	turtle := Pets{"Донателло", "красноухая", 5}

	return []Pets{cat, turtle}
}

func ShowPersons(persons []Persons) {
	fmt.Println("Семья:")
	for _, person := range persons {
		fmt.Print(person.Name, " ", person.Surname, " ", person.Age, " ", person.Height, "см\n")
	}
	fmt.Print("\n")
}

func ShowPets(pets []Pets) {
	fmt.Println("Домашние животные:")
	for _, pet := range pets {
		fmt.Print(pet.Name, " ", pet.Breed, " ", pet.Age, "\n")
	}
	fmt.Print("\n")
}
