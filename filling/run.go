package main

import "fmt"

func PrintHouseInfo(h House) {
	fmt.Printf("House Address: %v\n", h.Address)
	fmt.Println("\nFurniture:")
	for _, fur := range h.Furniture {
		fmt.Printf("%s is %s\n", fur.Name, fur.Color)
	}
	fmt.Println("\nAppliances:")
	for _, app := range h.Appliances {
		fmt.Printf("%s in the %s\n", app.Name, app.Room)
	}
	fmt.Println("\nFamily members:")
	for _, man := range h.Family {
		fmt.Printf("%s %s is %d y.o.\n", man.Role, man.Name, man.Age)
	}
	fmt.Println("\nAddress:")
	for _, add := range h.Address {
		fmt.Printf("%v\n", add.Address)
	}
	fmt.Println("\nHouse sizes:")
	for _, hou := range h.Size {
		fmt.Printf("House square is %d square meters\n", hou.Width*hou.Length)
	}
}

func runProject() {
	house := NewHouse()
	furs := getFur()
	apps := getApp()
	men := getMan()
	size := getSize()
	adds := getAddress()

	house.Address = adds
	house.Furniture = furs
	house.Appliances = apps
	house.Family = men
	house.Size = size

	PrintHouseInfo(house)
}
