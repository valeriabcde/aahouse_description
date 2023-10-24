package roomsInfo

import "fmt"

type Rooms struct {
	name   string
	length float64
	width  float64
	height float64
}

func RoomsList() []Rooms {
	bedroom := Rooms{"спальня", 12, 9.3, 2.5}
	kitchen := Rooms{"кухня", 8.2, 6.7, 2.5}
	return []Rooms{bedroom, kitchen}
}

func ShowRooms(rooms []Rooms) {
	fmt.Print("Размеры комнаты(длина, ширина, высота):\n")
	for _, room := range rooms {
		fmt.Print(room.name, " ", room.length, room.width, room.height, "\n")
	}
}
