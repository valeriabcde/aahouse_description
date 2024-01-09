package main

import (
	"house/myHouseInfo/appliancesInfo"
	"house/myHouseInfo/familyInfo"
	"house/myHouseInfo/furnitureInfo"
	"house/myHouseInfo/roomsInfo"
)

func ShowDb() {

	rms := ShowRoomsDb()
	bds := ShowBedsDb()
	cnds := ShowConditionersDb()
	lmps := ShowLampsDb()
	pers := ShowPersonsDb()
	pts := ShowPetsDb()
	tbls := ShowTablesChairsStorageDb()
	stvs := ShowStovesFridgesDb()
	tvs := ShowTVsLaptopsDb()

	roomsInfo.ShowRooms(rms)
	furnitureInfo.ShowBeds(bds)
	appliancesInfo.ShowConditioners(cnds)
	furnitureInfo.ShowLamps(lmps)
	familyInfo.ShowPersons(pers)
	familyInfo.ShowPets(pts)
	furnitureInfo.ShowTables(tbls)
	appliancesInfo.ShowStoves(stvs)
	appliancesInfo.ShowTVs(tvs)
}
