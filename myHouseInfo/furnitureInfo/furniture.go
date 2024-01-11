package furnitureInfo

import "fmt"

type Tables struct {
	Name   string
	Color  string
	Length float64
	Width  float64
	Height float64
}

type Chairs struct {
	Name   string
	Color  string
	Length float64
	Width  float64
	Height float64
}

type Beds struct {
	Name   string
	Color  string
	Size   string
	Length float64
	Width  float64
}

type StorageFurniture struct {
	Name   string
	Color  string
	Length float64
	Width  float64
	Height float64
}

type Lamps struct {
	Name   string
	Color  string
	Length float64
	Width  float64
	Power  float64
}

func TablesList() []Tables {
	table1 := Tables{"стол обеденный", "синий", 2, 1, 0.85}
	table2 := Tables{"стол рабочий", "белый", 1.8, 1.2, 0.9}

	return []Tables{table1, table2}
}

func ChairsList() []Chairs {
	chair1 := Chairs{"компьютерное кресло", "серое", 0.5, 0.6, 1.2}
	chair2 := Chairs{"стул", "синий", 0.5, 0.5, 0.9}
	chair3 := Chairs{"стул", "красный", 0.5, 0.5, 0.9}
	chair4 := Chairs{"стул", "желтый", 0.5, 0.5, 0.9}
	chair5 := Chairs{"стул", "зеленый", 0.5, 0.5, 0.9}

	return []Chairs{chair1, chair2, chair3, chair4, chair5}
}

func BedsList() []Beds {
	bed := Beds{"кровать", "белая", "двуспальная", 2.2, 1.7}

	return []Beds{bed}
}

func StorageFurnitureList() []StorageFurniture {
	closet := StorageFurniture{"шкаф-купе", "белый", 2.3, 1.2, 2}
	commode := StorageFurniture{"комод", "белый", 1.5, 0.7, 1.5}
	cupboard := StorageFurniture{"сервант", "черный", 1, 0.8, 1.8}
	bedsideTable := StorageFurniture{"тумбочка", "фиолетовая", 0.4, 0.4, 0.4}

	return []StorageFurniture{closet, commode, cupboard, bedsideTable}
}

func LampsList() []Lamps {
	chandelier1 := Lamps{"люстра", "белая", 0.4, 0.4, 100}
	chandelier2 := Lamps{"люстра", "прозрачная", 0.6, 0.6, 100}
	tableLamp := Lamps{"лампа", "серая", 0.15, 0.1, 40}

	return []Lamps{chandelier1, chandelier2, tableLamp}
}

func ShowTables(tables []Tables) {
	fmt.Println("Столы, стулья, хранение:")
	for _, oneOfTables := range tables {
		fmt.Print(oneOfTables.Name, " ", oneOfTables.Color, " ", oneOfTables.Length, "м ", oneOfTables.Width, "м ", oneOfTables.Height, "м\n")
	}
	fmt.Print("\n")
}

func ShowChairs(chairs []Chairs) {
	fmt.Println("Стулья:")
	for _, oneOfChairs := range chairs {
		fmt.Print(oneOfChairs.Name, " ", oneOfChairs.Color, " ", oneOfChairs.Length, "м ", oneOfChairs.Width, "м ", oneOfChairs.Height, "м\n")
	}
	fmt.Print("\n")
}

func ShowBeds(beds []Beds) {
	fmt.Println("Кровати:")
	for _, oneOfBeds := range beds {
		fmt.Print(oneOfBeds.Name, " ", oneOfBeds.Color, " ", oneOfBeds.Size, " ", oneOfBeds.Length, "м ", oneOfBeds.Width, "м\n")
	}
	fmt.Print("\n")
}

func ShowStorageFurniture(storageFurniture []StorageFurniture) {
	fmt.Println("Мебель для хранения:")
	for _, oneOfStorageFurniture := range storageFurniture {
		fmt.Print(oneOfStorageFurniture.Name, " ", oneOfStorageFurniture.Color, " ", oneOfStorageFurniture.Length, "м ", oneOfStorageFurniture.Width, "м ", oneOfStorageFurniture.Height, "м\n")
	}
	fmt.Print("\n")
}

func ShowLamps(lamps []Lamps) {
	fmt.Println("Лампы:")
	for _, oneOfLamps := range lamps {
		fmt.Print(oneOfLamps.Name, " ", oneOfLamps.Color, " ", oneOfLamps.Length, "м ", oneOfLamps.Width, "м ", oneOfLamps.Power, "Вт\n")
	}
	fmt.Print("\n")
}
