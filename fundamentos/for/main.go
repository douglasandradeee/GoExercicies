package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("What is the value of i? ", i)
	}

	value := 0
	test := true
	for test == true {
		value++
		if value%5 == 0 {
			test = false
		}
		fmt.Println("The number is: ", value)
	}

	for {
		value--
		fmt.Println("The number is: ", value)
		if value == 0 {
			break
		}
	}

	// Quebra as linhas entre as letras(incluindo o espaço) e me mostra o indicie da cada letra
	text := "I love writing programs using GO"
	for index, letter := range text {

		fmt.Printf("Text[%d] = %q\r\n", index, letter)
	}
}
