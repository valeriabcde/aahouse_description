package appliancesInfo

import "fmt"

type TVs struct {
	Name            string
	Color           string
	Screen_diagonal float64
}

type Laptops struct {
	Name            string
	Color           string
	Screen_diagonal float64
}

type Stoves struct {
	Name   string
	Color  string
	Length float64
	Width  float64
	Height float64
}

type Conditioners struct {
	Name   string
	Color  string
	Length float64
	Width  float64
	Power  float64
}

type Fridges struct {
	Name   string
	Color  string
	Length float64
	Width  float64
	Height float64
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
	fmt.Println("Телевизоры и ноутбуки:")
	for _, oneOfTVs := range tvs {
		fmt.Print(oneOfTVs.Name, " ", oneOfTVs.Color, " ", oneOfTVs.Screen_diagonal, "''\n")
	}
	fmt.Print("\n")
}

func ShowLaptops(laptops []Laptops) {
	fmt.Println("Ноутбуки:")
	for _, oneOfLaptops := range laptops {
		fmt.Print(oneOfLaptops.Name, " ", oneOfLaptops.Color, " ", oneOfLaptops.Screen_diagonal, "''\n")
	}
	fmt.Print("\n")
}

func ShowStoves(stoves []Stoves) {
	fmt.Println("Плиты и холодильники:")
	for _, oneOfStoves := range stoves {
		fmt.Print(oneOfStoves.Name, " ", oneOfStoves.Color, " ", oneOfStoves.Length, "м ", oneOfStoves.Width, "м ", oneOfStoves.Height, "м\n")
	}
	fmt.Print("\n")
}

func ShowConditioners(conditioners []Conditioners) {
	fmt.Println("Кондиционеры:")
	for _, oneOfConditioners := range conditioners {
		fmt.Print(oneOfConditioners.Name, " ", oneOfConditioners.Color, " ", oneOfConditioners.Length, "м ", oneOfConditioners.Width, "м ", oneOfConditioners.Power, "кВт\n")
	}
	fmt.Print("\n")
}

func ShowFridges(fridges []Fridges) {
	fmt.Println("Холодильники:")
	for _, oneOfFridges := range fridges {
		fmt.Print(oneOfFridges.Name, " ", oneOfFridges.Color, " ", oneOfFridges.Length, "м ", oneOfFridges.Width, "м ", oneOfFridges.Height, "м\n")
	}
	fmt.Print("\n")
}
