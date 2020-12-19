package main

import (
	"fmt"

	"github.com/douglasandradeee/GoExercicies/advancedFunctions/mathy"
)

func main() {
	result := mathy.Sum(2.1, 2.1)
	fmt.Printf("The result of sum is: %.1f\r\n", result)

	operationResults := mathy.Calculation(mathy.Multiplier, 4, 5)
	fmt.Printf("The result of the operation using function Mutiplier as parameter of the function Calculation is: %.d\r\n", operationResults)

	divResults, rest := mathy.DividerWithRest(12, 5)
	fmt.Printf("The result of division is: %d and the rest of division is: %d\r\n", divResults, rest)

	fmt.Println("The result of multiplication is: ", mathy.Multiplier(6, 6))
}
