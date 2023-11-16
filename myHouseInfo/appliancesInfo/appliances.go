package appliancesInfo

import "fmt"

type TVs struct {
	name            string
	color           string
	screen_diagonal float64
}

type Laptops struct {
	name            string
	color           string
	screen_diagonal float64
}

type Stoves struct {
	name   string
	color  string
	length float64
	width  float64
	height float64
}

type Conditioners struct {
	name   string
	color  string
	length float64
	width  float64
	power  float64
}

type Fridges struct {
	name   string
	color  string
	length float64
	width  float64
	height float64
}

func TVsList() []TVs {
	tv1 := TVs{"телевизор", "серый", 52}
	tv2 := TVs{"телевизор", "черный", 70}

	return []TVs{tv1, tv2}
}

func LaptopsList() []Laptops {
	laptop := Laptops{"ноутбук", "серый", 15.6}

	return []Laptops{laptop}
}

func StovesList() []Stoves {
	stove := Stoves{"плита", "черная", 0.6, 0.6, 0.85}

	return []Stoves{stove}
}

func ConditionersList() []Conditioners {
	conditioner1 := Conditioners{"кондиционер", "белый", 0.87, 0.29, 5}
	conditioner2 := Conditioners{"кондиционер", "белый", 1.06, 0.3, 7}

	return []Conditioners{conditioner1, conditioner2}
}

func FridgesList() []Fridges {
	fridge := Fridges{"холодильник", "серый", 0.7, 0.67, 1.9}

	return []Fridges{fridge}

}

func ShowTVs(tvs []TVs) {
	fmt.Println("Телевизоры:")
	for _, oneOfTVs := range tvs {
		fmt.Print(oneOfTVs.name, " ", oneOfTVs.color, " ", oneOfTVs.screen_diagonal, "''\n")
	}
	fmt.Print("\n")
}

func ShowLaptops(laptops []Laptops) {
	fmt.Println("Ноутбуки:")
	for _, oneOfLaptops := range laptops {
		fmt.Print(oneOfLaptops.name, " ", oneOfLaptops.color, " ", oneOfLaptops.screen_diagonal, "''\n")
	}
	fmt.Print("\n")
}

func ShowStoves(stoves []Stoves) {
	fmt.Println("Плиты:")
	for _, oneOfStoves := range stoves {
		fmt.Print(oneOfStoves.name, " ", oneOfStoves.color, " ", oneOfStoves.length, "м ", oneOfStoves.width, "м ", oneOfStoves.height, "м\n")
	}
	fmt.Print("\n")
}

func ShowConditioners(conditioners []Conditioners) {
	fmt.Println("Кондиционеры:")
	for _, oneOfConditioners := range conditioners {
		fmt.Print(oneOfConditioners.name, " ", oneOfConditioners.color, " ", oneOfConditioners.length, "м ", oneOfConditioners.width, "м ", oneOfConditioners.power, "кВт\n")
	}
	fmt.Print("\n")
}

func ShowFridges(fridges []Fridges) {
	fmt.Println("Холодильники:")
	for _, oneOfFridges := range fridges {
		fmt.Print(oneOfFridges.name, " ", oneOfFridges.color, " ", oneOfFridges.length, "м ", oneOfFridges.width, "м ", oneOfFridges.height, "м\n")
	}
	fmt.Print("\n")
}
