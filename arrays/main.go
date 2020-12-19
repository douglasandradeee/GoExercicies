package main

import "fmt"

func main() {

	var test [3]int
	test[0] = 3
	test[1] = 2
	test[2] = 1
	fmt.Println("What is the capacity of this array? ", len(test), "elements.")

	singer := [2]string{"Michael Jackson", "John Lennon"}
	fmt.Printf("What's in this array? \r\n%v\r\n", singer)

	capital := [...]string{"Lisboa", "Brasilia", "Luanda", "Maputo"}
	fmt.Println("What is the capacity of this array? ", len(capital), "elements.")
	for index, city := range capital {
		fmt.Printf("Capital[%d] is %s\r\n", index, city)
	}

}
