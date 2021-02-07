package main

import (
	"fmt"

	"github.com/douglasandradeee/GoExercicies/packages/operator"
	"github.com/douglasandradeee/GoExercicies/packages/prefix"
)

// UserName -
var UserName = "Douglas Andrade"

func main() {
	fmt.Printf("User name: %s\r\n", UserName)
	fmt.Printf("Capital prefix: %d\r\n", prefix.Capital)
	fmt.Printf("Operator name: %s\r\n", operator.OperatorName)
	fmt.Printf("Test Value: %s\r\n", prefix.TestWithPrefix)

}
