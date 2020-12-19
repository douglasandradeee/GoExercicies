package main

import "fmt"

func main() {

	capital := []string{"Lisboa"}
	//fmt.Println(capital, len(capital), cap(capital))
	capital = append(capital, "Brasilia")
	//fmt.Println(capital, len(capital), cap(capital))

	cities := make([]string, 5)
	cities[0] = "New York"
	cities[1] = "London"
	cities[2] = "Madeira"
	cities[3] = "Tokyo"
	cities[4] = "Singapore"
	//fmt.Println(cities, len(cities), cap(cities))
	for index, city := range cities {
		fmt.Printf("City[%d] = %s\r\n", index, city)
	}

	// First item starts with the index 0.
	// Second item starts with the index 1.
	asianCapitals := cities[3:5]
	fmt.Println(asianCapitals, len(asianCapitals), cap(asianCapitals))

	temp1 := cities[:2]
	fmt.Println(temp1)
	temp2 := cities[len(cities)-2:]
	fmt.Println(temp2)

	// Removing item from slice.
	indexRemovedItem := 2
	temp := cities[:indexRemovedItem]
	temp = append(temp, cities[indexRemovedItem+1:]...)
	cities = make([]string, len(temp))
	copy(cities, temp)
	fmt.Println("New Slice is :", cities)

}
