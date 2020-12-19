package main

import "fmt"

// Immobile -
type Immobile struct {
	X     int
	Y     int
	Name  string
	Value int
}

func main() {

	house := new(Immobile)
	fmt.Printf("House is: %p - %+v\r\n", &house, house)

	farm := Immobile{17, 28, "Beautiful FARM", 280000}
	fmt.Printf("Farm is: %p - %+v\r\n", &farm, farm)
	changeImmobile(&farm)
	fmt.Printf("Farm is: %p - %+v\r\n", &farm, farm)
}

func changeImmobile(immobile *Immobile) {
	immobile.X = immobile.X + 10
	immobile.Y = immobile.Y - 5
}
