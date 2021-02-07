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

	// w := []rune{`Hello World`}
	// fmt.Println(w)
	// fmt.Println(w[0])
	// fmt.Printf("%T", w[0])

	// for i, c := range w {
	// 	fmt.Printf("Indicie[%d] %s\r\n", i, c)
	// }

	// const placeOfInterest = `⌘`

	// fmt.Printf("plain string: ")
	// fmt.Printf("%s", placeOfInterest)
	// fmt.Printf("\n")

	// fmt.Printf("quoted string: ")
	// fmt.Printf("%+q", placeOfInterest)
	// fmt.Printf("\n")

	// fmt.Printf("hex bytes: ")
	// for i := 0; i < len(placeOfInterest); i++ {
	// 	fmt.Printf("%x ", placeOfInterest[i])
	// }
	// fmt.Printf("\n")

	// const nihongo = "Hello World"
	// for index, runeValue := range nihongo {
	// 	fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	// }

}
