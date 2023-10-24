package showMyHouse

import (
	"fmt"
	"house/myHouseInfo/appliancesInfo"
	"house/myHouseInfo/familyInfo"
	"house/myHouseInfo/furnitureInfo"
	"house/myHouseInfo/roomsInfo"
)

func ShowMyHouse() {
	rooms := roomsInfo.RoomsList()

	persons := familyInfo.PersonsList()
	pets := familyInfo.PetsList()

	tables := furnitureInfo.TablesList()
	chairs := furnitureInfo.ChairsList()
	beds := furnitureInfo.BedsList()
	storageFurniture := furnitureInfo.StorageFurnitureList()
	lamps := furnitureInfo.LampsList()

	tvs := appliancesInfo.TVsList()
	laptops := appliancesInfo.LaptopsList()
	stoves := appliancesInfo.StovesList()
	coditioners := appliancesInfo.ConditionersList()
	fridges := appliancesInfo.FridgesList()

	fmt.Println("МОЙ ДОМ")
	roomsInfo.ShowRooms(rooms)

	fmt.Println("В нем живут:")
	familyInfo.ShowPersons(persons)
	familyInfo.ShowPets(pets)

	furnitureInfo.ShowTables(tables)
	furnitureInfo.ShowChairs(chairs)
	furnitureInfo.ShowBeds(beds)
	furnitureInfo.ShowStorageFurniture(storageFurniture)
	furnitureInfo.ShowLamps(lamps)

	appliancesInfo.ShowTVs(tvs)
	appliancesInfo.ShowLaptops(laptops)
	appliancesInfo.ShowStoves(stoves)
	appliancesInfo.ShowConditioners(coditioners)
	appliancesInfo.ShowFridges(fridges)
}
