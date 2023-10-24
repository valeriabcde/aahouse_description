package furnitureInfo

import "fmt"

type Tables struct {
	name   string
	color  string
	length float64
	width  float64
}

type Chairs struct {
	name   string
	color  string
	length float64
	width  float64
}

type Beds struct {
	name   string
	color  string
	length float64
	width  float64
}

type StorageFurniture struct {
	name   string
	color  string
	length float64
	width  float64
}

type Lamps struct {
	name   string
	color  string
	length float64
	width  float64
}

func TablesList() []Tables {
	table1 := Tables{"стол обеденный", "синий", 2, 1}
	table2 := Tables{"стол рабочий", "белый", 1.8, 1.2}
	return []Tables{table1, table2}
}
func ChairsList() []Chairs {
	chair1 := Chairs{"компьютерное кресло", "серое", 0.5, 0.6}
	chair2 := Chairs{"стул", "синий", 0.5, 0.5}
	chair3 := Chairs{"стул", "красный", 0.5, 0.5}
	chair4 := Chairs{"стул", "желтый", 0.5, 0.5}
	chair5 := Chairs{"стул", "зеленый", 0.5, 0.5}
	return []Chairs{chair1, chair2, chair3, chair4, chair5}
}
func BedsList() []Beds {
	bed := Beds{"кровать", "белая", 2.2, 1.7}
	return []Beds{bed}
}
func StorageFurnitureList() []StorageFurniture {
	closet := StorageFurniture{"шкаф-купе", "белый", 3, 1.5}
	commode := StorageFurniture{"комод", "белый", 2.1, 0.7}
	cupboard := StorageFurniture{"сервант", "черный", 1.8, 0.8}
	bedsideTable := StorageFurniture{"тумбочка", "фиолетовая", 0.8, 0.8}
	return []StorageFurniture{closet, commode, cupboard, bedsideTable}
}
func LampsList() []Lamps {
	chandelier1 := Lamps{"люстра", "белая", 0.4, 0.4}
	chandelier2 := Lamps{"люстра", "бесцветная", 0.6, 0.6}
	tableLamp := Lamps{"лампа", "серая", 0.15, 0.1}
	return []Lamps{chandelier1, chandelier2, tableLamp}

}

func ShowTables(tables []Tables) {
	fmt.Println("Столы:")
	for _, oneOfTables := range tables {
		fmt.Print(oneOfTables.name, " ", oneOfTables.color, " ", oneOfTables.length, "м ", oneOfTables.width, "м\n")
	}
	fmt.Print("\n")
}
func ShowChairs(chairs []Chairs) {
	fmt.Println("Стулья:")
	for _, oneOfChairs := range chairs {
		fmt.Print(oneOfChairs.name, " ", oneOfChairs.color, " ", oneOfChairs.length, "м ", oneOfChairs.width, "м\n")
	}
	fmt.Print("\n")
}
func ShowBeds(beds []Beds) {
	fmt.Println("Кровати:")
	for _, oneOfBeds := range beds {
		fmt.Print(oneOfBeds.name, " ", oneOfBeds.color, " ", oneOfBeds.length, "м ", oneOfBeds.width, "м\n")
	}
	fmt.Print("\n")
}
func ShowStorageFurniture(storageFurniture []StorageFurniture) {
	fmt.Println("Мебель для хранения:")
	for _, oneOfStorageFurniture := range storageFurniture {
		fmt.Print(oneOfStorageFurniture.name, " ", oneOfStorageFurniture.color, " ", oneOfStorageFurniture.length, "м ", oneOfStorageFurniture.width, "м\n")
	}
	fmt.Print("\n")
}
func ShowLamps(lamps []Lamps) {
	fmt.Println("Лампы:")
	for _, oneOfLamps := range lamps {
		fmt.Print(oneOfLamps.name, " ", oneOfLamps.color, " ", oneOfLamps.length, "м ", oneOfLamps.width, "м\n")
	}
	fmt.Print("\n")
}
