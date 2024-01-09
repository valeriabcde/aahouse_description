package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"house/myHouseInfo/appliancesInfo"
	"house/myHouseInfo/familyInfo"
	"house/myHouseInfo/furnitureInfo"
	"house/myHouseInfo/roomsInfo"
	"log"
)

func ShowBedsDb() []furnitureInfo.Beds {
	db, err := sql.Open("postgres", "postgres://house:123@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, color, size, length, width FROM beds")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	bds := make([]*furnitureInfo.Beds, 0)
	for rows.Next() {
		bd := new(furnitureInfo.Beds)
		err := rows.Scan(&bd.Name, &bd.Color, &bd.Size, &bd.Length, &bd.Width)
		if err != nil {
			log.Fatal(err)
		}
		bds = append(bds, bd)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	resultBeds := make([]furnitureInfo.Beds, len(bds))
	for i, bd := range bds {
		resultBeds[i] = furnitureInfo.Beds{
			Name:   bd.Name,
			Color:  bd.Color,
			Size:   bd.Size,
			Length: bd.Length,
			Width:  bd.Width,
		}
	}
	return resultBeds
}

func ShowConditionersDb() []appliancesInfo.Conditioners {
	db, err := sql.Open("postgres", "postgres://house:123@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, color, length, width, power FROM conditioners")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cnds := make([]*appliancesInfo.Conditioners, 0)
	for rows.Next() {
		cnd := new(appliancesInfo.Conditioners)
		err := rows.Scan(&cnd.Name, &cnd.Color, &cnd.Length, &cnd.Width, &cnd.Power)
		if err != nil {
			log.Fatal(err)
		}
		cnds = append(cnds, cnd)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	resultConditioners := make([]appliancesInfo.Conditioners, len(cnds))
	for i, cnd := range cnds {
		resultConditioners[i] = appliancesInfo.Conditioners{
			Name:   cnd.Name,
			Color:  cnd.Color,
			Length: cnd.Length,
			Width:  cnd.Width,
			Power:  cnd.Power,
		}
	}
	return resultConditioners
}

func ShowLampsDb() []furnitureInfo.Lamps {
	db, err := sql.Open("postgres", "postgres://house:123@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, color, length, width, power FROM lamps")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	lmps := make([]*furnitureInfo.Lamps, 0)
	for rows.Next() {
		lmp := new(furnitureInfo.Lamps)
		err := rows.Scan(&lmp.Name, &lmp.Color, &lmp.Length, &lmp.Width, &lmp.Power)
		if err != nil {
			log.Fatal(err)
		}
		lmps = append(lmps, lmp)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	resultLamps := make([]furnitureInfo.Lamps, len(lmps))
	for i, lmp := range lmps {
		resultLamps[i] = furnitureInfo.Lamps{
			Name:   lmp.Name,
			Color:  lmp.Color,
			Length: lmp.Length,
			Width:  lmp.Width,
			Power:  lmp.Power,
		}
	}
	return resultLamps
}

func ShowPersonsDb() []familyInfo.Persons {
	db, err := sql.Open("postgres", "postgres://house:123@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, surname, age, height FROM persons")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	pers := make([]*familyInfo.Persons, 0)
	for rows.Next() {
		per := new(familyInfo.Persons)
		err := rows.Scan(&per.Name, &per.Surname, &per.Age, &per.Height)
		if err != nil {
			log.Fatal(err)
		}
		pers = append(pers, per)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	resultPersons := make([]familyInfo.Persons, len(pers))
	for i, per := range pers {
		resultPersons[i] = familyInfo.Persons{
			Name:    per.Name,
			Surname: per.Surname,
			Age:     per.Age,
			Height:  per.Height,
		}
	}
	return resultPersons
}

func ShowPetsDb() []familyInfo.Pets {
	db, err := sql.Open("postgres", "postgres://house:123@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, breed, age FROM pets")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	pts := make([]*familyInfo.Pets, 0)
	for rows.Next() {
		pt := new(familyInfo.Pets)
		err := rows.Scan(&pt.Name, &pt.Breed, &pt.Age)
		if err != nil {
			log.Fatal(err)
		}
		pts = append(pts, pt)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	resultPets := make([]familyInfo.Pets, len(pts))
	for i, pt := range pts {
		resultPets[i] = familyInfo.Pets{
			Name:  pt.Name,
			Breed: pt.Breed,
			Age:   pt.Age,
		}
	}
	return resultPets
}

func ShowRoomsDb() []roomsInfo.Rooms {
	db, err := sql.Open("postgres", "postgres://house:123@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, length, width, height FROM rooms")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	rms := make([]*roomsInfo.Rooms, 0)
	for rows.Next() {
		rm := new(roomsInfo.Rooms)
		var err = rows.Scan(&rm.Name, &rm.Length, &rm.Width, &rm.Height)
		if err != nil {
			log.Fatal(err)
		}
		rms = append(rms, rm)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	resultRooms := make([]roomsInfo.Rooms, len(rms))
	for i, rm := range rms {
		resultRooms[i] = roomsInfo.Rooms{
			Name:   rm.Name,
			Length: rm.Length,
			Width:  rm.Width,
			Height: rm.Height,
		}
	}
	return resultRooms
}

func ShowStovesFridgesDb() []appliancesInfo.Stoves {
	db, err := sql.Open("postgres", "postgres://house:123@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, color, length, width, height FROM stoves_fridges")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	stvs := make([]*appliancesInfo.Stoves, 0)
	for rows.Next() {
		stv := new(appliancesInfo.Stoves)
		err := rows.Scan(&stv.Name, &stv.Color, &stv.Length, &stv.Width, &stv.Height)
		if err != nil {
			log.Fatal(err)
		}
		stvs = append(stvs, stv)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	resultStovesFridges := make([]appliancesInfo.Stoves, len(stvs))
	for i, stv := range stvs {
		resultStovesFridges[i] = appliancesInfo.Stoves{
			Name:   stv.Name,
			Color:  stv.Color,
			Length: stv.Length,
			Width:  stv.Width,
			Height: stv.Height,
		}
	}
	return resultStovesFridges
}

func ShowTablesChairsStorageDb() []furnitureInfo.Tables {
	db, err := sql.Open("postgres", "postgres://house:123@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, color, length, width, height FROM tables_chairs_storage")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	tbls := make([]*furnitureInfo.Tables, 0)
	for rows.Next() {
		tbl := new(furnitureInfo.Tables)
		err := rows.Scan(&tbl.Name, &tbl.Color, &tbl.Length, &tbl.Width, &tbl.Height)
		if err != nil {
			log.Fatal(err)
		}
		tbls = append(tbls, tbl)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	resultTablesChairsStorage := make([]furnitureInfo.Tables, len(tbls))
	for i, tbl := range tbls {
		resultTablesChairsStorage[i] = furnitureInfo.Tables{
			Name:   tbl.Name,
			Color:  tbl.Color,
			Length: tbl.Length,
			Width:  tbl.Width,
			Height: tbl.Height,
		}
	}
	return resultTablesChairsStorage
}

func ShowTVsLaptopsDb() []appliancesInfo.TVs {
	db, err := sql.Open("postgres", "postgres://house:123@localhost:5436/test_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT name, color, screen_diagonal FROM tvs_laptops")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	tvs := make([]*appliancesInfo.TVs, 0)
	for rows.Next() {
		tv := new(appliancesInfo.TVs)
		err := rows.Scan(&tv.Name, &tv.Color, &tv.Screen_diagonal)
		if err != nil {
			log.Fatal(err)
		}
		tvs = append(tvs, tv)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	resultTVsLaptops := make([]appliancesInfo.TVs, len(tvs))
	for i, tv := range tvs {
		resultTVsLaptops[i] = appliancesInfo.TVs{
			Name:            tv.Name,
			Color:           tv.Color,
			Screen_diagonal: tv.Screen_diagonal,
		}
	}
	return resultTVsLaptops
}
