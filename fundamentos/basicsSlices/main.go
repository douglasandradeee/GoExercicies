package main

import "fmt"

func main() {

	var nums []int
	fmt.Println(nums, len(nums), cap(nums))
	nums = make([]int, 5)
	fmt.Println(nums, len(nums), cap(nums))

	capital := []string{"Lisboa"}
	fmt.Println(capital, len(capital), cap(capital))
	capital = append(capital, "Brasilia")
	fmt.Println(capital, len(capital), cap(capital))

	cities := make([]string, 4)
	cities[0] = "New York"
	cities[1] = "London"
	cities[2] = "Tokyo"
	cities[3] = "Singapore"
	fmt.Println(cities, len(cities), cap(cities))
	for index, city := range cities {
		fmt.Printf("City[%d] = %s\r\n", index, city)
	}
}
