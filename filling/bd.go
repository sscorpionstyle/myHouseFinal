package main

//
import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func openDb() *sql.DB {
	db, err := sql.Open("postgres", "postgres://house_new:123@localhost:5436/test_db1?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getFur() []Furniture {
	rows, err := openDb().Query("SELECT fur_name, fur_color FROM furniture")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	furs := make([]*Furniture, 0)
	for rows.Next() {
		fur := new(Furniture)
		err := rows.Scan(&fur.Name, &fur.Color)
		if err != nil {
			log.Fatal(err)
		}
		furs = append(furs, fur)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	convertedFurniture := make([]Furniture, len(furs))
	for i, fur := range furs {
		convertedFurniture[i] = Furniture{
			Name:  fur.Name,
			Color: fur.Color,
		}
	}
	return convertedFurniture
}

func getApp() []Appliance {
	rows, err := openDb().Query("SELECT app_name, app_room FROM appliance")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	apps := make([]*Appliance, 0)
	for rows.Next() {
		app := new(Appliance)
		err := rows.Scan(&app.Name, &app.Room)
		if err != nil {
			log.Fatal(err)
		}
		apps = append(apps, app)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	convertedAppliances := make([]Appliance, len(apps))
	for i, app := range apps {
		convertedAppliances[i] = Appliance{
			Name: app.Name,
			Room: app.Room,
		}
	}
	return convertedAppliances
}

func getMan() []FamilyMember {
	rows, err := openDb().Query("SELECT man_name, man_age, man_role FROM family")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	men := make([]*FamilyMember, 0)
	for rows.Next() {
		man := new(FamilyMember)
		err := rows.Scan(&man.Name, &man.Age, &man.Role)
		if err != nil {
			log.Fatal(err)
		}
		men = append(men, man)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	convertedMen := make([]FamilyMember, len(men))
	for i, man := range men {
		convertedMen[i] = FamilyMember{
			Name: man.Name,
			Age:  man.Age,
			Role: man.Role,
		}
	}
	return convertedMen
}

func getSize() []HouseSize {
	rows, err := openDb().Query("SELECT width, length FROM size")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	sizes := make([]*HouseSize, 0)
	for rows.Next() {
		size := new(HouseSize)
		err := rows.Scan(&size.Width, &size.Length)
		if err != nil {
			log.Fatal(err)
		}
		sizes = append(sizes, size)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	convertedSize := make([]HouseSize, len(sizes))
	for i, size := range sizes {
		convertedSize[i] = HouseSize{
			Width:  size.Width,
			Length: size.Length,
		}
	}
	return convertedSize
}

func getAddress() []HouseSize {
	rows, err := openDb().Query("SELECT address FROM size")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	adds := make([]*HouseSize, 0)
	for rows.Next() {
		add := new(HouseSize)
		err := rows.Scan(&add.Address)
		if err != nil {
			log.Fatal(err)
		}
		adds = append(adds, add)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	convertedAdd := make([]HouseSize, len(adds))
	for i, add := range adds {
		convertedAdd[i] = HouseSize{
			Address: add.Address,
		}
	}
	return convertedAdd
}
