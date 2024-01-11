package roomsInfo

import "fmt"

type Rooms struct {
	Name   string
	Length float64
	Width  float64
	Height float64
}

func RoomsList() []Rooms {
	bedroom := Rooms{"спальня", 12, 9.3, 2.5}
	kitchen := Rooms{"кухня", 8.2, 6.7, 2.5}

	return []Rooms{bedroom, kitchen}
}

func ShowRooms(rooms []Rooms) {
	fmt.Print("Размеры комнаты (длина, ширина, высота):\n")
	for _, room := range rooms {
		fmt.Print(room.Name, " ", room.Length, room.Width, room.Height, "\n")
	}
}
