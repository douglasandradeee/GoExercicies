package main

import "fmt"

var (
	address, phone string
	amount, stock  int
	bought         bool
	values         float64
	words          rune
)

func main() {
	test := "test value"
	fmt.Printf("address: %s\r\n", address)
	fmt.Printf("amount: %d\r\n", amount)
	fmt.Printf("bought: %v\r\n", bought)
	fmt.Printf("words: %v\r\n", words)
	fmt.Printf("test: %v\r\n", test)

}
