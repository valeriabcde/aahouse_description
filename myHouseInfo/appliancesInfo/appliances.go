package appliancesInfo

import "fmt"

type TVs struct {
	name   string
	color  string
	length float64
	width  float64
}

type Laptops struct {
	name   string
	color  string
	length float64
	width  float64
}

type Stoves struct {
	name   string
	color  string
	length float64
	width  float64
}

type Conditioners struct {
	name   string
	color  string
	length float64
	width  float64
}

type Fridges struct {
	name   string
	color  string
	length float64
	width  float64
}

func TVsList() []TVs {
	tv1 := TVs{"стол обеденный", "синий", 2, 1}
	tv2 := TVs{"стол рабочий", "белый", 1.8, 1.2}
	return []TVs{tv1, tv2}
}
func LaptopsList() []Laptops {
	laptop := Laptops{"компьютерное кресло", "серое", 0.3, 0.2}
	return []Laptops{laptop}
}
func StovesList() []Stoves {
	stove := Stoves{"кровать", "белая", 2.2, 1.7}
	return []Stoves{stove}
}
func ConditionersList() []Conditioners {
	conditioner1 := Conditioners{"шкаф-купе", "белый", 3, 1.5}
	conditioner2 := Conditioners{"комод", "белый", 2.1, 0.7}
	return []Conditioners{conditioner1, conditioner2}
}
func FridgesList() []Fridges {
	fridge := Fridges{"холодильник", "белый", 1.8, 1.7}
	return []Fridges{fridge}

}

func ShowTVs(tvs []TVs) {
	fmt.Println("Телевизоры:")
	for _, oneOfTVs := range tvs {
		fmt.Print(oneOfTVs.name, " ", oneOfTVs.color, " ", oneOfTVs.length, "м ", oneOfTVs.width, "м\n")
	}
	fmt.Print("\n")
}
func ShowLaptops(laptops []Laptops) {
	fmt.Println("Ноутбуки:")
	for _, oneOfLaptops := range laptops {
		fmt.Print(oneOfLaptops.name, " ", oneOfLaptops.color, " ", oneOfLaptops.length, "м ", oneOfLaptops.width, "м\n")
	}
	fmt.Print("\n")
}
func ShowStoves(stoves []Stoves) {
	fmt.Println("Кровати:")
	for _, oneOfStoves := range stoves {
		fmt.Print(oneOfStoves.name, " ", oneOfStoves.color, " ", oneOfStoves.length, "м ", oneOfStoves.width, "м\n")
	}
	fmt.Print("\n")
}
func ShowConditioners(conditioners []Conditioners) {
	fmt.Println("Столы:")
	for _, oneOfConditioners := range conditioners {
		fmt.Print(oneOfConditioners.name, " ", oneOfConditioners.color, " ", oneOfConditioners.length, "м ", oneOfConditioners.width, "м\n")
	}
	fmt.Print("\n")
}
func ShowFridges(fridges []Fridges) {
	fmt.Println("Холодильники:")
	for _, oneOfFridges := range fridges {
		fmt.Print(oneOfFridges.name, " ", oneOfFridges.color, " ", oneOfFridges.length, "м ", oneOfFridges.width, "м\n")
	}
	fmt.Print("\n")
}
