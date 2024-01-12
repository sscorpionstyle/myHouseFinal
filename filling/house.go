package main

import "fmt"

type House struct {
	Address    []HouseSize
	Size       []HouseSize
	Family     []FamilyMember
	Furniture  []Furniture
	Appliances []Appliance
}

type HouseSize struct {
	Width   int
	Length  int
	Address string
}

func NewHouse() House {
	fmt.Println("My house:")
	return House{}
}
