package main

import "fmt"

// Immobile - is a struct that stores a property.
type Immobile struct {
	X     int
	Y     int
	Name  string
	Value int
}

func main() {

	house := Immobile{}
	fmt.Printf("The house is: %+v\r\n", house)

	apartment := Immobile{17, 56, "My Apartment", 760000}
	fmt.Printf("The apartment is: %+v\r\n", apartment)

	farm := Immobile{
		Y:     85,
		X:     22,
		Name:  "FARM",
		Value: 55000,
	}
	fmt.Printf("The farm is: %+v\r\n", farm)

	house.Name = "Home Sweet Home"
	house.Value = 300000
	house.Y = 31
	house.X = 18

	fmt.Printf("The house is: %+v\r\n", house)
}
